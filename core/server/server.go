package server

import (
	"strconv"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/hypebeast/gojistatic"
	"github.com/rs/cors"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/routes"
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"net/http"
	"github.com/remony/lemonlog"
	"github.com/remony/Equipment-Rental-API/core/utils/cron"
)
// Start handles all route configuration and starts the http server
func Start(settings config.Config, context config.Context, mode int) {
	log.Println("こんにちは, listening on port :" + strconv.Itoa(settings.Production.Port))

	// Create the main router
	masterRouter := web.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
		OptionsPassthrough : false,
	})

	//Create subroutes
	apiRouter := web.New()
	imageRouter := web.New()

	// Assign sub routes to handle certain path requests

	masterRouter.Handle("/api/*", apiRouter)
	masterRouter.Handle("/data/*", imageRouter)
	masterRouter.Handle("/*", http.FileServer(http.Dir("C:/Users/remon/Git/equipment-rental-website/app")))

	// Apply SubRouter middleware to allow sub routing
	apiRouter.Use(middleware.SubRouter)
	imageRouter.Use(middleware.SubRouter)

	imageRouter.Use(gojistatic.Static("data/images/", gojistatic.StaticOptions{
		SkipLogging: true,
		Expires: nil,
	}))

	// Apply the CORS options to the main route handler
	masterRouter.Use(c.Handler)
	// apiRouter.Use(c.Handler)
	// imageRouter.Use(c.Handler)

	if (mode == 1) {
		masterRouter.Use(lemonlog.Logger)
	}
	// Create the routes
	routes.CreateRoutes(router.API{Router:apiRouter, Context:context, Config:settings})

	// Start Cron Job
	go cron.Cron(router.API{Router:apiRouter, Context:context, Config:settings});

	// Gracefully Serve
	if portIsFree(settings.Production.Port) {
		err := graceful.ListenAndServe(":" + strconv.Itoa(settings.Production.Port), masterRouter)
		if err != nil {
			//		If an error occurs, normally is port is already in use

			//		Don't panic
			panic(err)
		}
	}
}

// Checks if a port is free
func portIsFree(port int) bool {
	//	If the port is being used

	//	Return false

	//	if not in use
	return true
}

