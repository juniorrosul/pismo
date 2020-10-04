package operationtype

// Repository interface
type Repository interface {
	Find(id int) (*Service, error)
	Store(operationtype *OperationType) error
}
