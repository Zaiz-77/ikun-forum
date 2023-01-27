package dto

import "zaizwk/ginessential/model"

type UserDto struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Ps   string `json:"ps"`
	Role string `json:"role"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{Name: user.Name, Tel: user.Tel, Ps: user.PS, Role: user.Role}
}
