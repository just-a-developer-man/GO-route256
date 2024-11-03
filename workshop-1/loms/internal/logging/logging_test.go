package logging

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLogRequest(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		path   string
	}
	tests := []struct {
		name string
		args args
		want loggerRequestContext
	}{
		{
			name: "test1",
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				path:   "order/create",
			},
			want: loggerRequestContext{
				method: http.MethodGet,
				path:   "order/create",
			},
		},
		{
			name: "test2",
			args: args{
				ctx: context.WithValue(context.Background(), logCtxKey, loggerRequestContext{
					method: http.MethodPut,
					path:   "test/path",
				}),
				method: http.MethodGet,
				path:   "order/create",
			},
			want: loggerRequestContext{
				method: http.MethodGet,
				path:   "order/create",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithLogRequest(tt.args.ctx, tt.args.method, tt.args.path)
			if reqCtx, ok := got.Value(logCtxKey).(loggerRequestContext); ok {
				assert.Equal(t, tt.want, reqCtx)
			} else {
				t.Errorf("No loggerRequestContext in context")
			}
		})
	}
}

func TestWithLogUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name string
		args args
		want loggerRequestContext
	}{
		{
			name: "test1",
			args: args{
				ctx:    context.Background(),
				userID: 20,
			},
			want: loggerRequestContext{
				orderCreateInfo: &orderCreateContext{
					userID: 20,
				},
			},
		},
		{
			name: "test2",
			args: args{
				ctx: context.WithValue(context.Background(), logCtxKey, loggerRequestContext{
					orderCreateInfo: &orderCreateContext{
						userID: 10,
					},
				}),
				userID: 20,
			},
			want: loggerRequestContext{
				orderCreateInfo: &orderCreateContext{
					userID: 20,
				},
			},
		},
		{
			name: "test3",
			args: args{
				ctx: context.WithValue(context.Background(), logCtxKey, loggerRequestContext{
					method: http.MethodGet,
					path:   "create/order",
					orderCreateInfo: &orderCreateContext{
						userID: 10,
					},
				}),
				userID: 20,
			},
			want: loggerRequestContext{
				method: http.MethodGet,
				path:   "create/order",
				orderCreateInfo: &orderCreateContext{
					userID: 20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithLogUserID(tt.args.ctx, tt.args.userID)
			if reqCtx, ok := got.Value(logCtxKey).(loggerRequestContext); ok {
				assert.Equal(t, tt.want, reqCtx)
			} else {
				t.Errorf("No loggerRequestContext in context")
			}
		})
	}
}

func Test_newCustomLogger(t *testing.T) {
	defaultHanlder := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{})
	testMiddlwareHandler := newCustomLogger(defaultHanlder)

	assert.Exactly(t, defaultHanlder, testMiddlwareHandler.next)
}
