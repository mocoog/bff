package provider

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mocoog/bff/interface/gen/restapiv1"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations/basis"
)

func InitializeAPIServer() (*restapiv1.Server, error) {
	swaggerSpec, err := loads.Analyzed(restapiv1.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("InitializeAPIServer fail: %w", err)
	}

	api := operations.NewMocoogBffServerAPI(swaggerSpec)
	api.BasisGetHealthHandler = basis.GetHealthHandlerFunc(
		func(params basis.GetHealthParams) middleware.Responder {
			return basis.NewGetHealthOK()
		})
	return restapiv1.NewServer(api), nil
}
