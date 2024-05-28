package user

const (
	MARITAL_STATUS_NOT_SET = -1
	MARITAL_STATUS_FALSE   = 0
	MARITAL_STATUS_TRUE    = 1
)

type User struct {
	Id      uint32 // Unique identifier
	FName   string // First name
	City    string // City of residence
	Phone   uint32 // 10 digit phone number
	Height  uint8  // in inches
	Married int8   // Marital status -1 for not set, 0 for false, 1 for true
}

// UserStore is an interface that defines the methods that must be implemented by a user store
// Implementations can be for in-memory, database, or any other type of store
type UserStore interface {
	GetUserById(id uint32) (User, error)
	GetUsersListByIds(ids []uint32) ([]User, error)
	DeleteUser(id uint32) error
	UpdateUser(user User) error
	CreateUser(user User) error
	Search(query User) ([]User, error)
}
