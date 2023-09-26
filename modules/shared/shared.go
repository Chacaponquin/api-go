package shared

type Service interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hashedPassword string) bool
}
