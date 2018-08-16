package main

import (
	"net/http"
	"os"
	"os/signal"

	"go.planetmeican.com/titan/log"
	"go.planetmeican.com/titan/logging"
)

func main() {
	logging.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		log.Info("say hello")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":3336", nil))
	}()

	log.Info("server started")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	log.Printf("stop signal: %v", <-ch)
}
