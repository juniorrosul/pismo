package account

// Repository interface
type Repository interface {
	Find(id int) (*Model, error)
	Store(account *Model) error
}
