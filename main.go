package main

import (
	"bivrost_task2/controller"
	"bivrost_task2/database"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/koinworks/asgard-bivrost/libs"
	bv "github.com/koinworks/asgard-bivrost/service"
	hmodels "github.com/koinworks/asgard-heimdal/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database.StartDB()
}

func main() {
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	portNumber, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	serviceConfig := &hmodels.Service{
		Class:     "product-service",
		Key:       os.Getenv("APP_KEY"),
		Name:      os.Getenv("APP_NAME"),
		Version:   os.Getenv("APP_VERSION"),
		Host:      hostname,
		Port:      portNumber,
		Namespace: os.Getenv("K8S_NAMESPACE"),
		Metas:     make(hmodels.ServiceMetas),
	}

	registry, err := libs.InitRegistry(libs.RegistryConfig{
		Address:  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Service:  serviceConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	server, err := libs.NewServer(registry)
	if err != nil {
		log.Fatal(err)
	}

	bivrostSvc := server.AsGatewayService(
		"/v1",
	)

	bivrostSvc.Get("/ping", pingHandler)
	bivrostSvc.Post("/item", controller.CreateItem)
	bivrostSvc.Post("/order", controller.CreateOrder)
	bivrostSvc.Get("/items", controller.GetItems)
	bivrostSvc.Get("/orders", controller.GetOrders)

	err = server.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to Ping API",
			"id": "Selamat datang di Ping API",
		},
	})

}
