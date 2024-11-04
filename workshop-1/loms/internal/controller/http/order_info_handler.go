package controller_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateOrderInfoRequest(req, validateOrderID); err != nil {
		slog.ErrorContext(ctx, "request validation failed", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
