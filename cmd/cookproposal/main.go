package main

import (
	"example.com/cookproposal/internal/api"
	"example.com/cookproposal/internal/flow027"
	"example.com/cookproposal/internal/store"
	"flag"
	"log"
	"net/http"
)

func main() {
	path := flag.String("db", "cookproposal.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, err := store.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
	srv := api.NewServer(flow027.NewHandler(st))
	log.Printf("cookproposal listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, srv.HandlerFunc()))
}
