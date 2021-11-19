package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mocoog/bff/interface/gen/restapiv1"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations/basis"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func initializeAPI() (*operations.MocoogBffServerAPI, error) {
	swaggerSpec, err := loads.Analyzed(restapiv1.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	api := operations.NewMocoogBffServerAPI(swaggerSpec)
	return api, nil
}

func main() {
	api, err := initializeAPI()
	if err != nil {
		log.Fatalln(err)
	}
	server := restapiv1.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	api.BasisGetHealthHandler = basis.GetHealthHandlerFunc(
		func(params basis.GetHealthParams) middleware.Responder {
			return basis.NewGetHealthOK()
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
