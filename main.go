package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Handle requests with the echoHandler function
	http.HandleFunc("/", echoHandler)

	// Start the HTTP server on port 8080
	port := 9154
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// echoHandler echoes back the received request's body as the response.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Write the request body back as the response
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
