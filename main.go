package main

import (
	"log"
	"net/http"

	"github.com/cfabrica46/gokit-calculator/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Lshortfile)

	runServer("8080")
}

func runServer(port string) {
	svc := service.NewService()

	getAddHandler := httptransport.NewServer(
		service.MakeAddEndpoint(svc),
		service.DecodeRequest,
		service.EncodeResponse,
	)

	getSubtractHandler := httptransport.NewServer(
		service.MakeSubtractEndpoint(svc),
		service.DecodeRequest,
		service.EncodeResponse,
	)

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/add").Handler(getAddHandler)
	router.Methods(http.MethodPost).Path("/subtract").Handler(getSubtractHandler)

	log.Println("ListenAndServe on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
