package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/wickywaa/hotelgav/pkg/config"
	"github.com/wickywaa/hotelgav/pkg/handlers"
	"github.com/wickywaa/hotelgav/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	const port string = ":8080"

	fmt.Println(fmt.Sprintf("listening on port %s", port))

	srv := &http.Server{

		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
