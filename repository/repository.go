package repository

type DatabaseRepo interface {
	Close() error
	CreatedTableUsers() error
	VerifyUserPass(user, pass string) error
}
