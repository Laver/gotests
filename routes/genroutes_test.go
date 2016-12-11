package routes

import (
	"strings"
	"testing"
)

func TestGenRoutes(t *testing.T) {

	cases := []struct {
		in    int
		depth int
		want  string
	}{
		{8, 2, "/basetest0/wibble\n/basetest0/wobble\n/basetest0/flibble\n/basetest0/bubble\n/basetest0/trobble\n/basetest0/babble\n/basetest0/path1/wibble\n/basetest0/path1/wobble\n"},
		{3, 5, "/basetest0/wibble\n/basetest0/wobble\n/basetest0/flibble\n"},
		{2, 5, "/basetest0/wibble\n/basetest0/wobble\n"},
		{1, 5, "/basetest0/wibble\n"},
		{0, 5, ""},
		{0, 0, ""},
	}
	for _, c := range cases {
		buf := GenRoutes(c.in, c.depth)
		got := buf.String()
		if got != c.want {
			t.Errorf("genRoutes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGenRoutesLength(t *testing.T) {

	cases := []struct {
		count int
		depth int
		want  int
	}{
		{8, 2, 8},
		{100, 2, 100},
		{100, 50, 100},
		{1001, 2, 1001},
		{1001, 1, 1001},
		{1001, 5, 1001},
		{1001, 10, 1001},
		{100000, 10, 100000},
		{0, 2, 0},
		{0, 0, 0},
	}
	for _, c := range cases {
		buf := GenRoutes(c.count, c.depth)
		got := strings.Count(buf.String(), "\n")
		if got != c.want {
			t.Errorf("genRoutes(%q) == %q, want %q", c.count, got, c.want)
		}
	}
}

func TestChildRoute(t *testing.T) {

	cases := []struct {
		in   int
		want string
	}{
		{3, "/basetest/path3/subpath3/subsubpath3/"},
		{2, "/basetest/path2/subpath2/"},
		{1, "/basetest/path1/"},
		{0, "/basetest/"},
	}
	for _, c := range cases {
		got := childRoute("test", c.in)
		if got != c.want {
			t.Errorf("subName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSubName(t *testing.T) {

	cases := []struct {
		in   int
		want string
	}{
		{2, "subsubpathfolder"},
		{1, "subpathfolder"},
		{0, "pathfolder"},
	}
	for _, c := range cases {
		buf := subNameBuf("folder", c.in)
		got := buf.String()
		if got != c.want {
			t.Errorf("subName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkChildRoute(b *testing.B) { // generate a text file that contains n test routes
	depth := 10
	for n := 0; n < b.N; n++ {
		childRoute("test", depth)
	}
}
func BenchmarkSubNameBuf(b *testing.B) { // generate a text file that contains n test routes
	depth := 10
	for n := 0; n < b.N; n++ {
		subNameBuf("0", depth)
	}
}
