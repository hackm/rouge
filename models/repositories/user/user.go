package user

type User map[string]interface{}

type IUserRepository interface {
  GetUser(id int64) User
}

type UserRepository struct {}

func NewUserRepository() UserRepository {
  return UserRepository{}
}

func (r UserRepository) GetUser(id int64) User {
  u := make(User)
  u["name"] = "hackm"
  u["displayName"] = "Hackathon Monster"
  
  return u
}