package admin

import (
	"github.com/gopress/internal/action_bar"
	"github.com/gopress/internal/admin"
	"github.com/qorpress/help"
	"github.com/gopress/internal/media/asset_manager"
	"github.com/gopress/internal/media/media_library"

	// "github.com/gopress/qorpress/pkg/models/seo"
	"github.com/gopress/qorpress/pkg/config/application"
	"github.com/gopress/qorpress/pkg/config/i18n"
	"github.com/gopress/qorpress/pkg/models/settings"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

// AssetManager asset manager
var AssetManager *admin.Resource

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/admin"
	}
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	Admin := application.Admin

	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	ActionBar.RegisterAction(&action_bar.Action{
		Name: "New Post",
		Link: "/admin/posts/new",
	})

	// Add Translations
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: -1})

	// Add Setting
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Blog Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	SetupSEO(Admin)
	SetupWidget(Admin)
	SetupDashboard(Admin)

	application.Router.Mount(app.Config.Prefix, Admin.NewServeMux(app.Config.Prefix))
}
