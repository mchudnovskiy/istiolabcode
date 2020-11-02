package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	for {
		resp, err := http.Get("http://http-server:8686/ping")
		if err != nil {
			sugar.Infof("An error occured during the ping request: %s", err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					sugar.Error(err)
				}
				sugar.Infof("A pong response is recieved on the ping request: << %s >>", string(bodyBytes))
			}			
		}
		time.Sleep(time.Second)
	}
}
