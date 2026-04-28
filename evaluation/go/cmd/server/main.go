package main

import (
	"net/http"

	"openspec-superpowers-guide/evaluation/go/internal/handler"
	"openspec-superpowers-guide/evaluation/go/internal/store"
)

func main() {
	s := store.New()
	h := handler.New(s)
	mux := http.NewServeMux()
	h.Register(mux)
	http.ListenAndServe(":3000", mux)
}
