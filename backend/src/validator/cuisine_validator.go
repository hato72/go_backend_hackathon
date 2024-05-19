package validator

import (
	//"backend/src/model"
	"github.com/hato72/go_backend_hackathon/backend/src/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICuisineValidator interface {
	CuisineValidate(cuisine model.Cuisine) error
}

type cuisineValidator struct{}

func NewCuisineValidator() ICuisineValidator { //taskValidatorのインスタンスを生成するためのコンストラクタ
	return &cuisineValidator{}
}

func (tv *cuisineValidator) CuisineValidate(cuisine model.Cuisine) error {
	return validation.ValidateStruct(&cuisine,
		validation.Field(
			&cuisine.Title,
			validation.Required.Error("title is required"), //titleに値が存在するか
			//validation.RuneLength(1, 10).Error("limited max 10 char"), //1文字から10文字までの文字数になっているかどうか
		),
	)
}
