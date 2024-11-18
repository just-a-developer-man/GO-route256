package controller_http

import (
	"net/http"
	"testing"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestController_OrderInfoHandler(t *testing.T) {
	type fields struct {
		Usecases Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				Usecases: tt.fields.Usecases,
			}
			c.OrderInfoHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_validateOrderInfoRequest(t *testing.T) {
	type args struct {
		req   OrderInfoRequest
		funcs []validateOrderInfoRequestFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateOrderInfoRequest(tt.args.req, tt.args.funcs...); (err != nil) != tt.wantErr {
				t.Errorf("validateOrderInfoRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateOrderID(t *testing.T) {
	type args struct {
		req OrderInfoRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateOrderID(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateOrderID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_orderInfoToResponse(t *testing.T) {
	type args struct {
		info dto.OrderInfo
	}
	tests := []struct {
		name string
		args args
		want OrderInfoResponse
	}{
		{
			name: "Test 1",
			args: args{

				info: dto.OrderInfo{
					Order: models.Order{
						ID:     models.OrderID(1),
						UserID: models.UserID(1),
						Items: []models.ItemOrderInfo{
							{
								SKU:      1,
								Quantity: 1,
							},
							{
								SKU:      2,
								Quantity: 1,
							},
						},
					},
					Status: models.StatusAwaitingPayment,
				},
			},
			want: OrderInfoResponse{

				UserID: 1,
				Status: statusAwaitingPayment,
				Items: []ItemInfo{
					{
						SKU:      1,
						Quantity: 1,
					},
					{
						SKU:      2,
						Quantity: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, orderInfoToResponse(tt.args.info))
		})
	}
}
