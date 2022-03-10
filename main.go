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
	// http.ListenAndServe(":8080", http.HandlerFunc(redirect))
	http.ListenAndServeTLS(":8443", "/etc/letsencrypt/live/www.knatterburg.com/fullchain.pem", "/etc/letsencrypt/live/www.knatterburg.com/privkey.pem", http.HandlerFunc(greet))
}
