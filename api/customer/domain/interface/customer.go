package customerDomainInterface

import (
	"context"
	"net/http"

	customerDomainEntity "github.com/elhardian/go-clean-architecture/api/customer/domain/entity"
)

type CustomerHandler interface {
	GetCustomers() http.Handler
}

type CustomerUsecase interface {
	GetCustomers(ctx context.Context) ([]*customerDomainEntity.Customer, error)
}

type CustomerRepository interface {
	GetCustomers(ctx context.Context) ([]*customerDomainEntity.Customer, error)
}
