package controller_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/usecase"
)

type ItemInfo struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderRequest struct {
	UserID int64      `json:"user"`
	Items  []ItemInfo `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

func (c *Controller) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	// 0. Decode request
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 1. Validation
	if err := validateCreateOrderRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Transform delivery layer models to Domain/Usecase models
	orderInfo := extractCreateOrderInfoFromCreateOrderRequest(&req)

	// 3. Call usecases
	order, err := c.OrderManagementSystem.CreateOrder(ctx, models.UserID(req.UserID), orderInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Prepare answer
	resp := CreateOrderResponse{
		OrderID: int64(order.ID),
	}

	// 5. Encode answer & send response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func validateCreateOrderRequest(req *CreateOrderRequest) error {
	if err := validateUserID(req.UserID); err != nil {
		return fmt.Errorf("validateUserID: %w", err)
	}

	if err := validateItems(req.Items); err != nil {
		return fmt.Errorf("validateItems: %w", err)
	}

	return nil
}

func validateItems(items []ItemInfo) error {
	if len(items) <= 0 {
		return errors.New("no items in order")
	}

	for _, item := range items {
		if item.Count <= 0 {
			return fmt.Errorf("item quantity for sku=%d <= 0", item.SKU)
		}
	}

	return nil
}

func validateUserID(userID int64) error {
	if userID <= 0 {
		return fmt.Errorf("invalid userID: %d", userID)
	}
	return nil
}

func extractCreateOrderInfoFromCreateOrderRequest(req *CreateOrderRequest) usecase.CreateOrderInfo {
	info := usecase.CreateOrderInfo{
		Items: itemInfoToDTOItemInfo(req.Items),
	}

	return info
}

func itemInfoToDTOItemInfo(items []ItemInfo) []usecase.ItemInfo {
	dtoItems := make([]usecase.ItemInfo, len(items))
	for _, item := range items {
		dtoItems = append(dtoItems, usecase.ItemInfo{
			SKU:   item.SKU,
			Count: item.Count,
		})
	}
	return dtoItems
}