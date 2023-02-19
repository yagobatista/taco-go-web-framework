package handlers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type ModelUrlParams struct {
	ID string `params:"id"`
}

func CreateModel[Model any](ctx context.Context, urlParams struct{}, payload Model) (Model, error) {
	err := server.GetConnectionFromCtx(ctx).
		Create(&payload).
		Error

	return payload, err
}

func UpdateModel[Model any](ctx context.Context, urlParams ModelUrlParams, payload Model) (Model, error) {
	var model Model

	err := server.GetConnectionFromCtx(ctx).
		Model(model).
		Where("id = ? ", urlParams.ID).
		UpdateColumns(&payload).
		Error

	return payload, err
}

func GetModel[Model any](ctx context.Context, urlParams ModelUrlParams, payload struct{}) (Model, error) {
	var model Model
	var instance Model

	err := server.GetConnectionFromCtx(ctx).
		Model(model).
		Where("id = ?", urlParams.ID).
		Find(&instance).
		Error

	return instance, err
}

func ListModel[Filters any, Model any](ctx context.Context, urlParams ModelUrlParams, filters Filters) (results []Model, err error) {
	var model Model

	err = server.GetConnectionFromCtx(ctx).
		Model(model).
		Where(filters).
		Find(&results).
		Error

	return results, err
}
