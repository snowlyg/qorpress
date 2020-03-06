package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// "github.com/go-chi/docgen"
	"github.com/foomo/simplecert"
	"github.com/gopress/internal/admin"
	"github.com/gopress/internal/publish2"
	"github.com/gopress/internal/qor"
	"github.com/gopress/internal/qor/utils"

	"github.com/gopress/internal/qorpress/pkg/app/account"
	adminapp "github.com/gopress/internal/qorpress/pkg/app/admin"
	"github.com/gopress/internal/qorpress/pkg/app/api"
	"github.com/gopress/internal/qorpress/pkg/app/home"
	"github.com/gopress/internal/qorpress/pkg/app/pages"
	"github.com/gopress/internal/qorpress/pkg/app/posts"
	"github.com/gopress/internal/qorpress/pkg/app/static"
	"github.com/gopress/internal/qorpress/pkg/config"
	"github.com/gopress/internal/qorpress/pkg/config/application"
	"github.com/gopress/internal/qorpress/pkg/config/auth"
	"github.com/gopress/internal/qorpress/pkg/config/bindatafs"
	"github.com/gopress/internal/qorpress/pkg/config/db"
	_ "github.com/gopress/internal/qorpress/pkg/config/db/migrations"
	"github.com/gopress/internal/qorpress/pkg/utils/funcmapmaker"
)

func main() {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	cmdLine.Parse(os.Args[1:])

	var (
		Router = chi.NewRouter()
		Admin  = admin.New(&admin.AdminConfig{
			SiteName: "QORPRESS DEMO",
			Auth:     auth.AdminAuth{},
			DB:       db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
		})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// for demo, don't use this for your production site
			// to do: add to the yaml configuration file
			w.Header().Add("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.Header.Del("Authorization")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var (
				tx         = db.DB
				qorContext = &qor.Context{Request: req, Writer: w}
			)
			if locale := utils.GetLocale(qorContext); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}
			ctx := context.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	})

	Application.Use(api.New(&api.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(posts.New(&posts.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(pages.New(&pages.Config{}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/system"},
		Handler: utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))),
	}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"},
		Handler: bindatafs.AssetFS.FileServer(http.Dir(filepath.Join("themes", "qorpress", "public")), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"),
	}))

	// fmt.Println(docgen.MarkdownRoutesDoc(Application.Router, docgen.MarkdownOpts{ForceRelativeLinks: false}))

	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {

		if config.Config.HTTPS {
			fmt.Print("Listening on: 443\n")
			if err := simplecert.ListenAndServeTLS(":443", Application.NewServeMux(), "x0rzkov@protonmail.com", nil, "x0rzkov.com", "www.x0rzkov.com"); err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Listening on: %v\n", config.Config.Port)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), Application.NewServeMux()); err != nil {
				panic(err)
			}
		}
	}
}
