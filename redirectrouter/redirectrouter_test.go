package main

import "testing"

func BenchmarkRoutes100(b *testing.B) { // generate a text file that contains n test routes
	routeNum := 100
	depth := 10
	routes := genRoutes(routeNum, depth)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		loadRoutes(routes)
	}
}
func BenchmarkRoutes1000(b *testing.B) { // generate a text file that contains n test routes
	routeNum := 1000
	depth := 10
	routes := genRoutes(routeNum, depth)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		loadRoutes(routes)
	}
}
func BenchmarkRoutes10000(b *testing.B) { // generate a text file that contains n test routes
	routeNum := 10000
	depth := 10
	routes := genRoutes(routeNum, depth)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		loadRoutes(routes)
	}
}
func BenchmarkRoutes100000(b *testing.B) { // generate a text file that contains n test routes
	routeNum := 100000
	depth := 10
	routes := genRoutes(routeNum, depth)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		loadRoutes(routes)
	}
}
