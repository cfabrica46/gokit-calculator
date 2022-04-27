package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cfabrica46/gokit-calculator/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func main() {
	log.SetFlags(log.Lshortfile)

	runServer()
}

func runServer() {
	portString := strconv.Itoa(port)

	svc := service.NewService()

	getAddHandler := httptransport.NewServer(
		service.MakeOperationEndpoint(svc, service.NewAddState()),
		service.DecodeRequest,
		service.EncodeResponse,
	)

	getSubtractHandler := httptransport.NewServer(
		service.MakeOperationEndpoint(svc, service.NewSubtractState()),
		service.DecodeRequest,
		service.EncodeResponse,
	)

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/add").Handler(getAddHandler)
	router.Methods(http.MethodPost).Path("/subtract").Handler(getSubtractHandler)

	log.Println("ListenAndServe on localhost:" + portString)
	log.Fatal(http.ListenAndServe(":"+portString, router))
}
