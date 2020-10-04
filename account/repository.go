package account

// Repository interface
type Repository interface {
	Find(id int) (*Account, error)
	Store(account *Account) error
}
