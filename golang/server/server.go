package server

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/aituglo/rubyx-dashboard/golang/env"
	"github.com/aituglo/rubyx-dashboard/golang/errors"
	"github.com/aituglo/rubyx-dashboard/golang/server/write"
)

// isDev is used in the lag middleware... we don't want to read from env on every request
var isDev = false

func init() {
	isDev = os.Getenv("ENV") == "dev"
}

type server struct {
	env    env.Env
	router *httprouter.Router
}

// New initializes env (database connections and whatnot)
// and creates a server that implements ServeHTTP
func New() (*server, env.Env, error) {
	env, err := env.New()
	if err != nil {
		return nil, nil, err
	}

	srv := &server{
		env: env,
	}

	srv.ConfigureRouter()

	return srv, env, nil
}

// ServeHTTP handles routes that start with api
func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	// shift head and tail to get below "api/" part of the path
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head != "api" {
		write.Error(errors.RouteNotFound)
	}

	srv.router.ServeHTTP(w, r)
}

func (srv *server) Close() {
	srv.env.Close()
}
