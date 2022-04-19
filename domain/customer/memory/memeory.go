//纯内存，无数据库
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/wdfky/ddd-go/aggregate"
	"github.com/wdfky/ddd-go/domain/customer"
)

type MemoryRepository struct {
	customers map[string]*aggregate.Customer
	sync.Mutex
}

func NewCustomerMemory() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[string]*aggregate.Customer),
	}
}
func (r *MemoryRepository) Update(c aggregate.Customer) error {
	r.Lock()
	defer r.Unlock()
	//一系列数据操作，可能
	r.customers[c.GetID().String()] = &c
	return nil
}
func (r *MemoryRepository) Add(c aggregate.Customer) error {
	if r == nil {
		r.Lock()
		r = NewCustomerMemory()
		r.Unlock()
	}
	r.Lock()
	defer r.Unlock()
	if _, ok := r.customers[c.GetID().String()]; ok {
		return fmt.Errorf("custome is already exisit:%w", customer.ErrFailToAddCustomer)
	}
	r.customers[c.GetID().String()] = &c
	return nil
}
func (r *MemoryRepository) FindByID(id uuid.UUID) (aggregate.Customer, error) {
	r.Lock()
	defer r.Unlock()
	if c, ok := r.customers[id.String()]; ok {
		return *c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}
