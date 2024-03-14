package repository

import (
	"context"
	"synapsis/internal/core/category/models"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	GetAllCategory(context.Context) (*[]models.Category, *exception.Error)
}

type CategoryRepo struct {
	DBList *db.DatabaseList
}

func NewCategoryRepo(dbList *db.DatabaseList) CategoryRepo {
	return CategoryRepo{
		DBList: dbList,
	}
}

func (c CategoryRepo) GetAllCategory(ctx context.Context) (*[]models.Category, *exception.Error) {

	var response []models.Category

	err := c.DBList.DatabaseApp.Raw(qGetAllCategory).Scan(&response).Error
	if err != nil {
		return nil, exception.NewError(fiber.ErrBadRequest.Code, "Failed to get all category", "Gagal mengambil semua kategori")
	}

	return &response, nil

}
