package controller_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
)

type (
	validateOrderInfoRequestFunc func(OrderInfoRequest) error
	orderStatus                  string
)

const (
	statusNew             orderStatus = "new"
	statusAwaitingPayment orderStatus = "awaiting payment"
	statusFailed          orderStatus = "failed"
	statusPayed           orderStatus = "payed"
	statusCancelled       orderStatus = "cancelled"
)

type OrderInfoRequest struct {
	OrderID int64 `json:"orderID"`
}

type OrderInfoResponse struct {
	Status orderStatus `json:"status"`
	UserID int64       `json:"user"`
	Items  []ItemInfo  `json:"items"`
}

func (c *Controller) OrderInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req OrderInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, "request body decoding failed", "err", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := validateOrderInfoRequest(req, validateOrderID); err != nil {
		slog.ErrorContext(ctx, "request validation failed", "err", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	orderInfo, err := c.Usecases.OrderManagementSystem.OrderByID(ctx, models.OrderID(req.OrderID))
	if err != nil {
		slog.ErrorContext(ctx, "get order by id failed", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response := orderInfoToResponse(orderInfo)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(ctx, "cannot write encode or write response", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateOrderInfoRequest(req OrderInfoRequest, funcs ...validateOrderInfoRequestFunc) error {
	for _, f := range funcs {
		if err := f(req); err != nil {
			return fmt.Errorf("request validation: %w", err)
		}
	}

	return nil
}

func validateOrderID(req OrderInfoRequest) error {
	if req.OrderID < 0 {
		return errors.New("bad order ID")
	}
	return nil
}

func orderInfoToResponse(info dto.OrderInfo) OrderInfoResponse {
	items := make([]ItemInfo, 0, len(info.Items))

	for _, item := range info.Items {
		items = append(items, ItemInfo{
			SKU:      item.SKU,
			Quantity: item.Quantity,
		})
	}

	return OrderInfoResponse{
		Status: orderStatus(info.OrderStatus),
		UserID: int64(info.UserID),
		Items:  items,
	}
}
