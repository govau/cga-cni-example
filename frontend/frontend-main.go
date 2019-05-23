package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://backend-cni-test.apps.internal:8080")
	if err != nil {
		fmt.Fprintf(w, "Error talking to backend: %s\n", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
