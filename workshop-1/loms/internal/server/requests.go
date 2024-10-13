package server

type OrderCreateRequest struct {
	User  int64 `json:"user,omitempty"`
	Items []struct {
		SKU   uint32 `json:"sku,required"`
		Count uint16 `json:"count,omitempty"`
	} `json:"items,omitempty"`
}

type OrderCreateResponse struct {
	OrderID int64
}
