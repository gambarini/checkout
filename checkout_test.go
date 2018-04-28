package checkout

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOneItemCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemGoogleHomeSku]
	err := checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Empty(t, err)
	assert.Equal(t, 9, Items[ItemGoogleHomeSku].Qty)
	assert.Equal(t, 1, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, int64(4999), total)

}

func TestTwoItemCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemGoogleHomeSku]
	err := checkout.Scan(item)

	item, _ = Items[ItemDongleXSku]
	err = checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Empty(t, err)
	assert.Equal(t, 9, Items[ItemGoogleHomeSku].Qty)
	assert.Equal(t, 1, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, 1, Items[ItemDongleXSku].Qty)
	assert.Equal(t, 1, checkout.HasItem(ItemDongleXSku))
	assert.Equal(t, int64(7999), total)

}

func TestNotAvailableItemCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemGoogleHomeSku]
	item.Qty = 0
	err := checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, err, ErrItemNotAvailable)
	assert.Equal(t, 0, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, int64(0), total)

}

func TestMacFreeDongleSpecialCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemMacBookProSku]
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 1, checkout.HasItem(ItemMacBookProSku))
	assert.Equal(t, 1, checkout.HasItem(ItemDongleXSku))
	assert.Equal(t, int64(539999), total)

}

func TestTwoMacFreeDongleSpecialCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemMacBookProSku]
	checkout.Scan(item)
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 2, checkout.HasItem(ItemMacBookProSku))
	assert.Equal(t, 2, checkout.HasItem(ItemDongleXSku))
	assert.Equal(t, int64(1079998), total)

}

func TestThreeGoogleHomeSpecialCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemGoogleHomeSku]
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 3, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, int64(9998), total)

}

func TestSixGoogleHomeSpecialCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemGoogleHomeSku]
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 6, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, int64(19996), total)

}

func TestOverThreeAlexaSpeakerSpecialCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemAlexaSpeakerSku]
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 4, checkout.HasItem(ItemAlexaSpeakerSku))
	assert.Equal(t, int64(39420), total)

}

func TestAllSpecialsCheckout(t *testing.T) {

	Items = NewItems()

	checkout := NewCheckout(Rules)

	item, _ := Items[ItemAlexaSpeakerSku]
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)

	item, _ = Items[ItemMacBookProSku]
	checkout.Scan(item)

	item, _ = Items[ItemGoogleHomeSku]
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)
	checkout.Scan(item)

	item, _ = Items[ItemDongleXSku]
	checkout.Scan(item)

	total, _ := checkout.Total()

	assert.Equal(t, 4, checkout.HasItem(ItemAlexaSpeakerSku))
	assert.Equal(t, 1, checkout.HasItem(ItemMacBookProSku))
	assert.Equal(t, 4, checkout.HasItem(ItemGoogleHomeSku))
	assert.Equal(t, 2, checkout.HasItem(ItemDongleXSku))
	assert.Equal(t, int64(597416), total)

}
