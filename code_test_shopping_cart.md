# Code Test Description #

You are to build a checkout system with these items:

| SKU     | Name          | Price    | Inventory Qty     |
| --------|---------------|----------|-------------------|
| 120P90  | Google Home   | $49.99   | 10                |
| 43N23P  | MacBook Pro   | $5399.99 | 5                 |
| A304SD  | Alexa Speaker | $109.50  | 10                |
| 234234  | DongleX       | $30.00   | 2                 |

The system should have the following specials:

- Each sale of a MacBook Pro comes with a free DongleX
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers in the cart

# Implementation Requirements / Limitations #

- No frameworks unless used for testing
- Can be written in Go, Python, Scala, Java, .NET, Javascript, or Ruby
- Zip/tarball the code for submission
- Instructions on how to run the application
- Don't need a UI to make it look pretty

# Example Run With Pseudo Code #

```Go
checkout := NewCheckout(rules)
err := checkout.Scan(item1)
err = checkout.Scan(item2)
fmt.Println(checkout.Total())
...
```

# Example Scenarios #
```
Scanned Items: MacBook Pro, DongleX
Total: $5399.99

Scanned Items: Google Home, Google Home, Google Home
Total: $99.98

Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker
Total: $295.65
```

It should take about 2 hours to complete.