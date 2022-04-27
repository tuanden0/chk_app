package v1

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	handler *handler
}

func NewApp(handler *handler) (a *app) {
	return &app{handler: handler}
}

func (a *app) RunServer() {

	log.Printf("server start at: %v\n", SERVER_ADDR)
	server := &http.Server{Addr: SERVER_ADDR, Handler: newRouter(a.handler)}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {

		<-sig
		log.Println("got shutdown signal... shutdown server!")

		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}

		shutdownCancel()
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
