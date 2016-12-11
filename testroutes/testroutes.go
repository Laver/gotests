package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/laver/gotests/routes"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// generate a text file that contains n test routes

// generate a text file that contains n test routes
func main() {

	routeNumArg := "1000"
	routeDepthArg := "6"
	if len(os.Args) > 2 {
		routeNumArg = os.Args[1]
		routeDepthArg = os.Args[2]
	}

	numRoutes, err := strconv.Atoi(routeNumArg)
	if err != nil {
		fmt.Println("error routeNumArg")
	}

	maxDepth, err2 := strconv.Atoi(routeDepthArg)
	if err2 != nil {
		fmt.Println("error routeDepthArg")
	}
	fmt.Printf("numRoutes: %d maxDepth: %d\n", numRoutes, maxDepth)

	filePath := "./routes.txt"
	routes.WriteRouteFile(numRoutes, maxDepth, filePath)
}
