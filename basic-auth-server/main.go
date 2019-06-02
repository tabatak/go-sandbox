package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const (
	basicAuthUser     = "user"
	basicAuthPassword = "password"
)

func main() {
	http.HandleFunc("/basic",
		func(w http.ResponseWriter, r *http.Request) {
			dump, _ := httputil.DumpRequest(r, true)
			fmt.Println(string(dump))

			if user, pass, ok := r.BasicAuth(); !ok || user != basicAuthUser || pass != basicAuthPassword {
				fmt.Println("Not Authed!")

				w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
				w.WriteHeader(http.StatusUnauthorized) // 401
				http.Error(w, "Not authorized", 401)
				return
			}

			fmt.Println("Authed!")
			fmt.Fprintf(w, "Authed!!")
		},
	)
	http.ListenAndServe(":18888", nil)
}
