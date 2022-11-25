package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

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
		err := errors.New("path " + request.URL.Path + " is not registered")
		response, _ := json.Marshal(BuildErrorResponse(http.StatusNotFound, err))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
		LogError(http.StatusNotFound, request.URL.Path, err)
		return
	}

	if !methodExists {
		err := errors.New("no handler for the " + request.Method + " method in the " + request.URL.Path + " path")
		response, _ := json.Marshal(BuildErrorResponse(http.StatusMethodNotAllowed, err))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response)
		LogError(http.StatusMethodNotAllowed, request.URL.Path, err)
		return
	}
	handler(w, request)
}
