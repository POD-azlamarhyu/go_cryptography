package cryptos

type IHashService interface {
	ExecuteSHA256(message string) ([]byte, error)
	ExecuteSHA512(message string) ([]byte, error)
}

type HashService struct {
	executer IHashService
}

func NewHashService(hashExecuter IHashService) *HashService {
	return &HashService{
		executer: hashExecuter,
	}
}

func (s *HashService) ExecuteSHA256(message string) ([]byte, error) {
	return s.executer.ExecuteSHA256(message)
}

func (s *HashService) ExecuteSHA512(message string) ([]byte, error) {
	return s.executer.ExecuteSHA512(message)
}