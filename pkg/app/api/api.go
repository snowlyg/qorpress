package api

import (
	"github.com/gopress/internal/admin"
	"github.com/gopress/internal/qor"

	"github.com/gopress/qorpress/pkg/config/application"
	"github.com/gopress/qorpress/pkg/config/db"

	// "github.com/gopress/qorpress/pkg/models/orders"
	"github.com/gopress/qorpress/pkg/models/posts"
	"github.com/gopress/qorpress/pkg/models/users"
)

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/api"
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
	API := admin.New(&qor.Config{DB: db.DB})

	API.AddResource(&posts.Post{})

	API.AddResource(&users.User{})
	// User := API.AddResource(&users.User{})
	// userOrders, _ := User.AddSubResource("Orders")
	// userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})

	API.AddResource(&posts.Category{})

	application.Router.Mount(app.Config.Prefix, API.NewServeMux(app.Config.Prefix))
}
