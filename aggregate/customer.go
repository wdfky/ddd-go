package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wdfky/ddd-go/entity"
	"github.com/wdfky/ddd-go/valueobject"
)

var (
	ErrInvalidCustomer = errors.New("a customer must have name")
)

type Customer struct {
	Person       *entity.Person            `json:"person" bson:"person"`
	Products     []*entity.Item            `json:"items" bson:"items"`
	Transactions []valueobject.Transaction `json:"transactions" bson:"transactions"`
}

func NewCustomer(name string, age int) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidCustomer
	}
	return Customer{
		Person: &entity.Person{
			Id:   uuid.New(),
			Name: name,
			Age:  age,
		},
	}, nil
}
