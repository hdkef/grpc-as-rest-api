package usecase

type Bcrypt_ interface {
	CompareHashAndPassword([]byte, []byte) error
	GenerateFromPassword([]byte, int) ([]byte, error)
}
