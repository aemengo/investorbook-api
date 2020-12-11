package server

import (
	"github.com/aemengo/investorbook-api/controller"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	con *controller.Controller
}

func New(con *controller.Controller) *server{
	return &server{
		con: con,
	}
}

func (s *server) Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/investors/:investorId/connections", s.connections)
	router.GET("/investors/:investorId/mutual/:otherInvestorId", s.mutualConnections)
	router.GET("/investors/:investorId/search", s.search)
	return router
}