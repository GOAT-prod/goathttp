package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request)

type Router struct {
	router *mux.Router
}

type Server struct {
	ctx    context.Context
	logger goatlogger.Logger
	server *http.Server
}

type HandlerDefinition struct {
	Path    string
	Method  string
	Handler Handler
}

func New(ctx context.Context, router Router, port int, appName string) *Server {
	return &Server{
		ctx:    ctx,
		logger: goatlogger.New(appName),
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
			Handler: router.router,
		},
	}
}

func NewRouter() Router {
	return Router{
		router: mux.NewRouter(),
	}
}

func (r *Router) Get(path string, handler Handler) {
	r.router.HandleFunc(path, handler).Methods(http.MethodGet)
}

func (r *Router) Post(path string, handler Handler) {
	r.router.HandleFunc(path, handler).Methods(http.MethodPost)
}

func (r *Router) Put(path string, handler Handler) {
	r.router.HandleFunc(path, handler).Methods(http.MethodPut)
}

func (r *Router) Delete(path string, handler Handler) {
	r.router.HandleFunc(path, handler).Methods(http.MethodDelete)
}

func (r *Router) NewSubRouter(pathPrefix string, handlers []HandlerDefinition, middlewares []mux.MiddlewareFunc) {
	s := r.router.PathPrefix(pathPrefix).Subrouter()
	s.Use(middlewares...)

	for _, handler := range handlers {
		s.HandleFunc(handler.Path, handler.Handler).Methods(handler.Method)
	}
}

func (s *Server) Start(router *mux.Router) error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Shutdown(s.ctx)
}

func (s *Server) Setup(router Router) {
	s.server.Handler = router.router
}

func (r *Router) SetMiddlewares(middleware ...mux.MiddlewareFunc) {
	r.router.Use(middleware...)
}
