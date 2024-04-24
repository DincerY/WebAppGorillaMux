package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	//r.Handle("/", MyHandler{})

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{key}", ProductsHandler)
	r.HandleFunc("/query", DenemeHandler).Queries("a", "abc")

	r.Handle("/deneme", MyHandler{})

	r.PathPrefix("/help").Handler(MyHandler{})
	r.PathPrefix("/peh").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusUnauthorized)
	})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	server := http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Hata alındı")
	}

}

type MyHandler struct {
}

func (d MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

}
func ProductsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	if vars["key"] == "abc" {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func DenemeHandler(writer http.ResponseWriter, request *http.Request) {

}
