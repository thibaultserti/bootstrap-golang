package main

import (
	"fmt"
	"net/http"
	"os"

	"bootstrap-golang/pkg/config"

	logging "github.com/sirupsen/logrus"
)

var port string
var hostname string

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logging.SetFormatter(&logging.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logging.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	logging.SetLevel(logging.InfoLevel)

	configPath := "config/configuration.yaml"
	viper, err := config.LoadConfig(configPath)
	if err != nil {
		logging.Errorf("Fail to read file %s", configPath)
	}

	port = viper.GetString("port")
	hostname = viper.GetString("hostname")

}
func main() {
	logging.Info(fmt.Sprintf("Serving on http://%s:%s ...", hostname, port))
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		logging.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
