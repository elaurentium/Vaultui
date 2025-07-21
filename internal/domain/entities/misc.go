package entities

type Misc struct {
	Price           float64 `json:"price"`
	Currency        string  `json:"currency"`
	OrderId         string  `json:"order_id"`
	RedirectUrl     string  `json:"redirect_url"`
	NotificationUrl string  `json:"notification_url"`
	BuyerEmail      string  `json:"buyer_email"`
}
