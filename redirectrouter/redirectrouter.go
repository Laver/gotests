package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"

	"bytes"

	"github.com/julienschmidt/httprouter"
	"github.com/laver/gotests/routes"
)

func main() {
	routeNum := 1000
	depth := 10
	routes := genRoutes(routeNum, depth)
	router := httprouter.New()
	loadRoutes(router, routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func redirect(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello")
	//code := 301 // Permanent redirect, request with GET method
	//http.Redirect(w, r, req.URL.String(), code)
}

func genRoutes(numRoutes int, depth int) bytes.Buffer {
	buf := routes.GenRoutes(numRoutes, depth)
	return buf
}

func loadRoutes(router *httprouter.Router, routes bytes.Buffer) {

	scanner := bufio.NewScanner(strings.NewReader(routes.String()))
	for scanner.Scan() {
		router.GET(scanner.Text(), redirect)
	}
}
