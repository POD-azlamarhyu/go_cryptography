package cryptos

// 依存性逆転をここで利用する。
type IAESService interface {
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}

type AESService struct {
	AESExecuter IAESService
}

func NewAESService(service IAESService) *AESService {
	return &AESService{
		AESExecuter: service,
	}
}

func (s *AESService) Encrypt(plainText []byte) ([]byte, error) {
	return s.AESExecuter.Encrypt(plainText)
}

func (s *AESService) Decrypt(cipherText []byte) ([]byte, error) {
	return s.AESExecuter.Decrypt(cipherText)
}