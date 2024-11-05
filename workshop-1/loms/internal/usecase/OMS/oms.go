package oms

import (
	"context"

	"github.com/google/uuid"
	controller_http "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/controller/http"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/usecase"
)

// Объявляем интерфейсы зависимостей в месте использования!
// (Задаем контракт поведения для адаптеров)
type (
	// WarehouseManagementSystem - то что отвечает за стоки товаров
	WarehouseManagementSystem interface {
		// ReserveStocks - резервация стоков на складах
		ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	// OMSRepository - репозиторий сервиса OMS
	OMSRepository interface {
		// CreateOrder - создание записи заказа в БД
		CreateOrder(ctx context.Context, order models.Order) error
	}
)

// Deps - зависимости нашего usecase
type Deps struct {
	WarehouseManagementSystem
	OMSRepository
}

type omsUsecase struct {
	Deps
}

// check that we implement usecase contarct correctly
var _ controller_http.OrderManagementSystem = (*omsUsecase)(nil)

// NewOMSUsecase - возвращаем реализацию usecase.OrderManagementSystem
func NewOMSUsecase(d Deps) *omsUsecase {
	return &omsUsecase{
		Deps: d,
	}
}

// CreateOrder - создание заказа
func (oms *omsUsecase) CreateOrder(ctx context.Context, userID models.UserID, info dto.CreateOrderInfo) (models.OrderID, error) {
	// Резервируем стоки на складах
	if err := oms.WarehouseManagementSystem.ReserveStocks(ctx, userID, info.Items); err != nil {
		return 0, usecase.ErrReserveStocks
	}

	// Формируем запись о заказе
	var (
		orderUUID = uuid.New()
		order     = models.Order{
			ID:     models.OrderID(orderUUID.ID()),
			UserID: userID,
			Items:  info.Items,
		}
	)

	// Создаем заказ в БД
	if err := oms.OMSRepository.CreateOrder(ctx, order); err != nil {
		return 0, usecase.ErrCreateOrder
	}

	return order.ID, nil
}

func (omsusecase *omsUsecase) OrderByID(ctx context.Context, orderID models.OrderID) (dto.OrderInfo, error) {
	panic("not implemented") // TODO: Implement
}
