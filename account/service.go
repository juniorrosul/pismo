package account

// Service interface
type Service interface {
	Find(id int) (*Model, error)
	Store(account *Model) error
}
