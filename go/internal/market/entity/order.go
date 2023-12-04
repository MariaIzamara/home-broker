package entity

// TODO: OrderType mudar de string para constante com as duas possibilidades de tipo (buy ou sell)
type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}

func (o *Order) UpdatePendingShares(shares int) {
	o.PendingShares -= shares
}

func (o *Order) CloseOrder() {
	if o.PendingShares == 0 {
		o.Status = "CLOSED"
	}
}
