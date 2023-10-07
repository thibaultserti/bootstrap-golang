package main

import (
	"fmt"
	"net/http"
	"os"

	"bootstrap-golang/pkg/config"

	logging "github.com/sirupsen/logrus"
)

func main() {

	var configuration config.Configuration

	configPath := "config/configuration.yaml"
	configuration, err := config.LoadConfig(configPath)
	if err != nil {
		logging.Fatal("Cannot load config")
	}
	level, err := logging.ParseLevel(configuration.LogLevel)
	if err != nil {
		logging.Fatal("Log level invalid")
	}

	if configuration.Env != "prod" && configuration.Env != "prd" {
		logging.SetFormatter(&logging.TextFormatter{})
	} else {
		logging.SetFormatter(&logging.JSONFormatter{})
	}
	logging.SetOutput(os.Stdout)
	logging.SetLevel(level)

	logging.Info(fmt.Sprintf("Serving on http://%s:%s ...", configuration.Hostname, configuration.Port))
	http.HandleFunc("/", HelloServer)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configuration.Port), nil)
	if err != nil {
		logging.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
