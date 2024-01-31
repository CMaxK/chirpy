package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	// Serve files from the '.' directory at the '/app/' path, after stripping the '/app' prefix from the request path.
	//eg "/app/assets/logo.png" would serve the file at "./assets/logo.png".
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	// Handle requests to /healthz path
	mux.HandleFunc("/healthz", handlerReadiness)

	// Serves headers
	corsMux := middlewareCors(mux)

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
