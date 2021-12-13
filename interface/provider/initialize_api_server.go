package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mocoog/bff/interface/gen/restapiv1"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations"
	"github.com/mocoog/bff/interface/gen/restapiv1/operations/basis"
	gPB "github.com/mocoog/grpc-go-packages/gateway/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             5 * time.Second,
	PermitWithoutStream: true,
}

func healthCheckImp(params basis.GetHealthParams) middleware.Responder {
	address := "172.23.0.1:9881"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		// grpc.WithBlock(),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		panic(errors.Wrap(err, "コネクションエラー"))
	}
	defer conn.Close()

	client := gPB.NewGatewayServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	healthCheckRequest := gPB.HealthCheckRequest{
		Status: 200,
	}

	res, err := client.HealthCheck(ctx, &healthCheckRequest)

	if err != nil {
		log.Println(errors.Wrap(err, "受取り失敗"))
	} else {
		log.Printf("サーバからの受け取り %d", res.GetStatus())
	}

	// return request(client, a, b)
	return basis.NewGetHealthOK()
}

func InitializeAPIServer() (*restapiv1.Server, error) {
	swaggerSpec, err := loads.Analyzed(restapiv1.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("InitializeAPIServer fail: %w", err)
	}

	api := operations.NewMocoogBffServerAPI(swaggerSpec)
	api.BasisGetHealthHandler = basis.GetHealthHandlerFunc(healthCheckImp)
	return restapiv1.NewServer(api), nil
}
