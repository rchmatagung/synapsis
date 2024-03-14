package server

import (
	"fmt"
	"log"
	"synapsis/config"
	"synapsis/internal/middleware"
	"synapsis/internal/wrapper/handler"
	"synapsis/internal/wrapper/repository"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/infra/db"

	coreCategory "synapsis/internal/core/category"
	coreCustomers "synapsis/internal/core/customers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {

	//* Initial Engine
	engine := html.New("./views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:      conf.App.Name,
		ServerHeader: "Go Fiber",
		Views:        engine,
		BodyLimit:    conf.App.BodyLimit * 1024 * 1024,
	})

	//* Initial Data Middleware
	middleware.InitMiddlewareConfig(app, dbList, conf, appLoger)

	//* General Middleware
	middleware.CORSMiddleware()
	middleware.DefaultLimitterMiddleware()
	//middleware.RecoverMiddleware()

	//* Initial Wrapper

	if dbList.DatabaseApp == nil {
		log.Println("is nil")
	}

	repo := repository.NewRepository(conf, dbList, appLoger)
	usecase := usecase.NewUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, appLoger)

	//* Root Endpoint
	app.Get("/", handler.General.Root.GetRoot)

	//* Api Endpoint
	api := app.Group(conf.App.Endpoint)

	//* General Routes
	//generalEncyrption.NewRoutes(api, handler)

	//* Core Routes
	coreCategory.NewRoutes(api, handler)
	coreCustomers.NewRoutes(api, handler)

	//* CMS Routes
	// cmsWorkOfType.NewRoutes(api, handler)

	//* Not found
	app.All("*", handler.General.NotFound.GetNotFound)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.App.Port)))
}
