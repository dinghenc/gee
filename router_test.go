package gee

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assert/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	cases := []struct {
		pattern string
		parts   []string
	}{
		{
			pattern: "/p/:name",
			parts:   []string{"p", ":name"},
		}, {
			pattern: "/p/*/",
			parts:   []string{"p", "*"},
		}, {
			pattern: "/p/*name/*",
			parts:   []string{"p", "*name"},
		},
	}

	for i, c := range cases {
		ok := reflect.DeepEqual(parsePattern(c.pattern), c.parts)
		if !ok {
			t.Logf("%d case failed: pattern = %s, parts = %v, calculate_parts = %v\n", i, c.pattern, c.parts, parsePattern(c.pattern))
			continue
		}
		t.Logf("%d/%d case pass\n", i+1, len(cases))
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()

	cases := []struct {
		path    string
		pattern string
		params  map[string]string
	}{
		{
			path:    "/hello/huster",
			pattern: "/hello/:name",
			params:  map[string]string{"name": "huster"},
		}, {
			path:    "/",
			pattern: "/",
			params:  map[string]string{},
		}, {
			path:    "/hello/b/c",
			pattern: "/hello/b/c",
			params:  map[string]string{},
		}, {
			path:    "/assert/image/hust.jpg",
			pattern: "/assert/*filepath",
			params:  map[string]string{"filepath": "image/hust.jpg"},
		},
	}

	for i, c := range cases {
		n, params := r.getRoute("GET", c.path)
		if n == nil {
			t.Logf("%d get route is nil: path = %s\n", i, c.path)
			continue
		}
		if n.pattern != c.pattern || !reflect.DeepEqual(params, c.params) {
			t.Logf("%d get route pattern error: path = %s\n, pattern = (%s, %s), params = (%v, %v)\n", i, c.path, c.pattern, n.pattern, c.params, params)
			continue
		}
		t.Logf("%d/%d case pass\n", i+1, len(cases))
	}
}
