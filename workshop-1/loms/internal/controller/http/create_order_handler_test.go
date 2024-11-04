package controller_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/logging"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	goodOrderCreateRequest      = `{"user":1,"items":[{"sku":1,"count":10},{"sku":2,"count":11},{"sku":3,"count":12},{"sku":4,"count":13}]}`
	badUserIDOrderCreateRequest = `{"user":0,"items":[{"sku":1,"count":10},{"sku":2,"count":11},{"sku":3,"count":12},{"sku":4,"count":13}]}`
	badItemsOrderCreateRequest  = `{"user":1,"items":[]}`
	noItemsOrderCreateRequest   = `{"user":1}`
)

func TestController_CreateOrderHandler(t *testing.T) {
	logging.InitLogger(logging.TextHandler, slog.LevelDebug, os.Stdout)
	type args struct {
		controller      *Controller
		responseRecoder *httptest.ResponseRecorder
		request         *http.Request
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantBody   CreateOrderResponse
	}{
		{
			name: "good request",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: noErrorOMSMock(t),
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodPost, "/v1/order/create", goodOrderCreateRequest),
			},
			wantStatus: http.StatusOK,
			wantBody: CreateOrderResponse{
				OrderID: 1,
			},
		},
		{
			name: "bad userID",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: nil,
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodPost, "/v1/order/create", badUserIDOrderCreateRequest),
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "bad items",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: nil,
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodPost, "/v1/order/create", badItemsOrderCreateRequest),
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "no items",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: nil,
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodPost, "/v1/order/create", noItemsOrderCreateRequest),
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "oms error",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: errorOMSMock(t),
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodPost, "/v1/order/create", goodOrderCreateRequest),
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "bad method",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: nil,
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createRequest(t, http.MethodGet, "/v1/order/create", goodOrderCreateRequest),
			},
			wantStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.controller.CreateOrderHandler(tt.args.responseRecoder, tt.args.request)
			assert.Equal(t, tt.wantStatus, tt.args.responseRecoder.Result().StatusCode)
			if tt.wantStatus != http.StatusOK {
				return
			}
			var response CreateOrderResponse
			err := json.NewDecoder(tt.args.responseRecoder.Body).Decode(&response)
			if err != nil {
				t.Errorf("error decoding response: %v", err)
			}
			assert.Equal(t, tt.wantBody, response)
		})
	}
}

func createRequest(t *testing.T, method string, url string, jsonRequest string) *http.Request {
	bodyReader := bytes.NewBuffer([]byte(jsonRequest))
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatalf("http.NewRequest: %v", err)
	}
	return req
}

func noErrorOMSMock(t *testing.T) *mocks.OrderManagementSystem {
	oms := mocks.NewOrderManagementSystem(t)
	oms.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(models.Order{ID: 1}, nil)
	return oms
}

func errorOMSMock(t *testing.T) *mocks.OrderManagementSystem {
	oms := mocks.NewOrderManagementSystem(t)
	oms.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(models.Order{ID: 0}, errors.New("error create order"))
	return oms
}
