package cryptos

// 依存性逆転をここで利用する。
type IAESService interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

type AESService struct {
	AESExecuter IAESService
}

func NewAESService(service IAESService) *AESService {
	return &AESService{
		AESExecuter: service,
	}
}

func (s *AESService) Encrypt(plainText string) (string, error) {
	return s.AESExecuter.Encrypt(plainText)
}

func (s *AESService) Decrypt(cipherText string) (string, error) {
	return s.AESExecuter.Decrypt(cipherText)
}