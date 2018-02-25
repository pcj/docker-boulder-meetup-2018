package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"github.com/pcj/docker-boulder-meetup-2018/go/registry"
)

const (
	DockerRegistryPathEnvVarName = "DOCKER_REGISTRY_PATH"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func debuggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", dump)
		next.ServeHTTP(w, r)
	})
}

func main() {
	address := ":8000"
	
	r := mux.NewRouter()

	registry, err := registry.NewFilesystemRegistry(os.Getenv(DockerRegistryPathEnvVarName), r)
	if err != nil {
		log.Fatalf("Failed to initialize registry: %v (did you set %s?)", err, DockerRegistryPathEnvVarName)
	}
	
	//r.Use(debuggingMiddleware)
	r.Use(loggingMiddleware)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("NOT FOUND: %s", dump)
		log.Printf("NOT FOUND: %s %s", r.Method, r.RequestURI)
		http.NotFound(w, r)
	});
	
	// Bind to a port and pass our router in
	log.Printf("Serving docker registry protocol v2.2 at %s, rootDir=%s\n", address, registry.Path)
	log.Fatal(http.ListenAndServe(address, r))
}

