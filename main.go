package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	r := chi.NewRouter()
	// Serve files from the '.' directory at the '/app/' path, after stripping the '/app' prefix from the request path.
	//eg "/app/assets/logo.png" would serve the file at "./assets/logo.png".
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	cfg := &apiConfig{}
	//configue routes with r.Handle and r.Get
	r.Handle("/app/*", cfg.middlewareMetricsInc(handler))
	r.Handle("/app", cfg.middlewareMetricsInc(handler))
	r.Get("/healthz", handlerReadiness)
	r.Get("/metrics", cfg.getHitsCountHandler)
	r.Get("/reset", cfg.resetHitsCountHandler)

	// Serves headers
	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

// Responds to HTTP requests with status code 200 and a 'Content-Type' header set to 'text/plain'
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) getHitsCountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hits: %d", cfg.fileserverHits)
}

func (cfg *apiConfig) resetHitsCountHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hits count reset to 0")
}
