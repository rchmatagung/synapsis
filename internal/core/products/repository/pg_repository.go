package repository

import (
	"context"
	"synapsis/internal/core/products/models"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	GetAllProductByCategoryID(ctx context.Context, categoryId int) (*[]models.Product, *exception.Error)
}

type ProductsRepo struct {
	DBList *db.DatabaseList
}

func NewProductsRepo(dbList *db.DatabaseList) ProductsRepo {
	return ProductsRepo{
		DBList: dbList,
	}
}

func (r ProductsRepo) GetAllProductByCategoryID(ctx context.Context, categoryId int) (*[]models.Product, *exception.Error) {

	var response []models.Product

	err := r.DBList.DatabaseApp.Raw(qGetAllProductByCategoryID, categoryId, categoryId).Scan(&response).Error
	if err != nil {
		return nil, exception.NewError(fiber.ErrBadRequest.Code, "Failed to get products", "Gagal mengambil data produk")
	}

	return &response, nil

}
