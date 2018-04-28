package checkout

import "fmt"

func MacFreeDongleSpecial(checkout *Checkout) (err error) {

	macPurchase, ok := checkout.Purchases[ItemMacbookProSku]

	if !ok {
		return nil
	}

	dongleItem, _ := Items[ItemDongleXSku]

	for i := 0; i < macPurchase.Qty; i++ {

		if dongleItem.Qty == 0 {
			return fmt.Errorf("cannot apply MacFreeDongleSpecial, dongle not available")
		}

		donglePurchase, ok := checkout.Purchases[ItemDongleXSku]

		if !ok {
			checkout.Purchases[ItemDongleXSku] = &Purchase{1, dongleItem.Price}
		} else {
			donglePurchase.Qty += 1
			donglePurchase.Discount += dongleItem.Price
		}

		dongleItem.Qty -= 1
	}

	return nil
}

func ThreeGoogleHomeSpecial(checkout *Checkout) (err error) {

	googlePurchase, ok := checkout.Purchases[ItemGoogleHomeSku]

	if !ok {
		return nil
	}

	googleItem, _ := Items[ItemGoogleHomeSku]

	for i := 1; i <= googlePurchase.Qty; i++ {

		if i % 3 == 0 {

			googlePurchase.Discount += googleItem.Price
		}
	}

	return nil
}

func OverThreeAlexaSpeakerSpecial(checkout *Checkout) (err error) {

	alexaPurchase, ok := checkout.Purchases[ItemAlexaSpeakerSku]

	if !ok {
		return nil
	}

	alexaItem, _ := Items[ItemAlexaSpeakerSku]

	if alexaPurchase.Qty > 3 {
		total := (alexaItem.Price * int64(alexaPurchase.Qty)) / 100

		discount := float32(total) * float32(0.1)

		alexaPurchase.Discount += int64(discount * 100)
	}

	return nil
}
