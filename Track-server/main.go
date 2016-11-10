// Track-server project main.go
package main

import (
	"encoding/base64"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.EscapedPath()[1:]
		url, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			log.Fatal("This address does not exist or an error occurred:", err)
		}

		http.Redirect(w, r, "http://"+string(url), http.StatusFound)
	})

	http.ListenAndServe(":8085", nil)
}
