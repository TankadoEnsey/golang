// 123 project main.go
package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request, a string) {
	http.Redirect(w, r, a, 301)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect(w, r, "http://ya.ru")
	})

	//http.HandleFunc("/", redirect)
	http.ListenAndServe(":9090", nil)
}
