package services

import (
	"realtime-score/internal/dto"
	"realtime-score/internal/models"
	"realtime-score/internal/pkg"
	"realtime-score/internal/repository"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repository.UserRepository
	s3aws    *pkg.S3aws
}

func NewUser(userRepo *repository.UserRepository, s3aws *pkg.S3aws) *UserService {
	return &UserService{
		userRepo: userRepo,
		s3aws:    s3aws,
	}
}

func (s *UserService) Register(req dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	isEmailExist := s.userRepo.IsEmailExist(req.Email)
	if isEmailExist {
		return dto.UserCreateResponse{}, dto.EmailAlreadyExist
	}

	isUsernameExist := s.userRepo.IsUsernameExist(req.Username)
	if isUsernameExist {
		return dto.UserCreateResponse{}, dto.UsernameAlreadyExist
	}
	hashedPassword, err := pkg.HashPassword(req.Password)
	user_id := uuid.New()

	allowedPhotoType := []string{"image/jpeg", "image/png"}
	linkProfilePict := ""
	if req.ProfilePicture != nil {
		objectKey, err := s.s3aws.FileUpload("PICT-"+user_id.String(), req.ProfilePicture, "ProfilePicture", allowedPhotoType...)
		if err != nil {
			return dto.UserCreateResponse{}, err
		}
		linkProfilePict = s.s3aws.GetPublicLink(objectKey)
	}
	data := models.User{
		ID:             user_id,
		Username:       req.Username,
		Email:          req.Email,
		Password:       hashedPassword,
		ProfilePicture: linkProfilePict,
		Role:           "user",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	register, err := s.userRepo.CreateUser(data)
	if err != nil {
		return dto.UserCreateResponse{}, dto.CantCreateUser

	}
	return dto.UserCreateResponse{
		Username:       register.Username,
		Email:          register.Email,
		ProfilePicture: register.ProfilePicture,
	}, nil
}

func (s *UserService) Login(req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	checkUserExist, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.InvalidCredentials
	}

	checkPass := pkg.CheckPassword(req.Password, checkUserExist.Password)
	if !checkPass {
		return dto.UserLoginResponse{}, dto.InvalidCredentials
	}

	token, err := pkg.GenerateToken(checkUserExist.ID, checkUserExist.Role)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{
		Token: token,
		Role:  checkUserExist.Role,
	}, nil
}

func (s *UserService) GetUserByID(userid string) (dto.GetUser, error) {
	user, err := s.userRepo.GetUserById(userid)
	if err != nil {
		return dto.GetUser{}, dto.UserNotFound
	}
	return dto.GetUser{
		ID:             user.ID.String(),
		Username:       user.Username,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Role:           user.Role,
	}, nil
}

func (s *UserService) GetAllUser() (dto.GetAllUserResponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return dto.GetAllUserResponse{}, dto.ErrGetAllUser
	}
	return users, nil
}

func (s *UserService) UpdateUser(req dto.UserUpdateRequest, id string) (dto.UserCreateResponse, error) {
	userExist, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dto.UserCreateResponse{}, dto.UserNotFound
	}

	allowedPhotoType := []string{"image/jpeg", "image/png"}

	if req.ProfilePicture != nil && userExist.ProfilePicture != "" {
		objectKey := s.s3aws.GetObjectKeyFromLink(userExist.ProfilePicture)
		objectKey, err = s.s3aws.UpdateFile(objectKey, req.ProfilePicture, allowedPhotoType...)
		if err != nil {
			return dto.UserCreateResponse{}, err
		}
		userExist.ProfilePicture = s.s3aws.GetPublicLink(objectKey)
	} else {
		objectKey, err := s.s3aws.FileUpload("PICT-"+userExist.ID.String(), req.ProfilePicture, "ProfilePicture", allowedPhotoType...)
		if err != nil {
			return dto.UserCreateResponse{}, err
		}
		userExist.ProfilePicture = s.s3aws.GetPublicLink(objectKey)
	}

	if req.Email != "" {
		emailExist := s.userRepo.IsEmailExist(req.Email)
		if emailExist {
			return dto.UserCreateResponse{}, dto.EmailAlreadyExist
		}
		userExist.Email = req.Email
	}
	if req.Username != "" {
		usernameExist := s.userRepo.IsUsernameExist(req.Username)
		if usernameExist {
			return dto.UserCreateResponse{}, dto.UsernameAlreadyExist
		}
		userExist.Username = req.Username
	}

	timeupdate := time.Now()
	userExist.UpdatedAt = timeupdate
	updatedUser, err := s.userRepo.UpdateUser(userExist)
	if err != nil {
		return dto.UserCreateResponse{}, dto.ErrUpdateUser
	}
	return dto.UserCreateResponse{
		Username:       updatedUser.Username,
		Email:          updatedUser.Email,
		ProfilePicture: updatedUser.ProfilePicture,
	}, nil
}

func (s *UserService) DeleteUser(user_id string) error {
	if err := s.userRepo.DeleteUser(user_id); err != nil {
		return err
	}
	return nil
}
