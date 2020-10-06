package serializer

import (
	"encoding/json"

	"github.com/juniorrosul/pismo/account"
)

type Account struct{}

func (a *Account) Decode(input []byte) (*account.Model, error) {
	account := &account.Model{}

	if err := json.Unmarshal(input, account); err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Encode(input *account.Model) ([]byte, error) {
	encoded, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return encoded, nil
}
