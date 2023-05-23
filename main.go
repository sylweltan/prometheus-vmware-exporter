package main

import (
	// "./controller"
	// "github.com/sylweltan/prometheus-vmware-exporter"
	"flag"
	"fmt"
	"net/http"
	"os"

	controllers "github.com/sylweltan/prometheus-vmware-exporter/controller"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	listen   = ":9512"
	host     = ""
	username = ""
	password = ""
	logLevel = "info"
)

func env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}

func init() {
	flag.StringVar(&listen, "listen", env("ESX_LISTEN", listen), "listen port")
	flag.StringVar(&host, "host", env("ESX_HOST", host), "URL ESX host ")
	flag.StringVar(&username, "username", env("ESX_USERNAME", username), "User for ESX")
	flag.StringVar(&password, "password", env("ESX_PASSWORD", password), "password for ESX")
	flag.StringVar(&logLevel, "log", env("ESX_LOG", logLevel), "Log level must be, debug or info")
	flag.Parse()
	controllers.RegistredMetrics()
	collectMetrics()
}

func collectMetrics() {
	logger, err := initLogger()
	if err != nil {
		fmt.Println(err.Error())
	}
	go func() {
		logger.Debugf("Start collect host metrics")
		controllers.NewVmwareHostMetrics(host, username, password, logger)
		logger.Debugf("End collect host metrics")
	}()
	go func() {
		logger.Debugf("Start collect datastore metrics")
		controllers.NewVmwareDsMetrics(host, username, password, logger)
		logger.Debugf("End collect datastore metrics")
	}()
	go func() {
		logger.Debugf("Start collect VM metrics")
		controllers.NewVmwareVmMetrics(host, username, password, logger)
		logger.Debugf("End collect VM metrics")
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		collectMetrics()
	}
	h := promhttp.Handler()
	h.ServeHTTP(w, r)
}

func initLogger() (*log.Logger, error) {
	logger := log.New()
	logrusLogLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		return logger, err
	}
	logger.SetLevel(logrusLogLevel)
	logger.Formatter = &log.TextFormatter{DisableTimestamp: false, FullTimestamp: true}
	return logger, nil
}

func main() {
	logger, err := initLogger()
	if err != nil {
		logger.Fatal(err)
	}
	if host == "" {
		logger.Fatal("Yor must configured systemm env ESX_HOST or key -host")
	}
	if username == "" {
		logger.Fatal("Yor must configured system env ESX_USERNAME or key -username")
	}
	if password == "" {
		logger.Fatal("Yor must configured system env ESX_PASSWORD or key -password")
	}
	msg := fmt.Sprintf("Exporter start on port %s", listen)
	logger.Info(msg)
	http.HandleFunc("/metrics", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>VMware Exporter</title></head>
			<body>
			<h1>VMware Exporter</h1>
			<p><a href="` + "/metrics" + `">Metrics</a></p>
			</body>
			</html>`))
	})
	logger.Fatal(http.ListenAndServe(listen, nil))
}
