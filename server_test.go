package main

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mocoog/bff/interface/gen/restapiv1"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations/basis"
)

func TestHealthCheck(t *testing.T) {
	t.Helper()
	cases := []struct {
		Name string
		Path string
	}{
		{
			Name: "(Success)Health Check",
			Path: "/health",
		},
	}
	api, err := initializeAPI()
	if err != nil {
		t.Errorf("Failed Health Check! 1: %v", err.Error())
	}
	server := restapiv1.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag
	api.BasisGetHealthHandler = basis.GetHealthHandlerFunc(
		func(params basis.GetHealthParams) middleware.Responder {
			return basis.NewGetHealthOK()
		})
	server.ConfigureAPI()

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ts := httptest.NewServer(server.GetHandler())
			defer ts.Close()
			res, err := http.Get(ts.URL + "/v1/health")
			if err != nil {
				t.Errorf("Failed Health Check! 2: %v", err.Error())
			}
			if res.StatusCode != 200 {
				t.Errorf("Failed Health Check! 2: %v", "Invalid Status Code")
			}
		})
	}
}
