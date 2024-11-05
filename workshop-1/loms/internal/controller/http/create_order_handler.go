package controller_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/logging"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
)

type validateCreateOrderRequestFunc func(*CreateOrderRequest) error

type ItemInfo struct {
	SKU      uint32 `json:"sku"`
	Quantity uint16 `json:"count"`
}

type CreateOrderRequest struct {
	UserID int64      `json:"user"`
	Items  []ItemInfo `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

// CreateOrderHandler handles `order/create` request
func (c *Controller) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 0. Decode request
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, "request body decoding failed", "err", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx = logging.WithLogCreateOrder(ctx, req.UserID)

	slog.DebugContext(ctx, "decoded json request from body", "request", req)

	// 1. Validation
	if err := validateCreateOrderRequest(&req, validateUserID, validateItems); err != nil {
		slog.ErrorContext(ctx, "request validation failed", "err", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// 2. Transform delivery layer models to Domain/Usecase models
	orderInfo := extractOrderInfo(&req)
	slog.DebugContext(ctx, "extracted order info from request", "order info", orderInfo)

	// 3. Call usecases
	orderID, err := c.OrderManagementSystem.CreateOrder(ctx, models.UserID(req.UserID), orderInfo)
	if err != nil {
		slog.ErrorContext(ctx, "order creation in OMSystem failed", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// 4. Prepare answer
	resp := CreateOrderResponse{
		OrderID: int64(orderID),
	}

	// 5. Encode answer & send response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(ctx, "order creation response encoding failed", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func validateCreateOrderRequest(req *CreateOrderRequest, funcs ...validateCreateOrderRequestFunc) error {
	for _, f := range funcs {
		if err := f(req); err != nil {
			return fmt.Errorf("request validation: %w", err)
		}
	}

	return nil
}

func validateItems(req *CreateOrderRequest) error {
	items := req.Items
	if len(items) <= 0 {
		return errors.New("no items in order")
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return fmt.Errorf("item quantity for sku=%d <= 0", item.SKU)
		}
	}

	return nil
}

func validateUserID(req *CreateOrderRequest) error {
	userID := req.UserID
	if userID <= 0 {
		return fmt.Errorf("invalid userID: %d", userID)
	}
	return nil
}

func extractOrderInfo(req *CreateOrderRequest) dto.CreateOrderInfo {
	info := dto.CreateOrderInfo{
		Items: itemInfoToItemOrderInfo(req.Items),
	}

	return info
}

func itemInfoToItemOrderInfo(items []ItemInfo) []models.ItemOrderInfo {
	modelItems := make([]models.ItemOrderInfo, 0, len(items))
	for _, item := range items {
		modelItems = append(modelItems, models.ItemOrderInfo{
			SKU:      item.SKU,
			Quantity: item.Quantity,
		})
	}
	return modelItems
}
