package repository

import (
	"rest-api-go/internal/dto"
	"rest-api-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) IsEmailExist(email string) bool {
	var count int64
	if err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
func (r *UserRepository) IsUserExist(id string) bool {
	var count int64
	if err := r.db.Model(&models.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (r *UserRepository) IsUsernameExist(username string) bool {
	var count int64
	if err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(reqUser models.User) (models.User, error) {
	if err := r.db.Create(&reqUser).Error; err != nil {
		return models.User{}, err
	}
	return reqUser, nil
}

func (r *UserRepository) GetUserById(user_id string) (models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUser() (dto.GetAllUserResponse, error) {
	var users []models.User

	if err := r.db.Model(&models.User{}).Find(&users).Error; err != nil {
		return dto.GetAllUserResponse{}, err
	}

	var userResponses []dto.GetUser
	for _, user := range users {
		userResponses = append(userResponses, dto.GetUser{
			ID:             user.ID.String(),
			Username:       user.Username,
			Email:          user.Email,
			ProfilePicture: user.ProfilePicture,
			Role:           user.Role,
		})
	}

	return dto.GetAllUserResponse{
		Users: userResponses,
	}, nil
}

func (r *UserRepository) UpdateUser(req models.User) (models.User, error) {
	if err := r.db.Updates(&req).Error; err != nil {
		return models.User{}, err
	}
	return req, nil
}

func (r *UserRepository) DeleteUser(user_id string) error {
	var user models.User
	if err := r.db.Where("id = ?", user_id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
