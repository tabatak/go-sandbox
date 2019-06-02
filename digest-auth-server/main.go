package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/digest",
		func(w http.ResponseWriter, r *http.Request) {

			dump, _ := httputil.DumpRequest(r, true)
			fmt.Println(string(dump))

			if _, ok := r.Header["Authorization"]; !ok {
				fmt.Println("Not Authed!")
				w.Header().Add("WWW-Authenticate", `Digest realm="my private area", nonce="RMH1usDrAwA=6dc290ea3304de42a7347e0a94089ff5912ce0de", algorithm=MD5, qop="auth"`)
				w.WriteHeader(http.StatusUnauthorized) // 401
				http.Error(w, "Not authorized", 401)
				return
			}

			fmt.Println("Authed!")
			fmt.Fprintf(w, "Authed!")
		},
	)
	http.ListenAndServe(":18888", nil)
}
