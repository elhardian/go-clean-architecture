package customerHandler

import (
	"net/http"

	response "github.com/elhardian/go-clean-architecture/helper/response"
	manager "github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/rs/zerolog/log"

	customerDomainInterface "github.com/elhardian/go-clean-architecture/api/customer/domain/interface"
	customerUsecase "github.com/elhardian/go-clean-architecture/api/customer/usecase"
)

type Customer struct {
	Usecase customerDomainInterface.CustomerUsecase
	mgr     manager.Manager
}

func NewCustomerHandler(mgr manager.Manager) customerDomainInterface.CustomerHandler {
	handler := new(Customer)
	handler.Usecase = customerUsecase.NewCustomerUsecase(mgr)
	handler.mgr = mgr

	return handler
}

func (h *Customer) GetCustomers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		code := ""
		results, err := h.Usecase.GetCustomers(ctx)
		if err != nil {
			code = "[Handler] GetCustomers-1 = "
			log.Error().Err(err).Msg(code + err.Error())
			response.OnError(w, false, http.StatusBadRequest, err.Error(), code)
			return
		}

		response.OnSuccess(w, true, "", results, nil)
	})
}
