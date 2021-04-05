package users

import "encoding/json"

type PublicUer struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPulbic bool) interface{} {
	if isPulbic {
		return PublicUer{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)

	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}

func (users Users) Marshall(isPulbic bool) interface{} {
	result := make([]interface{}, len(users))

	for i, user := range users {
		result[i] = user.Marshall(isPulbic)
	}
	return result
}
