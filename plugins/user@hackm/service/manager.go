package service

type Manager struct {

}

func NewUserManager() *Manager {
	return &Manager{}
}

func (*Manager) GetIdByName(name string) int64 {
	return 1000
}

func (*Manager) ValidateUser(username, password string) bool {
	return password == "1234"
}
