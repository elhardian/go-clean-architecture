package customerRoute

import (
	customerHandler "github.com/elhardian/go-clean-architecture/api/customer/delivery/handler"
	"github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/gorilla/mux"
)

func NewCustomerRoute(mgr manager.Manager, route *mux.Router) {
	customerHandler := customerHandler.NewCustomerHandler(mgr)

	route.Handle("", customerHandler.GetCustomers()).Methods("GET")
}
