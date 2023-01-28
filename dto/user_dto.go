package dto

import "zaizwk/ginessential/model"

type UserDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Ps   string `json:"ps"`
	Role string `json:"role"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{ID: user.ID, Name: user.Name, Tel: user.Tel, Ps: user.PS, Role: user.Role}
}
