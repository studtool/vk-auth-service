package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/rest"

	"github.com/studtool/vk-auth-service/beans"
	"github.com/studtool/vk-auth-service/config"
)

type Server struct {
	server *rest.Server
}

func NewServer() *Server {
	srv := &Server{
		server: rest.NewServer(
			rest.ServerConfig{
				Host: consts.EmptyString,
				Port: config.ServerPort.Value(),
			},
		),
	}
	srv.server.SetLogger(beans.Logger)

	mx := mux.NewRouter()

	h := srv.server.WithRecover(mx)
	if config.ShouldLogRequests.Value() {
		h = srv.server.WithLogs(h)
	}
	if config.ShouldAllowCORS.Value() {
		h = srv.server.WithCORS(h, rest.CORS{
			Origins: []string{"*"},
			Methods: []string{
				http.MethodGet, http.MethodHead,
				http.MethodPost, http.MethodPatch,
				http.MethodDelete, http.MethodOptions,
			},
			Headers: []string{
				"User-Agent", "Authorization",
				"Content-Type", "Content-Length",
			},
			Credentials: false,
		})
	}

	srv.server.SetHandler(h)

	return srv
}

func (srv *Server) Run() error {
	return srv.server.Run()
}

func (srv *Server) Shutdown() error {
	return srv.server.Shutdown()
}
