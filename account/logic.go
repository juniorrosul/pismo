package account

import "errors"

var (
	ErrAccountNotFound     = errors.New("Account not found")
	ErrAccountStoreInvalid = errors.New("Account not stored")
)

type accountService struct {
	accountRepository Repository
}

func (a *accountService) Find(id int) (*Model, error) {
	return a.accountRepository.Find(id)
}

func (a *accountService) Store(account *Model) error {
	return a.accountRepository.Store(account)
}
