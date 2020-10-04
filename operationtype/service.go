package operationtype

// Service interface
type Service interface {
	Find(id int) (*OperationType, error)
	Store(operationType *OperationType) error
}
