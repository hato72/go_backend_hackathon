package usecase

import (
	"backend/src/model"
	"backend/src/repository"
	"backend/src/validator"
)

type ICuisineUsecase interface {
	GetAllCuisines(userId uint) ([]model.CuisineResponse, error)
	GetCuisineById(userId uint, cuisineId uint) (model.CuisineResponse, error)
	CreateCuisine(cuisine model.Cuisine) (model.CuisineResponse, error)
	UpdateCuisine(cuisine model.Cuisine, userId uint, cuisineId uint) (model.CuisineResponse, error)
	DeleteCuisine(userId uint, cuisineId uint) error
}

type cuisineUsecase struct {
	cr repository.ICuisineRepository
	cv validator.ICuisineValidator
}

func NewCuisineUsecase(tr repository.ICuisineRepository, tv validator.ICuisineValidator) ICuisineUsecase { //コンストラクタ
	return &cuisineUsecase{tr, tv}
}

func (cu *cuisineUsecase) GetAllCuisines(userId uint) ([]model.CuisineResponse, error) {
	cuisines := []model.Cuisine{}
	if err := cu.cr.GetAllCuisines(&cuisines, userId); err != nil {
		return nil, err
	}
	resCuisines := []model.CuisineResponse{}
	for _, v := range cuisines {
		t := model.CuisineResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resCuisines = append(resCuisines, t)
	}
	return resCuisines, nil
}

func (cu *cuisineUsecase) GetCuisineById(userId uint, cuisineId uint) (model.CuisineResponse, error) {
	cuisine := model.Cuisine{}
	if err := cu.cr.GetCuisineById(&cuisine, userId, cuisineId); err != nil {
		return model.CuisineResponse{}, err
	}
	rescuisine := model.CuisineResponse{
		ID:        cuisine.ID,
		Title:     cuisine.Title,
		CreatedAt: cuisine.CreatedAt,
		UpdatedAt: cuisine.UpdatedAt,
	}
	return rescuisine, nil
}

func (cu *cuisineUsecase) CreateCuisine(cuisine model.Cuisine) (model.CuisineResponse, error) {
	if err := cu.cv.CuisineValidate(cuisine); err != nil {
		return model.CuisineResponse{}, err
	}
	if err := cu.cr.CreateCuisine(&cuisine); err != nil {
		return model.CuisineResponse{}, err
	}
	rescuisine := model.CuisineResponse{
		ID:        cuisine.ID,
		Title:     cuisine.Title,
		CreatedAt: cuisine.CreatedAt,
		UpdatedAt: cuisine.UpdatedAt,
	}
	return rescuisine, nil
}

func (cu *cuisineUsecase) UpdateCuisine(cuisine model.Cuisine, userId uint, cuisineId uint) (model.CuisineResponse, error) {
	if err := cu.cr.UpdateCuisine(&cuisine, userId, cuisineId); err != nil {
		return model.CuisineResponse{}, err
	}
	rescuisine := model.CuisineResponse{
		ID:        cuisine.ID,
		Title:     cuisine.Title,
		CreatedAt: cuisine.CreatedAt,
		UpdatedAt: cuisine.UpdatedAt,
	}
	return rescuisine, nil
}

func (cu *cuisineUsecase) DeleteCuisine(userId uint, cuisineId uint) error {
	if err := cu.cr.DeleteCuisine(userId, cuisineId); err != nil {
		return err
	}
	return nil
}
