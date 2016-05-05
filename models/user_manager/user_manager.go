package user_manager

type UserManager struct{
  
}

func NewUserManager() *UserManager{
  return &UserManager{}
}

func (*UserManager) GetIdByName(string name) int64 {
  return 1000
} 

func (*UserManager) ValidateUser(string username, string password) bool{
  return password == "1234"
} 