package customerRepository

import (
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	customerDomainEntity "github.com/elhardian/go-clean-architecture/api/customer/domain/entity"
	customerDomainInterface "github.com/elhardian/go-clean-architecture/api/customer/domain/interface"
)

type customer struct {
	DB *gorm.DB
}

func NewCustomerRepository(database *gorm.DB) customerDomainInterface.CustomerRepository {
	repo := new(customer)
	repo.DB = database

	return repo
}

func (repo *customer) GetCustomers(ctx context.Context) ([]*customerDomainEntity.Customer, error) {
	code := ""
	results := []*customerDomainEntity.Customer{}
	err := repo.DB.
		Select("user_id", "email", "email", "password", "name").
		Find(&results).Error
	if err != nil {
		code = "[Repository] GetCustomers-1 = "
		log.Error().Err(err).Msg(code + err.Error())
		return nil, err
	}

	return results, nil
}
