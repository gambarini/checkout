package checkout

import "errors"

var (
	ErrItemNotAvailable = errors.New("item not available")
)

type Checkout struct {
	Rules     []Rule
	Purchases map[string]*Purchase
}

type Purchase struct {
	Qty      int
	Discount int64
}

func NewCheckout(rules []Rule) (checkout Checkout) {

	purchases := make(map[string]*Purchase, len(Items))

	return Checkout{
		Rules:     rules,
		Purchases: purchases,
	}
}

func (checkout *Checkout) Scan(item *Item) (err error) {

	if item.Qty == 0 {
		return ErrItemNotAvailable
	}

	purchase, ok := checkout.Purchases[item.Sku]

	if !ok {
		checkout.Purchases[item.Sku] = &Purchase{1, 0}
	} else {
		purchase.Qty += 1
	}

	item.Qty -= 1

	return nil

}

func (checkout Checkout) HasItem(sku string) (qty int) {

	purchase, ok := checkout.Purchases[sku]

	if !ok {
		return 0
	}

	return purchase.Qty
}

func (checkout *Checkout) Total() (total int64, err error) {

	for _, rule := range checkout.Rules {
		err = rule.Apply(checkout)

		if err != nil {
			return total, err
		}
	}

	for sku, purchase := range checkout.Purchases {
		total += (Items[sku].Price * int64(purchase.Qty)) - purchase.Discount
	}

	return total, nil
}
