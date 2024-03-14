package repository

import (
	"context"
	"synapsis/internal/core/customers/models"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"
	"synapsis/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	Register(ctx context.Context, req models.CustomerRegister) (*models.Customer, *exception.Error)
	Login(ctx context.Context, username, password string) (*models.LoginResponse, *exception.Error)
}

type CustomersRepo struct {
	DBList *db.DatabaseList
}

func NewCustomersRepo(dbList *db.DatabaseList) CustomersRepo {
	return CustomersRepo{
		DBList: dbList,
	}
}

func (c CustomersRepo) Register(ctx context.Context, req models.CustomerRegister) (*models.Customer, *exception.Error) {

	var response models.Customer

	passwordHash, _ := utils.HashingPassword(req.Password)

	err := c.DBList.DatabaseApp.Raw(qRegisterCustomer, req.Username, passwordHash, req.Email, time.Now(), "system").Scan(&response).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusBadRequest, "Failed to register customer", "Gagal mendaftarkan customer")
	}

	return &response, nil

}

func (c CustomersRepo) Login(ctx context.Context, username, password string) (*models.LoginResponse, *exception.Error) {

	var response models.LoginResponse

	err := c.DBList.DatabaseApp.Raw(qLoginCustomer, username).Scan(&response).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusBadRequest, "customer not found", "customer tidak ditemukan")
	}

	if !utils.CheckHashedPassword(response.Password, password) {
		return nil, exception.NewError(fiber.StatusBadRequest, "password is incorrect", "password salah")
	}

	return &response, nil

}
