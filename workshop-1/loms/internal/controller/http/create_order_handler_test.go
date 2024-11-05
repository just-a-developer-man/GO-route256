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

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
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
			name: "bad body reader request",
			args: args{
				controller: NewController(Usecases{
					OrderManagementSystem: nil,
				}),
				responseRecoder: httptest.NewRecorder(),
				request:         createBadBodyReaderRequest(t),
			},
			wantStatus: http.StatusBadRequest,
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

type BadReader struct{}

func (r *BadReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("bad reader")
}

type BadWriter struct{}

func (w *BadWriter) Write(buf []byte) (n int, err error) {
	return 0, errors.New("bad writer")
}
func (w *BadWriter) Header() http.Header {
	return make(http.Header)
}
func (w *BadWriter) WriteHeader(status int) {}

func createBadBodyReaderRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest(http.MethodPut, "", &BadReader{})
	if err != nil {
		t.Fatalf("http.NewRequest: %v", err)
	}
	return req

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
	oms.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(models.OrderID(1), nil)
	return oms
}

func errorOMSMock(t *testing.T) *mocks.OrderManagementSystem {
	oms := mocks.NewOrderManagementSystem(t)
	oms.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(models.OrderID(1), errors.New("error create order"))
	return oms
}

func Test_itemInfoToItemOrderInfo(t *testing.T) {
	type args struct {
		items []ItemInfo
	}
	tests := []struct {
		name string
		args args
		want []models.ItemOrderInfo
	}{
		{
			"simple test",
			args{
				[]ItemInfo{
					{
						SKU:      10,
						Quantity: 10,
					},
					{
						SKU:      20,
						Quantity: 30,
					},
					{
						SKU:      20,
						Quantity: 30,
					},
				},
			},
			[]models.ItemOrderInfo{
				{
					SKU:      10,
					Quantity: 10,
				},
				{
					SKU:      20,
					Quantity: 30,
				},
				{
					SKU:      20,
					Quantity: 30,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := itemInfoToItemOrderInfo(tt.args.items)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_extractOrderInfo(t *testing.T) {
	type args struct {
		req *CreateOrderRequest
	}
	tests := []struct {
		name string
		args args
		want dto.CreateOrderInfo
	}{
		{
			"Request 1",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items:  []ItemInfo{{SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}},
				},
			},
			dto.CreateOrderInfo{
				Items: []models.ItemOrderInfo{{SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}, {SKU: 10, Quantity: 10}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractOrderInfo(tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateUserID(t *testing.T) {
	type args struct {
		req *CreateOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good test",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
					},
				},
			},
			false,
		},
		{
			"Bad userID",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
					},
				},
			},
			true,
		},
		{
			"No items",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items:  []ItemInfo{},
				},
			},
			true,
		},
		{
			"Empty item 1",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{},
						{},
					},
				},
			},
			true,
		},
		{
			"Empty item 2",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{
							SKU:      1,
							Quantity: 0,
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUserID(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateItems(t *testing.T) {
	type args struct {
		req *CreateOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good test",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
					},
				},
			},
			false,
		},
		{
			"No items",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items:  []ItemInfo{},
				},
			},
			true,
		},
		{
			"Empty item 1",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{},
						{},
					},
				},
			},
			true,
		},
		{
			"Empty item 2",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{
							SKU:      1,
							Quantity: 0,
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateItems(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCreateOrderRequest(t *testing.T) {
	type args struct {
		req   *CreateOrderRequest
		funcs []validateCreateOrderRequestFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good test",
			args{
				&CreateOrderRequest{
					UserID: 1,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
					},
				},
				[]validateCreateOrderRequestFunc{validateItems, validateUserID},
			},
			false,
		},
		{
			"Bad userID",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
					},
				},
				[]validateCreateOrderRequestFunc{validateItems, validateUserID},
			},
			true,
		},
		{
			"No items",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items:  []ItemInfo{},
				},
				[]validateCreateOrderRequestFunc{validateItems, validateUserID},
			},
			true,
		},
		{
			"Empty item 1",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{},
						{},
					},
				},
				[]validateCreateOrderRequestFunc{validateItems, validateUserID},
			},
			true,
		},
		{
			"Empty item 2",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{
							SKU:      1,
							Quantity: 0,
						},
					},
				},
				[]validateCreateOrderRequestFunc{validateItems, validateUserID},
			},
			true,
		},
		{
			"No ID validate",
			args{
				&CreateOrderRequest{
					UserID: 0,
					Items: []ItemInfo{
						{
							SKU:      1,
							Quantity: 1,
						},
						{
							SKU:      1,
							Quantity: 3,
						},
					},
				},
				[]validateCreateOrderRequestFunc{validateItems},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreateOrderRequest(tt.args.req, tt.args.funcs...); (err != nil) != tt.wantErr {
				t.Errorf("validateCreateOrderRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
