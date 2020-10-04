package operationtype

// Repository interface
type Repository interface {
	Find(id int) (*OperationType, error)
}
