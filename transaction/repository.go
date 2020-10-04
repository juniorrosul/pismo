package transaction

// Repository interface
type Repository interface {
	Find(id int) (*Transaction, error)
	Store(transaction *Transaction) error
}
