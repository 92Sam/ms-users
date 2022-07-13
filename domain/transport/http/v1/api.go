package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartServer(router *mux.Router, errChanServer chan error, statusChanServer chan bool) {
	portServer := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	srv := &http.Server{
		Addr:    portServer,
		Handler: router,
	}

	go func() {
		fmt.Println("Server Started On Port " + portServer)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChanServer <- err
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
