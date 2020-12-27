package router

import (
	"io"
	"log"
	"net/http"
	"os"
)

type baseHandler interface{}

type Router struct {
	NotFoundHandler http.Handler
	Routing         Routing
	routes          []*Route
	outLog          *log.Logger
	errLog          *log.Logger
}

// Routing is represents routing tree
type Routing interface {
	Lookup(method, path string) (HandlerData, error)
	Insert(method, path string, handler baseHandler) error
}

// HandlerData is represents handler function and args
type HandlerData struct {
	handler baseHandler
	params  []interface{}
}

// Get register handler via GET
func (r *Router) Get(path string, h baseHandler) *Route { return r.HandleFunc("GET", path, h) }

// Head register handler via HEAD
func (r *Router) Head(path string, h baseHandler) *Route { return r.HandleFunc("HEAD", path, h) }

// Post register handler via POST
func (r *Router) Post(path string, h baseHandler) *Route { return r.HandleFunc("POST", path, h) }

// Put register handler via PUT
func (r *Router) Put(path string, h baseHandler) *Route { return r.HandleFunc("PUT", path, h) }

// Patch register handler via PATCH
func (r *Router) Patch(path string, h baseHandler) *Route { return r.HandleFunc("PATCH", path, h) }

// Delete register handler via DELETE
func (r *Router) Delete(path string, h baseHandler) *Route { return r.HandleFunc("DELETE", path, h) }

// Options register handler via OPTIONS
func (r *Router) Options(path string, h baseHandler) *Route { return r.HandleFunc("OPTIONS", path, h) }

func (r *Router) errorLogf(args ...interface{}) {
	if env := os.Getenv("GO_ROUTER_ENABLE_LOGGING"); len(env) != 0 {
		r.errLog.Printf("[error] %s\n", args)
	}
}


func newLogger(w io.Writer) *log.Logger {
	return log.New(w, "", log.LstdFlags|log.Lshortfile)
}


// NewRouter return created Router
func NewRouter() *Router {
	return &Router{
		NotFoundHandler: http.NotFoundHandler(),
		Routing:         NewTrie(),
		outLog:          newLogger(os.Stdout),
		errLog:          newLogger(os.Stderr),
	}
}

// AddRoute add route in router
func (r *Router) AddRoute() *Route {
	route := &Route{}
	r.routes = append(r.routes, route)
	return route
}

// HandleFunc register handler each HTTP method
func (r *Router) HandleFunc(method, path string, h baseHandler) *Route {
	route := r.AddRoute().HandleFunc(method, path, h)
	err := r.Routing.Insert(route.method, route.path, route.handler)
	if err != nil {
		r.errorLogf("failed registered path. path=%s, error=%v", path, err)
	}
	return route
}
