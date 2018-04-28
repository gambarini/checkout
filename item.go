package checkout

const (
	ItemGoogleHomeSku   = "120P90"
	ItemMacBookProSku   = "43N23P"
	ItemAlexaSpeakerSku = "A304SD"
	ItemDongleXSku      = "234234"
)

var (
	Items = NewItems()
)

type Item struct {
	Sku   string
	Name  string
	Price int64
	Qty   int
}

func NewItems() (items map[string]*Item) {

	items = make(map[string]*Item, 4)

	items[ItemGoogleHomeSku] = &Item{
		Sku:   ItemGoogleHomeSku,
		Name:  "Google Home",
		Qty:   10,
		Price: 4999,
	}

	items[ItemMacBookProSku] = &Item{
		Sku:   ItemMacBookProSku,
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
