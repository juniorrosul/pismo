package transaction

// Service interface
type Service interface {
	Find(id int) (*Transaction, error)
	Store(transaction *Transaction) error
}
