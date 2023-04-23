package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Iann221/goBookings/pkg/config"
	"github.com/Iann221/goBookings/pkg/handler"
	"github.com/Iann221/goBookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig // ditaro di luar agar bisa dipake di middleware juga yang packagenya juga main
var session *scs.SessionManager

func main() {

	// change to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // sessionnya mau 24 jam
	session.Cookie.Persist = true // should the session persist when the browser window is closed? true
	session.Cookie.SameSite = http.SameSiteLaxMode // how strict u wanna be about what site this cookie applies to
	session.Cookie.Secure = app.InProduction // mastiin cookie encrypted. di production hrs true, klo localhost false
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false // klo false, klo page.tmpl diupdate, auto kerender ulang

	// set the appconfig to render.go
	render.NewTemplates(&app)

	// untuk membuat objek repository di handler.go
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// /: url that listens to this. jika client ke /, here's what i want to give them.
		// handlFunc ini listening ot a request sent by a web browser
		n, err := fmt.Fprintf(w, "hello world") // fprintf itu bwt nulis hello world di webnya
		if err != nil {
			fmt.Println("ada error")
		}
		fmt.Println(n)
	})

	// http.HandleFunc("/home", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	// start a webserver that listens for a request. jdi ini listen to port 8080 dengan no handler
	fmt.Printf("starting on %s", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
