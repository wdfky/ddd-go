package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wdfky/ddd-go/aggregate"
	"github.com/wdfky/ddd-go/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type test struct {
		id    uuid.UUID
		want  aggregate.Customer
		want1 error
	}
	cust, err := aggregate.NewCustomer("wdf", 0)
	if err != nil {
		t.Errorf("newCustomer should return nil")
	}
	id := cust.GetID()
	repo := &MemoryRepository{
		customers: map[string]*aggregate.Customer{
			id.String(): &cust,
		},
	}
	tests := []test{}
	tests = append(tests, test{
		id:    id,
		want:  cust,
		want1: nil,
	})
	tests = append(tests, test{
		id:    uuid.New(),
		want:  aggregate.Customer{},
		want1: customer.ErrCustomerNotFound,
	})
	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			got, got1 := repo.FindByID(tt.id)
			if got1 == nil && got.GetName() != tt.want.GetName() {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCustomer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestMemory_AddCustomer(t *testing.T) {
	type test struct {
		cust aggregate.Customer
		want error
	}
	cust, err := aggregate.NewCustomer("wdf", 0)
	if err != nil {
		t.Errorf("newCustomer should return nil")
	}
	repo := &MemoryRepository{
		customers: map[string]*aggregate.Customer{},
	}
	tests := []test{}
	tests = append(tests, test{
		cust: cust,
		want: nil,
	})
	for _, tt := range tests {
		t.Run(tt.cust.GetID().String(), func(t *testing.T) {
			got := repo.Add(tt.cust)
			if got != tt.want {
				t.Errorf("AddCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMemory_UpdateCustomer(t *testing.T) {
	type test struct {
		cust aggregate.Customer
		want error
	}
	cust, err := aggregate.NewCustomer("wdf", 0)
	if err != nil {
		t.Errorf("newCustomer should return nil")
	}
	repo := &MemoryRepository{
		customers: map[string]*aggregate.Customer{
			cust.GetID().String(): &cust,
		},
	}
	tests := []test{}
	tests = append(tests, test{
		cust: cust,
		want: nil,
	})
	for _, tt := range tests {
		t.Run(tt.cust.GetID().String(), func(t *testing.T) {
			got := repo.Update(tt.cust)
			if got != tt.want {
				t.Errorf("UpdateCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}

}
