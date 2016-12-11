package routes

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//WriteRouteFile generates and writes routes out to a file
func WriteRouteFile(numRoutes int, maxDepth int, filePath string) {
	// Generate the route tree
	r := GenRoutes(numRoutes, maxDepth)
	// Count the lines in the route tree
	lines := strings.Count(r.String(), "\n")
	// Report the number of lines
	fmt.Printf("Generated %d routes\n", lines)

	// write the Route tree to a text file
	writeErr := ioutil.WriteFile(filePath, r.Bytes(), 0644)
	check(writeErr)
}

// GenRoutes Generate N routes in a buffer
func GenRoutes(numRoutes int, maxDepth int) bytes.Buffer {
	foldernames := []string{"wibble", "wobble", "flibble", "bubble", "trobble", "babble"}
	var buffer bytes.Buffer
	routeCount := 0
	if maxDepth == 0 || len(foldernames) == 0 {
		return buffer
	}
	// figure out how many loops we need to do for the given number of routes based on depth and how many folder names we have
	iterations := numRoutes / maxDepth / len(foldernames)
	for i := 0; i <= iterations; i++ {
		// loop to the max depth
		for d := 0; d < maxDepth; d++ {
			// create a base path for the child route
			strRoute := childRoute("test"+strconv.Itoa(i), d)
			// loop over the folder names to generate some child paths
			for _, c := range foldernames {
				// only run if we've not hit desired total number of routes
				if routeCount < numRoutes {
					buffer.WriteString(strRoute)
					buffer.WriteString(c)
					buffer.WriteString("\n")
					routeCount++
				}
			}
		}
	}
	return buffer
}

func childRoute(label string, depth int) string {
	var buffer bytes.Buffer
	depthStr := strconv.Itoa(depth)

	buffer.WriteString("/base")
	buffer.WriteString(label)
	buffer.WriteString("/")

	for i := 0; i < depth; i++ {
		name := subNameBuf(depthStr, i)
		buffer.Write(name.Bytes())
		buffer.WriteString("/")
	}
	return buffer.String()
}

func subNameBuf(label string, depth int) bytes.Buffer {
	var subBuf bytes.Buffer
	for j := 0; j < depth; j++ {
		subBuf.WriteString("sub")
	}
	subBuf.WriteString("path")
	subBuf.WriteString(label)
	return subBuf
}
