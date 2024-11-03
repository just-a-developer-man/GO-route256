package controller_http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const goodOrderCreateRequest = `{"user":1,"items":[{"sku":1,"count":10},{"sku":2,"count":11},{"sku":3,"count":12},{"sku":4,"count":13}]}`

func TestController_CreateOrderHandler(t *testing.T) {
	type fields struct {
		Usecases Usecases
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCode     int
		wantResponse CreateOrderResponse
	}{
		{
			name: "Create order good",
			fields: fields{
				Usecases: Usecases{
					OrderManagementSystem: noErrorOMSMock(t),
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: createRequest(t, http.MethodPost, goodOrderCreateRequest),
			},
			wantCode: http.StatusOK,
			wantResponse: CreateOrderResponse{
				OrderID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				Usecases: tt.fields.Usecases,
			}
			c.CreateOrderHandler(tt.args.w, tt.args.r)
			assert.Equal(t, tt.wantCode, tt.args.w.Result().StatusCode)

			var response CreateOrderResponse
			if err := json.NewDecoder(tt.args.w.Body).Decode(&response); err != nil {
				t.Fatalf("cannot decode body from response: %v", err)
			}

			assert.Equal(t, tt.wantResponse, response)
		})
	}
}

func createRequest(t *testing.T, method string, jsonRequest string) *http.Request {
	bodyReader := bytes.NewBuffer([]byte(jsonRequest))
	req, err := http.NewRequest(method, "order/create", bodyReader)
	if err != nil {
		t.Fatalf("http.NewRequest: %v", err)
	}
	return req
}

func noErrorOMSMock(t *testing.T) *mocks.OrderManagementSystem {
	oms := mocks.NewOrderManagementSystem(t)
	oms.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(models.Order{ID: 1}, nil).Once()
	return oms
}
