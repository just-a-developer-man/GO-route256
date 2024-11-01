package usecase

// CreateOrderInputInfo - DTO заказа (для создания заказа)
type CreateOrderInfo struct {
	Items []ItemInfo // Товары в заказе
}

type ItemInfo struct {
	SKU   uint32
	Count uint16
}
