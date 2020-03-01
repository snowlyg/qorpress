package posts

import (
	"bytes"
	"fmt"
	"html/template"

	// "github.com/jinzhu/gorm"
	"github.com/qorpress/admin"
	"github.com/qorpress/media"
	"github.com/qorpress/media/media_library"
	"github.com/qorpress/qor"
	"github.com/qorpress/render"

	"github.com/qorpress/qorpress-example/pkg/config/application"
	"github.com/qorpress/qorpress-example/pkg/models/posts"
	"github.com/qorpress/qorpress-example/pkg/utils/funcmapmaker"
)

// var Genders = []string{"Men", "Women", "Kids"}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	controller := &Controller{View: render.New(&render.Config{
		AssetFileSystem: application.AssetFS.NameSpace("posts"),
	}, "pkg/app/posts/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)

	application.Router.Get("/posts", controller.Index)
	application.Router.Get("/posts/{code}", controller.Show)
	// application.Router.Get("/{gender:^(men|women|kids)$}", controller.Gender)
	application.Router.Get("/category/{code}", controller.Category)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	// Produc Management
	Admin.AddMenu(&admin.Menu{Name: "Post Management", Priority: 1})
	// color := Admin.AddResource(&posts.Color{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -5})
	// Admin.AddResource(&posts.Size{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -4})
	// Admin.AddResource(&posts.Material{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -4})

	category := Admin.AddResource(&posts.Category{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -3})
	category.Meta(&admin.Meta{Name: "Categories", Type: "select_many"})

	collection := Admin.AddResource(&posts.Collection{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -2})

	// Add PostImage as Media Libraray
	PostImagesResource := Admin.AddResource(&posts.PostImage{}, &admin.Config{Menu: []string{"Post Management"}, Priority: -1})

	PostImagesResource.Filter(&admin.Filter{
		Name:       "SelectedType",
		Label:      "Media Type",
		Operations: []string{"contains"},
		Config:     &admin.SelectOneConfig{Collection: [][]string{{"video", "Video"}, {"image", "Image"}, {"file", "File"}, {"video_link", "Video Link"}}},
	})
	//PostImagesResource.Filter(&admin.Filter{
	//	Name:   "Color",
	//	Config: &admin.SelectOneConfig{RemoteDataResource: color},
	//})
	PostImagesResource.Filter(&admin.Filter{
		Name:   "Category",
		Config: &admin.SelectOneConfig{RemoteDataResource: category},
	})
	PostImagesResource.IndexAttrs("File", "Title")

	// Add Post
	post := Admin.AddResource(&posts.Post{}, &admin.Config{Menu: []string{"Post Management"}})
	// post.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: Genders, AllowBlank: true}})

	postPropertiesRes := post.Meta(&admin.Meta{Name: "PostProperties"}).Resource
	postPropertiesRes.NewAttrs(&admin.Section{
		Rows: [][]string{{"Name", "Value"}},
	})
	postPropertiesRes.EditAttrs(&admin.Section{
		Rows: [][]string{{"Name", "Value"}},
	})

	post.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
		{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
		{Name: "table", Source: "/vendors/redactor_table.js"},
	},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/admin/post_images",
		},
	}})
	post.Meta(&admin.Meta{Name: "Category", Config: &admin.SelectOneConfig{AllowBlank: true}})
	post.Meta(&admin.Meta{Name: "Collections", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})

	post.Meta(&admin.Meta{Name: "MainImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: PostImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"original": {Width: 560, Height: 700},
		},
	}})
	post.Meta(&admin.Meta{Name: "MainImageURL", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*posts.Post); ok {
			result := bytes.NewBufferString("")
			tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
			tmpl.Execute(result, map[string]string{"image": p.MainImageURL()})
			return template.HTML(result.String())
		}
		return ""
	}})

	post.Filter(&admin.Filter{
		Name:   "Collections",
		Config: &admin.SelectOneConfig{RemoteDataResource: collection},
	})

	post.Filter(&admin.Filter{
		Name: "Featured",
	})

	post.Filter(&admin.Filter{
		Name: "Name",
		Type: "string",
	})

	post.Filter(&admin.Filter{
		Name: "Code",
	})

	//post.Filter(&admin.Filter{
	//	Name: "Price",
	//	Type: "number",
	//})

	post.Filter(&admin.Filter{
		Name: "CreatedAt",
	})

	post.Action(&admin.Action{
		Name:        "Import Post",
		URLOpenType: "slideout",
		URL: func(record interface{}, context *admin.Context) string {
			return "/admin/workers/new?job=Import Posts"
		},
		Modes: []string{"collection"},
	})

	type updateInfo struct {
		CategoryID  uint
		Category    *posts.Category
		//MadeCountry string
		//Gender      string
	}

	updateInfoRes := Admin.NewResource(&updateInfo{})
	post.Action(&admin.Action{
		Name:     "Update Info",
		Resource: updateInfoRes,
		Handler: func(argument *admin.ActionArgument) error {
			newPostInfo := argument.Argument.(*updateInfo)
			for _, record := range argument.FindSelectedRecords() {
				fmt.Printf("%#v\n", record)
				if post, ok := record.(*posts.Post); ok {
					if newPostInfo.Category != nil {
						post.Category = *newPostInfo.Category
					}
					//f newPostInfo.MadeCountry != "" {
					//	post.MadeCountry = newPostInfo.MadeCountry
					//}
					//if newPostInfo.Gender != "" {
					//	post.Gender = newPostInfo.Gender
					//}
					argument.Context.GetDB().Save(post)
				}
			}
			return nil
		},
		Modes: []string{"batch"},
	})

	post.UseTheme("grid")

	// variationsResource := post.Meta(&admin.Meta{Name: "Variations", Config: &variations.VariationsConfig{}}).Resource
	// if imagesMeta := variationsResource.GetMeta("Images"); imagesMeta != nil {
	// 	imagesMeta.Config = &media_library.MediaBoxConfig{
	// 		RemoteDataResource: PostImagesResource,
	// 		Sizes: map[string]*media.Size{
	// 			"icon":    {Width: 50, Height: 50},
	// 			"thumb":   {Width: 100, Height: 100},
	// 			"display": {Width: 300, Height: 300},
	// 		},
	// 	}
	// }

	// variationsResource.EditAttrs("-ID", "-Post")
	// oldSearchHandler := post.SearchHandler
	// post.SearchHandler = func(keyword string, context *qor.Context) *gorm.DB {
	// 	context.SetDB(context.GetDB().Preload("Variations.Color").Preload("Variations.Size").Preload("Variations.Material"))
	// 	return oldSearchHandler(keyword, context)
	// }
	/*
	colorVariationMeta := post.Meta(&admin.Meta{Name: "ColorVariations"})
	colorVariation := colorVariationMeta.Resource
	colorVariation.Meta(&admin.Meta{Name: "Images", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: PostImagesResource,
		Sizes: map[string]*media.Size{
			"icon":    {Width: 50, Height: 50},
			"preview": {Width: 300, Height: 300},
			"listing": {Width: 640, Height: 640},
		},
	}})

	colorVariation.NewAttrs("-Post", "-ColorCode")
	colorVariation.EditAttrs("-Post", "-ColorCode")

	sizeVariationMeta := colorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	sizeVariation := sizeVariationMeta.Resource
	sizeVariation.EditAttrs(
		&admin.Section{
			Rows: [][]string{
				{"Size", "AvailableQuantity"},
				{"ShareableVersion"},
			},
		},
	)
	sizeVariation.NewAttrs(sizeVariation.EditAttrs())
	*/

	post.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
	post.IndexAttrs("MainImageURL", "Name", "Featured", "VersionName", "PublishLiveNow")
	post.EditAttrs(
		&admin.Section{
			Title: "Seo Meta",
			Rows: [][]string{
				{"Seo"},
			}},
		&admin.Section{
			Title: "Basic Information",
			Rows: [][]string{
				{"Name", "Featured"},
				{"Code"},
				{"MainImage"},
			}},
		&admin.Section{
			Title: "Organization",
			Rows: [][]string{
				{"Category"},
				{"Collections"},
			}},
		"PostProperties",
		"Description",
		// "ColorVariations",
		"PublishReady",
	)
	// post.ShowAttrs(post.EditAttrs())
	post.NewAttrs(post.EditAttrs())

	/*
	for _, gender := range Genders {
		var gender = gender
		post.Scope(&admin.Scope{Name: gender, Group: "Gender", Handler: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Where("gender = ?", gender)
		}})
	}
	*/

	post.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if post, ok := record.(*posts.Post); ok {
				return fmt.Sprintf("/posts/%v", post.Code)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})

}
