package user

type Repository interface {
	getAllUsers() map[string]string
}