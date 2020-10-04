package operationtype

type operationTypeService struct {
	operationTypeRepository Repository
}

func newOperationTypeService(operationTypeRepository Repository) Service {
	return &operationTypeService{
		operationTypeRepository,
	}
}

func (ot *operationTypeService) Find(id int) (*OperationType, error) {
	return ot.operationTypeRepository.Find(id)
}
