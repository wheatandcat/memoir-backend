package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

func (r UpdateUser) Validate() error {
	err := validation.ValidateStruct(&r,
		validation.Field(
			&r.DisplayName,
			validation.Required.Error("名前の入力は必須です"),
			validation.RuneLength(1, 30).Error("名前は最大30文字までです"),
		),
	)

	return ce.NewValidationError(err.Error())
}

func (r NewItem) Validate() error {
	err := validation.ValidateStruct(&r,
		validation.Field(
			&r.Title,
			validation.Required.Error("タイトルの入力は必須です"),
			validation.RuneLength(1, 100).Error("タイトルは最大100文字までです"),
		),
	)

	return ce.NewValidationError(err.Error())
}

func (r UpdateItem) Validate() error {
	err := validation.ValidateStruct(&r,
		validation.Field(
			&r.Title,
			validation.Required.Error("タイトルの入力は必須です"),
			validation.RuneLength(1, 100).Error("タイトルは最大100文字までです"),
		),
	)

	return ce.NewValidationError(err.Error())
}
