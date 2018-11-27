package routs

import (
	"net/http"
	"regexp"
	)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handle(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}
func (h *RegexpHandler) DefaultHandle(pattern string, handler http.Handler) {
	h.routes = append(h.routes, &route{regexp.MustCompile(pattern), handler})
}
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}


func (h *RegexpHandler) DefaultHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {

	h.routes = append(h.routes, &route{regexp.MustCompile(pattern), http.HandlerFunc(handler)})
}
