package repository

import (
	"backend/src/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICuisineRepository interface {
	GetAllCuisines(cuisines *[]model.Cuisine, userId uint) error              //作成したタスクの一覧を取得
	GetCuisineById(cuisine *model.Cuisine, userId uint, cuisineId uint) error //引数のcuisineIdに一致するタスクを返す
	CreateCuisine(cuisine *model.Cuisine) error                               //タスクの新規作成
	UpdateCuisine(cuisine *model.Cuisine, userId uint, cuisineId uint) error  //タスクの更新
	DeleteCuisine(userId uint, cuisineId uint) error                          //タスクの削除
	UpdateCuisine_Image(cuisine *model.Cuisine, image []byte, userId uint, cuisineId uint) error
	AddURL(cuisine *model.Cuisine, userId uint, cuisineId uint) error
}

type cuisineRepository struct {
	db *gorm.DB
}

func NewCuisineRepository(db *gorm.DB) ICuisineRepository { //コンストラクタ
	return &cuisineRepository{db}
}

func (cr *cuisineRepository) GetAllCuisines(cuisines *[]model.Cuisine, userId uint) error {
	if err := cr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(cuisines).Error; err != nil { //タスクの一覧から引数のユーザーidに一致するタスクを取得する　その時、作成日時があたらしいものが末尾に来るようにする
		return err
	}
	return nil
}

func (cr *cuisineRepository) GetCuisineById(cuisine *model.Cuisine, userId uint, cuisineId uint) error {
	if err := cr.db.Joins("User").Where("user_id=?", userId).Find(cuisine, cuisineId).Error; err != nil { //引数のユーザーidに一致するタスクを取得し、その中でcuisineの主キーが引数で受け取ったcuisineIdに一致するタスクを取得する
		return err
	}
	return nil
}

func (cr *cuisineRepository) CreateCuisine(cuisine *model.Cuisine) error {
	if err := cr.db.Create(cuisine).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cuisineRepository) UpdateCuisine(cuisine *model.Cuisine, userId uint, cuisineId uint) error {
	result := cr.db.Model(cuisine).Clauses(clause.Returning{}).Where("id=? AND user_id=?", cuisineId, userId).Update("title", cuisine.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}

func (cr *cuisineRepository) DeleteCuisine(userId uint, cuisineId uint) error {
	result := cr.db.Where("id=? AND user_id=?", cuisineId, userId).Delete(&model.Cuisine{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}

func (cr *cuisineRepository) UpdateCuisine_Image(cuisine *model.Cuisine, image []byte, userId uint, cuisineId uint) error {
	result := cr.db.Model(cuisine).Clauses(clause.Returning{}).Where("id=? AND user_id=?", cuisineId, userId).Update("image", image)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}

func (cr *cuisineRepository) AddURL(cuisine *model.Cuisine, userId uint, cuisineId uint) error {
	result := cr.db.Model(cuisine).Clauses(clause.Returning{}).Where("id=? AND user_id=?", cuisineId, userId).Update("url", cuisine.URL)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}
