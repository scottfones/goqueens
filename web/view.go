// Generalizes View construction

package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Shortcut to layout directory
var layoutDir string = "web/templates"

// NewView prepares a view for export.
func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View consists of a template and layout.
type View struct {
	Template *template.Template
	Layout   string
}

// Render returns a constructed template
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles adds all gohtml files to a slice
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "/*.gohtml")
	if err != nil {
		log.Println(err)
	}
	return files
}
