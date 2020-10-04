package serializer

import (
	"encoding/json"

	"madsonjr.com/pismo/account"
)

type Account struct{}

func (a *Account) Decode(input []byte) (*account.Account, error) {
	account := &account.Account{}

	if err := json.Unmarshal(input, account); err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Encode(input *account.Account) ([]byte, error) {
	encoded, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return encoded, nil
}
