package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/crypto_utils"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
)

var (
	UserService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindByStatus(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Vaildate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if len(user.FirstName) != 0 {
		current.FirstName = user.FirstName
	}

	if len(user.LastName) != 0 {
		current.LastName = user.LastName
	}

	if len(user.Email) != 0 {
		current.Email = user.Email
	}

	if len(user.Password) != 0 {
		current.Password = user.Password
	}

	if len(user.Status) != 0 {
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	return user.Delete()
}

func (s *usersService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	user := &users.User{Status: status}
	return user.FindByStatus()

}
