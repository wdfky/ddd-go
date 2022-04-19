package services

import (
	"github.com/google/uuid"
	"github.com/wdfky/ddd-go/domain/customer"
)

type OrderConfigurator func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrdeService(cfgs ...OrderConfigurator) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		if err := cfg(os); err != nil {
			return nil, err
		}
	}
	return os, nil
}
func (o *OrderService) CreateOrder(customId uuid.UUID, productIds []uuid.UUID) error {
	// cust, err := o.customers.FindByID(customId)
	// if err != nil {
	// 	return err
	// }
	// for _, pid := range productIds {
	// 	cust.AddProduct(pid)
	// }
	return nil
}
