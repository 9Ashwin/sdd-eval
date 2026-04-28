package main

import (
	"net/http"

	"sdd-eval/evaluation/go/internal/handler"
	"sdd-eval/evaluation/go/internal/store"
)

func main() {
	s := store.New()
	h := handler.New(s)
	mux := http.NewServeMux()
	h.Register(mux)
	http.ListenAndServe(":3000", mux)
}
