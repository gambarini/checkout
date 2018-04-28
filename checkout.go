package checkout

import "errors"


const (
	ItemGoogleHomeSku   = "120P90"
	ItemMacbookProSku   = "43N23P"
	ItemAlexaSpeakerSku = "A304SD"
	ItemDongleXSku      = "234234"
)

var (
	ErrItemNotAvailable = errors.New("item not available")
	Items               = NewItems()
	Rules				= NewRules()
)

type Item struct {
	Sku   string
	Name  string
	Price int64
	Qty   int
}

type RuleApply func(checkout *Checkout) (err error)

type Rule struct {
	Apply RuleApply
}

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

	return total, err
}

func NewRules() (rules []Rule) {

	rules = make([]Rule, 3)

	rules[0] = Rule{MacFreeDongleSpecial}
	rules[1] = Rule{ThreeGoogleHomeSpecial}
	rules[2] = Rule{OverThreeAlexaSpeakerSpecial}

	return rules
}

func NewItems() (items map[string]*Item) {

	items = make(map[string]*Item, 4)

	items[ItemGoogleHomeSku] = &Item{
		Sku:   ItemGoogleHomeSku,
		Name:  "Google Home",
		Qty:   10,
		Price: 4999,
	}

	items[ItemMacbookProSku] = &Item{
		Sku:   ItemMacbookProSku,
		Name:  "MacBook Pro",
		Qty:   5,
		Price: 539999,
	}

	items[ItemAlexaSpeakerSku] = &Item{
		Sku:   ItemAlexaSpeakerSku,
		Name:  "Alexa Speaker",
		Qty:   10,
		Price: 10950,
	}

	items[ItemDongleXSku] = &Item{
		Sku:   ItemDongleXSku,
		Name:  "DongleX",
		Qty:   2,
		Price: 3000,
	}

	return items
}
