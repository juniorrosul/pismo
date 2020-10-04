package account

// Service interface
type Service interface {
	Find(id int) (*Account, error)
	Store(account *Account) error
}
