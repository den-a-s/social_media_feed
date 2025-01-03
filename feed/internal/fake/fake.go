package fake

const (
	AdminId = 1
)

func IsAdmin(userId int) bool {
	return AdminId == userId
}