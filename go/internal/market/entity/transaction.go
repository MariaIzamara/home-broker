package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) DoTransaction() {
	buyingShares := t.BuyingOrder.PendingShares
	sellingShares := t.SellingOrder.PendingShares

	sharesSoldOrBought := sellingShares
	if buyingShares < sharesSoldOrBought {
		sharesSoldOrBought = buyingShares
	}

	t.BuyingOrder.Investor.UpdateAssetPosition(t.BuyingOrder.Asset.ID, sharesSoldOrBought)
	t.BuyingOrder.UpdatePendingShares(sharesSoldOrBought)

	t.SellingOrder.Investor.UpdateAssetPosition(t.SellingOrder.Asset.ID, -sharesSoldOrBought)
	t.SellingOrder.UpdatePendingShares(sharesSoldOrBought)

	t.BuyingOrder.CloseOrder()
	t.SellingOrder.CloseOrder()
}
