package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wdfky/ddd-go/aggregate"
)

var (
	ErrFailToUpdateCustomer = errors.New("update Customer failed")
	ErrFailToAddCustomer    = errors.New("fail to add customer")
	ErrCustomerNotFound     = errors.New("fail to find customer")
)

type CustomerRepository interface {
	Update(aggregate.Customer) error
	Add(aggregate.Customer) error
	FindByID(id uuid.UUID) (aggregate.Customer, error)
}
