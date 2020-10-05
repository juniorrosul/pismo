package transaction

// Repository interface
type Repository interface {
	Store(transaction *Model) error
}
