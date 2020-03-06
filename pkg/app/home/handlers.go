package home

import (
	"net/http"

	"github.com/gopress/internal/qor"
	"github.com/gopress/internal/qor/utils"
	"github.com/gopress/internal/render"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
}

// SwitchLocale switch locale
func (ctrl Controller) SwitchLocale(w http.ResponseWriter, req *http.Request) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: req.URL.Query().Get("locale")}, &qor.Context{Request: req, Writer: w})
	http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
}
