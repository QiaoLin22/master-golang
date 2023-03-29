package main

import (
	"flag"
	"net/http"

	"github.com/QiaoLin22/goFolder/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":49999", "todo")
	flag.Parse()

	http.HandleFunc("/user", api.HandleGetUser)
	http.HandleFunc("/account", api.HandleGetAccount)
	http.ListenAndServe(*listenAddr, nil)
}
