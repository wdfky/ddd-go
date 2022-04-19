package valueobject

import "github.com/google/uuid"

type Transaction struct {
	Amount  int       `json:"amount" bson:"amount"`
	From    uuid.UUID `json:"from" bson:"from"`
	To      uuid.UUID `json:"to" bson:"to"`
	CreatAt int64     `json:"creatAt" bson:"creatAt"`
}
