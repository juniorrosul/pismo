package serializer

import (
	"encoding/json"

	"madsonjr.com/pismo/transaction"
)

type Transaction struct{}

func (t *Transaction) Decode(input []byte) (*transaction.Model, error) {
	tr := &transaction.Model{}

	if err := json.Unmarshal(input, tr); err != nil {
		return nil, err
	}

	return tr, nil
}
