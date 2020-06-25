package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RechargeAPI struct {
	router *mux.Router
}

func NewRechargeAPI() *RechargeAPI {
	api := RechargeAPI{router: mux.NewRouter()}
	api.initRoutes()
	return &api
}

func (api *RechargeAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}
