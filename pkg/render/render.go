package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Iann221/goBookings/pkg/config"
	"github.com/Iann221/goBookings/pkg/models"
)

var app *config.AppConfig

// set the config for the template function
func NewTemplates(a *config.AppConfig) {
	app = a
}

// bikin default model klo2 ada data yg mau dipake di banyak template
func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	// specify atribut yg pgn dipake publik disini
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) { // loads file from root level

	var tc map[string]*template.Template
	if app.UseCache{
		// get template cache from appconfig
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("err")
	}

	// create buffer (optional) for final grain error checking. utk mastiin beneran bisa diexecute
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf,td)

	// render the template
	_, err := buf.WriteTo(w)
	if(err != nil){
		fmt.Println("error writing template to browser",err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	// bisa jg create map kek gini
	myCache := map[string]*template.Template{}

	// get all files *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if(err != nil){
		return myCache, err
	}

	// range through all *.pages.tmpl files
	for _, page := range pages {
		name := filepath.Base(page) // return last element of path jdi cm about.page.tmpl
		ts, err := template.New(name).ParseFiles(page) // New: give the template a name
		if(err != nil){
			return myCache, err
		}
		
		// get all layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if(err != nil){
			return myCache, err
		}

		if len(matches) > 0 {
			// biasanya kan parseFile file page dan layoutnya. ini pagenya udah diparse, ditambah layoutnya
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if(err != nil){
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}