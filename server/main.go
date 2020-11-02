package main

import (
	"fmt"
	"net/http"
	"os"
	"go.uber.org/zap"
	"log"

)

func pingHandler(w http.ResponseWriter, req *http.Request) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	sugar.Infof("Recieved ping request. Preparing pong response: << Pong from %s >>", hostName)
	fmt.Fprintf(w, "Pong from %s\n", hostName)
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("Lab server is started on port 8686 and is waiting for requests.")
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8686", nil)
}