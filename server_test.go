package main

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mocoog/bff/interface/provider"
)

func TestHealthCheck(t *testing.T) {
	t.Helper()
	cases := []struct {
		Name       string
		Path       string
		StatusCode int
	}{
		{
			Name:       "(Success)Health Check",
			Path:       "/v1/health",
			StatusCode: 200,
		},
		{
			Name:       "(Success)Not Found",
			Path:       "/v1/aaaas",
			StatusCode: 404,
		},
	}
	server, _ := provider.InitializeAPIServer()
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag
	server.ConfigureAPI()

	ts := httptest.NewServer(server.GetHandler())
	defer ts.Close()

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res, err := http.Get(ts.URL + c.Path)
			if err != nil {
				t.Errorf("Failed %s: %v", c.Path, err.Error())
			}
			if res.StatusCode != c.StatusCode {
				t.Errorf("Failed %s: Invalid Status Code", c.Path)
			}
		})
	}
}
