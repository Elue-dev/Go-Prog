package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.gohtml")
	err:= parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := templateCache[t]
// 	if !inMap {
// 		// need to create the template
// 		fmt.Println("creating template and adding to cache")
// 		err := createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 	} else {
// 		// we have the template in the cache
// 		fmt.Println("using cached template")
// 	}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.gohtml",
// 	}

// 		// parse the template
// 		tmpl, err := template.ParseFiles(templates...) 
// 		if err !=nil {
// 			return err
// 		}

// 		// add template to cache (map)
// 		templateCache[t] = tmpl

// 		return nil
// }

func RenderTemplate(w http.ResponseWriter, t string) {
	// create template cache
	templateCache, err := createTemplateCache()
	if err !=nil {
		log.Fatal(err)
	}

	// get rendered template from cache
	template, ok := templateCache[t]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	err = template.Execute(buffer, nil)
	if err !=nil {
		fmt.Println(err)
	}

	// render the template
	_, err = buffer.WriteTo(w)
	if err !=nil {
		fmt.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// create entire cache at once, and populate that map with every available template, layout etc
	// get all of the files named *.page.gohtml from templates folder
	pages, err := filepath.Glob("/templates/*page.gohtml")
	if err != nil {
		return myCache, err
	}

	// range through all files starting with *page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("/templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}