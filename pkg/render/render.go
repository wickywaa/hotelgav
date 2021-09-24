package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wickywaa/hotelgav/pkg/config"
	"github.com/wickywaa/hotelgav/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {

	app = a

}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}

	}

	t, ok := tc[tmpl]
	fmt.Println("hello", t)

	if !ok {

		log.Fatal("could not get templeate from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {

		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {

			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {

			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {

				return myCache, err
			}
		}

		myCache[name] = ts
		fmt.Println("page is cureently ", name)

	}
	fmt.Println("phaylo")
	return myCache, nil

}
