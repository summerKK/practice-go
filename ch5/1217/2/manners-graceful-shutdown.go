package main

import (
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
	"fmt"
)

type handler struct{}

func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)

}
func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
