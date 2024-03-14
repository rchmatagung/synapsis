package repository

import (
	"context"
	"synapsis/internal/core/shopping_cart/models"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	AddProductShoppingCart(ctx context.Context, req models.ShoppingCartRequest, customerId int64, createdBy string) (*models.ShoppingCart, *exception.Error)
	GetShoppingCart(ctx context.Context, customerId int64) (*models.ShoppingCart, *exception.Error)
	DeleteShoppingCartProduct(ctx context.Context, shoppingCartProductId int64, deletedBy string) *exception.Error
	CheckoutShoppingCart(ctx context.Context, customerId int64, createdBy string) (*models.ShoppingCart, *exception.Error)
}

type Shopping_cartRepo struct {
	DBList *db.DatabaseList
}

func NewShopping_cartRepo(dbList *db.DatabaseList) Shopping_cartRepo {
	return Shopping_cartRepo{
		DBList: dbList,
	}
}

func (s Shopping_cartRepo) AddProductShoppingCart(ctx context.Context, req models.ShoppingCartRequest, customerId int64, createdBy string) (*models.ShoppingCart, *exception.Error) {

	tx := s.DBList.DatabaseApp.Begin()
	defer tx.Rollback()

	var shoppingCart models.ShoppingCart

	err := tx.Raw(qAddProductShoppingCart, customerId, req.StatusCode, req.PaymentTotal, time.Now(), createdBy).Scan(&shoppingCart).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to add product to shopping cart", "Gagal menambahkan produk ke keranjang belanja")
	}

	for _, val := range req.ShoppingCardProducts {
		var shoppingCartProduct models.ShoppingCartProduct

		err = tx.Raw(qAddProductShoppingCartProduct, shoppingCart.ShoppingCartId, val.ProductId, val.Quantity, val.SubTotal, val.Price, time.Now(), createdBy).Scan(&shoppingCartProduct).Error
		if err != nil {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to add product to shopping cart", "Gagal menambahkan produk ke keranjang belanja")
		}

		shoppingCart.ShoppingCardProducts = append(shoppingCart.ShoppingCardProducts, shoppingCartProduct)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to commit transaction", "Gagal melakukan commit transaksi")
	}

	return &shoppingCart, nil

}

func (s Shopping_cartRepo) GetShoppingCart(ctx context.Context, customerId int64) (*models.ShoppingCart, *exception.Error) {

	var shoppingCart models.ShoppingCart

	err := s.DBList.DatabaseApp.Raw(qGetShoppingCart, customerId).Scan(&shoppingCart).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get shopping cart", "Gagal mendapatkan keranjang belanja")
	}

	var shoppingCartProducts []models.ShoppingCartProduct

	err = s.DBList.DatabaseApp.Raw(qGetShoppingCartProduct, shoppingCart.ShoppingCartId).Scan(&shoppingCartProducts).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get shopping cart", "Gagal mendapatkan keranjang belanja")
	}

	shoppingCart.ShoppingCardProducts = shoppingCartProducts

	return &shoppingCart, nil

}

func (s Shopping_cartRepo) DeleteShoppingCartProduct(ctx context.Context, shoppingCartProductId int64, deletedBy string) *exception.Error {

	err := s.DBList.DatabaseApp.Exec(qDeleteShoppingCartProduct, shoppingCartProductId).Error
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to delete shopping cart product", "Gagal menghapus produk keranjang belanja")
	}

	err = s.DBList.DatabaseApp.Exec(qDecreasePaymentTotal, shoppingCartProductId, shoppingCartProductId).Error
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to delete shopping cart product", "Gagal menghapus produk keranjang belanja")
	}

	return nil

}

func (s Shopping_cartRepo) CheckoutShoppingCart(ctx context.Context, customerId int64, updatedBy string) (*models.ShoppingCart, *exception.Error) {

	tx := s.DBList.DatabaseApp.Begin()
	defer tx.Rollback()

	var shoppingCart models.ShoppingCart

	err := tx.Exec(qPaymentTransaction, time.Now(), updatedBy, customerId).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to checkout shopping cart", "Gagal checkout keranjang belanja")
	}

	err = tx.Raw(qCheckoutShoppingCart, time.Now(), updatedBy, customerId).Scan(&shoppingCart).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to checkout shopping cart", "Gagal checkout keranjang belanja")
	}

	var ShoppingCardProducts []models.ShoppingCartProduct

	err = tx.Raw(qSelectShoppingCartProduct, shoppingCart.ShoppingCartId).Scan(&ShoppingCardProducts).Error
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to checkout shopping cart", "Gagal checkout keranjang belanja")
	}

	for _, val := range ShoppingCardProducts {
		err = tx.Debug().Exec(qDecreaseStockProduct, val.Quantity, val.ProductId).Error
		if err != nil {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to checkout shopping cart", "Gagal checkout keranjang belanja")
		}

	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to commit transaction", "Gagal melakukan commit transaksi")
	}

	return &shoppingCart, nil

}
