package data

import "Test/domain"

type User struct {
	Username    string `json:"username" form:"username" validate:"required"`
	Fullname    string `json:"fullname" form:"fullname"`
	Email       string `json:"email" form:"email"`
	City        string `json:"city" form:"city"`
	Phonenumber string `json:"phonenumber" form:"phonenumber"`
	Password    string `json:"password" form:"password" validate:"required"`
	Points      int
	Balance     int
	Role        string
}

func (u *User) ToModel() domain.User {
	return domain.User{
		Username:    u.Username,
		Email:       u.Email,
		City:        u.City,
		Phonenumber: u.Phonenumber,
		Password:    u.Password,
		Points:      u.Points,
		Balance:     u.Balance,
		Role:        u.Role,
	}
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.Username
	res.Email = data.Email
	res.City = data.City
	res.Phonenumber = data.Phonenumber
	res.Password = data.Password
	res.Points = data.Points
	res.Balance = data.Balance
	res.Role = data.Role
	return res
}
