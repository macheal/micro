package main

import (
	"example-service/handler"
	example "example-service/proto"

	"gitee.com/smartsteps/go-micro/v2"
	"gitee.com/smartsteps/go-micro/v2/util/log"
	"github.com/micro/micro/v2/cmd/protoc-gen-micro/plugin/micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.example"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
