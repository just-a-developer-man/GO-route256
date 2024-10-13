package models

import (
	"github.com/google/uuid"
	"time"
)

// UserID is the type of user id
type UserID int64

// SKUID is the type of product id
type SKUID int64

// WarehouseID is the type of warehouse id
type WarehouseID int64

// DeliveryVarID is the type of delivery kind id
type DeliveryVarID int64

// SKU is the account unit of warehouse
type SKU struct {
	ID    SKUID
	Name  string
	Price uint64
}

// Order is information about order
type Order struct {
	UUID   uuid.UUID
	UserID UserID
	Items  []ItemOrderInfo
	DeliveryOrderInfo
}

// DeliveryOrderInfo is information about order delivery
type DeliveryOrderInfo struct {
	DeliveryVariantID DeliveryVarID
	DeliveryDate      time.Time
}

// ItemOrderInfo is information about order composition
type ItemOrderInfo struct {
	SKU         SKU
	Quantity    uint16
	WarehouseID WarehouseID
}
