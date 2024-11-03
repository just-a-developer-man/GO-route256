package main

import (
	"log"
	"net/http"
	"os"

	controller_http "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/controller/http"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/controller/http/middleware"
	repository "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/repository/postgres"
	wms "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/services/WMS"
	oms "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/usecase/OMS"
)

func main() {
	// Repository layer
	omsRepo := repository.NewOMSRepostiory( /* ... */ )

	// Other external services (adapters) layer
	wmsService := wms.NewWMSService( /* ... */ )

	// Usecase layer
	omsUsecase := oms.NewOMSUsecase(oms.Deps{
		WarehouseManagementSystem: wmsService,
		OMSRepository:             omsRepo,
	})

	// Delivery || Gateway || Transport || Controller layer
	controller := controller_http.NewController(controller_http.Usecases{
		OrderManagementSystem: omsUsecase,
	})

	// Router layer
	router := controller.NewRouter()

	// Middleware layer
	router = middleware.AddLoggingCtxMiddleware(router)
	router = middleware.WithHTTPRecoverMiddleware(router)
	router = middleware.LogHandlingTimeMiddleware(router)

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
