package customerUsecase

import (
	"context"

	"github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/rs/zerolog/log"

	customerDomainEntity "github.com/elhardian/go-clean-architecture/api/customer/domain/entity"
	customerDomainInterface "github.com/elhardian/go-clean-architecture/api/customer/domain/interface"
	customerRepository "github.com/elhardian/go-clean-architecture/api/customer/repository"
)

type Customer struct {
	repo customerDomainInterface.CustomerRepository
	mgr  manager.Manager
}

func NewCustomerUsecase(mgr manager.Manager) customerDomainInterface.CustomerUsecase {
	usecase := new(Customer)
	usecase.repo = customerRepository.NewCustomerRepository(mgr.GetGorm())
	usecase.mgr = mgr
	return usecase
}

func (uc *Customer) GetCustomers(ctx context.Context) ([]*customerDomainEntity.Customer, error) {
	code := ""
	results, err := uc.repo.GetCustomers(ctx)
	if err != nil {
		code = "[Usecase] GetCustomers-1 = "
		log.Error().Err(err).Msg(code + err.Error())
		return results, err
	}

	return results, nil
}
