package dto

import (
	"errors"
	"mime/multipart"
)

var (
	EmailAlreadyExist    = errors.New("email already exist")
	UsernameAlreadyExist = errors.New("username already exist")
	CantCreateUser       = errors.New("cant create user")
	InvalidCredentials   = errors.New("wrong password or email")
	UserNotFound         = errors.New("user not found")
	ErrGetAllUser        = errors.New("error get all user")
	ErrUploadFile        = errors.New("error upload file")
	ErrUpdateUser        = errors.New("error update user")
)

type (
	UserCreateRequest struct {
		Username       string                `json:"username" form:"username" binding:"required"`
		Email          string                `json:"email" form:"email" binding:"required"`
		Password       string                `json:"password" form:"password" binding:"required"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture" binding:"omitempty"`
	}
	UserCreateResponse struct {
		Username       string `json:"username"`
		Email          string `json:"email"`
		ProfilePicture string `json:"profile_picture"`
	}
	GetUser struct {
		ID             string `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		ProfilePicture string `json:"profile_picture"`
		Role           string `json:"role"`
	}
	UserLoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}
	GetAllUserResponse struct {
		Users []GetUser `json:"users"`
	}
	UserUpdateRequest struct {
		Username       string                `json:"username" form:"username"`
		Email          string                `json:"email" form:"email"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture" binding:"omitempty"`
	}
)
