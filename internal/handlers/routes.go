package handlers

import "net/http"

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", ListProducts)
	mux.HandleFunc("/", SayHi)
	return mux
}
