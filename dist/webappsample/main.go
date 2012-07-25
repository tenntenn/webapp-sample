package webappsample

import (
	"text/template"
	"net/http"
	"appengine"
	"code.google.com/p/goweb/goweb"
	"webappsample/task"
)

// Parse templates
var templates = template.Must(template.ParseGlob("template/*.html"))

// Initilize webapp
func init() {

	// register handler for "/"
	http.HandleFunc("/", index)

	// initilize task handlers
	task.Init()
}

// A handler for "/"
func index(w http.ResponseWriter, r *http.Request) {

	// Context of this request
	c := appengine.NewContext(r)
	
	// Execute index template
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		// execute error is occured
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}