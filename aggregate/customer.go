package aggregate

import (
	"github.com/wdfky/ddd-go/domain/entity"
	"github.com/wdfky/ddd-go/domain/valueobject"
)

type Customer struct {
	Person       *entity.Person             `json:"person" bson:"person"`
	Items        []*entity.Item             `json:"items" bson:"items"`
	Transactions []*valueobject.Transaction `json:"transactions" bson:"transactions"`
}
