package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"bytes"

	"github.com/julienschmidt/httprouter"
	"github.com/laver/routes"
)

func main() {

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

func loadRoutes(routes bytes.Buffer) {
	router := httprouter.New()
	scanner := bufio.NewScanner(strings.NewReader(routes.String()))
	for scanner.Scan() {
		router.GET(scanner.Text(), redirect)
	}
}
