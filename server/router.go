package server

import (
	"net/http"
)

// The Router handler
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// Create a Router instance
func BuildRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Finds the assigned handler for the provided path
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExists := r.rules[path]
	if !pathExists {
		return nil, pathExists, false
	}
	handler, methodExists := r.rules[path][method]
	return handler, pathExists, methodExists
}

// Checks the request path against the registered routes. If the path
// does not exist, returns a 404 (Not Found). Otherwise, returns the
// handler assigned to that path.
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, pathExists, methodExists := r.FindHandler(request.URL.Path, request.Method)
	if !pathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
