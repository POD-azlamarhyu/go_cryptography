package cryptos

// 依存性逆転をここで利用する。
type IDESService interface {
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}

type DESService struct {
	DESExecuter IDESService
}

func NewDESService(service IDESService) *DESService {
	return &DESService{
		DESExecuter: service,
	}
}

func (s *DESService) Encrypt(plainText []byte) ([]byte, error) {
	return s.DESExecuter.Encrypt(plainText)
}

func (s *DESService) Decrypt(cipherText []byte) ([]byte, error) {
	return s.DESExecuter.Decrypt(cipherText)
}