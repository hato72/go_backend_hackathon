package repository

import (
	//"backend/src/model"
	"github.com/hato72/go_backend_hackathon/backend/src/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil { //emailが引数でうけとった値に一致するユーザーを探す
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil { //ユーザーの作成
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *model.User) error {
	if user.Email != "" {
		if err := ur.db.Model(user).Where("id = ?", user.ID).Update("email", user.Email).Error; err != nil {
			return err
		}
	}
	//log.Print("email", user.Email)

	if user.Name != "" {
		if err := ur.db.Model(user).Where("id = ?", user.ID).Update("name", user.Name).Error; err != nil {
			return err
		}
	}
	//log.Print("name", user.Name)

	if user.Password != "" {
		if err := ur.db.Model(user).Where("id = ?", user.ID).Update("password", user.Password).Error; err != nil {
			return err
		}
	}
	//log.Print("pass", user.Password)

	if user.IconUrl != nil {
		if err := ur.db.Model(user).Where("id = ? ", user.ID).Update("icon_url", user.IconUrl).Error; err != nil {
			return err
		}
	}
	//log.Print("icon", user.IconUrl)
	return nil
}
