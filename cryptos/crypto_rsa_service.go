package cryptos

type IRSAService interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

type RSAService struct {
	rsaExecuter IRSAService
}

func NewRSAService(rsaExecuter IRSAService) *RSAService {
	return &RSAService{
		rsaExecuter: rsaExecuter,
	}
}

func (s *RSAService) Encrypt(plainText string) (string, error) {
	return s.rsaExecuter.Encrypt(plainText)
}

func (s *RSAService) Decrypt(cipherText string) (string, error) {
	return s.rsaExecuter.Decrypt(cipherText)
}