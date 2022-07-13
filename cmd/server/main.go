package main

import (
	"fmt"

	"github.com/92Sam/ms-users/infra"
	"github.com/92Sam/ms-users/utils"
)

func main() {
	errChanServer := make(chan error)
	statusChanServer := make(chan bool)

	utils.GetEnviroment("env")
	infra.InitApp(errChanServer, statusChanServer)

	for {
		select {
		case err := <-errChanServer:
			fmt.Println("Channel Error on server ->", err)
			panic(err)
		case status := <-statusChanServer:
			fmt.Println("Channel Init Server ->", status)
		default:
		}
	}
}
