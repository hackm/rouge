package services

type UserManager struct {

}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (*UserManager) GetIdByName(name string) int64 {
	return 1000
}

func (*UserManager) ValidateUser(username, password string) bool {
	return password == "1234"
}
