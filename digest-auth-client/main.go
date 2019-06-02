package main

import (
	"fmt"
	"net/http/httputil"
	"os"

	"github.com/bobziuchkovski/digest"
)

func main() {
	t := digest.NewTransport("digest-auth-username", "digest-auth-password")
	c, err := t.Client()
	if err != nil {
		os.Exit(0)
	}
	resp, err := c.Get("http://127.0.0.1:18888/digest")
	if err != nil {
		os.Exit(0)
	}
	dump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))

}
