package server

import (
	"agency-banking/pkg/token"
	"agency-banking/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// ///// pkg
type Server struct {
	Config     util.Config
	Router     *gin.Engine
	TokenMaker token.Maker
	ChiRouter  *chi.Mux
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	chiRouter := chi.NewRouter()

	return &Server{Config: config, TokenMaker: tokenMaker, ChiRouter: chiRouter}, nil
}

func (s *Server) StartChi(address string) {
	s.setupRoutes()
	http.Handle("/", s.ChiRouter)
	http.ListenAndServe(address, nil)
}

func (s *Server) StartGin(address string) error {
	s.Router = gin.Default()
	return s.Router.Run(address)

}

//// here

func (s *Server) setupRoutes() {
	s.ChiRouter.Get("/doc/swagger", httpSwagger.WrapHandler)
	s.ChiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("agency banking api running at : %s", s.Config.HTTPServerAddress)))
	})
}
