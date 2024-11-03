package logging

import (
	"context"
	"log/slog"
)

type customLoggerMiddleware struct {
	next slog.Handler
}

func newCustomLogger(next slog.Handler) *customLoggerMiddleware {
	return &customLoggerMiddleware{
		next: next,
	}
}

func (c *customLoggerMiddleware) Enabled(ctx context.Context, level slog.Level) bool {
	return c.next.Enabled(ctx, level)
}

func (c *customLoggerMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if loggerCtx, ok := ctx.Value(logCtxKey).(loggerRequestContext); ok {
		rec = handleRequestContext(loggerCtx, rec)
	}
	return c.next.Handle(ctx, rec)
}

func (c *customLoggerMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &customLoggerMiddleware{next: c.next.WithAttrs(attrs)}
}

func (c *customLoggerMiddleware) WithGroup(name string) slog.Handler {
	return &customLoggerMiddleware{next: c.next.WithGroup(name)}
}

type orderCreateContext struct {
	userID int64
}

type loggerRequestContext struct {
	method          string
	path            string
	orderCreateInfo *orderCreateContext
}

type keyType uint8

const logCtxKey = keyType(0)

func WithLogRequest(ctx context.Context, method string, path string) context.Context {
	if c, ok := ctx.Value(logCtxKey).(loggerRequestContext); ok {
		c.method = method
		c.path = path
		return context.WithValue(ctx, logCtxKey, c)
	}
	return context.WithValue(ctx, logCtxKey, loggerRequestContext{
		method: method,
		path:   path,
	})
}

func WithLogCreateOrder(ctx context.Context, userID int64) context.Context {
	if c, ok := ctx.Value(logCtxKey).(loggerRequestContext); ok {
		if c.orderCreateInfo != nil {
			c.orderCreateInfo.userID = userID
			return context.WithValue(ctx, logCtxKey, c)
		}
	}
	return context.WithValue(ctx, logCtxKey, loggerRequestContext{
		orderCreateInfo: &orderCreateContext{
			userID: userID,
		},
	})
}

func handleRequestContext(logCtx loggerRequestContext, rec slog.Record) slog.Record {
	rec.Add("method", logCtx.method)
	rec.Add("path", logCtx.path)

	if orderCtx := logCtx.orderCreateInfo; orderCtx != nil {
		return handleCreateOrderContext(*orderCtx, rec)
	}

	return rec
}

func handleCreateOrderContext(orderCreateInfo orderCreateContext, rec slog.Record) slog.Record {
	rec.Add("userID", orderCreateInfo.userID)
	return rec
}
