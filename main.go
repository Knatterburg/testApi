package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.ListenAndServe("knatterburg.com:8080", http.HandlerFunc(greet))
	// r := mux.NewRouter()
	// r.HandleFunc("/", greet)
	// err := http.ListenAndServeTLS(":8443", "/etc/letsencrypt/live/knatterburg.com/fullchain.pem", "/etc/letsencrypt/live/knatterburg.com/privkey.pem", r)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}
