package handler

import (
	"net/http"

	"github.com/Iann221/goBookings/pkg/config"
	"github.com/Iann221/goBookings/pkg/models"
	"github.com/Iann221/goBookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr // eve
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // nyimpen value ke session
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	//some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

