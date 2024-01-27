package customerRoutes

import (
	customerRoute "github.com/elhardian/go-clean-architecture/api/customer/delivery/route"
	"github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/gorilla/mux"
)

func NewRoutes(r *mux.Router, mgr manager.Manager) {
	api := r.PathPrefix("/customers").Subrouter()

	customerRoute.NewCustomerRoute(mgr, api)
}
