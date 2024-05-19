package usecase

import (
	// "backend/src/model"
	// "backend/src/repository"
	// "backend/src/validator"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/hato72/go_backend_hackathon/backend/src/model"
	"github.com/hato72/go_backend_hackathon/backend/src/repository"
	"github.com/hato72/go_backend_hackathon/backend/src/validator"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	Update(user model.User, newEmail string, newName string, newPassword string, iconFile *multipart.FileHeader) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:      newUser.ID,
		Name:    newUser.Name,
		Email:   newUser.Email,
		IconUrl: newUser.IconUrl,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{} //空のユーザーオブジェクト
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) //パスワードの検証
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), //jwtの有効期限
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) //jwtトークンの生成
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) Update(user model.User, newEmail string, newName string, newPassword string, iconFile *multipart.FileHeader) (model.UserResponse, error) {

	if newPassword != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return model.UserResponse{}, err
		}
		user.Password = string(hash)
	}

	if iconFile != nil {
		src, err := iconFile.Open()
		if err != nil {
			return model.UserResponse{}, err
		}
		defer src.Close()

		data, err := io.ReadAll(src)
		if err != nil {
			return model.UserResponse{}, err
		}

		hasher := sha256.New()
		hasher.Write(data)
		hashValue := hex.EncodeToString(hasher.Sum(nil))

		ext := filepath.Ext(iconFile.Filename)

		iconUrl := "icons/" + hashValue + ext

		dst, err := os.Create("public/user_images/" + iconUrl)
		if err != nil {
			return model.UserResponse{}, err
		}

		defer dst.Close()

		if _, err := dst.Write(data); err != nil {
			return model.UserResponse{}, nil
		}

		user.IconUrl = &iconUrl

	}

	updatedUser := model.User{
		ID:       user.ID,
		Name:     newName,
		Email:    newEmail,
		Password: newPassword,
		IconUrl:  user.IconUrl,
	}
	//log.Print("updateUser:", updatedUser)

	if err := uu.ur.UpdateUser(&updatedUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID:      updatedUser.ID,
		Name:    updatedUser.Name,
		Email:   updatedUser.Email,
		IconUrl: updatedUser.IconUrl,
	}
	//log.Print("resUser:", resUser)

	return resUser, nil

}
