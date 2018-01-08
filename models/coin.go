package models

type Coin struct {
	Name string
	Symbol string
	MarketCapUSD int64
	MarketCapBTC float64
	PriceUSD float64
	PriceBTC float64
	CirculatingSupply int64
	Volume24HUSD int64
	Volume24HBTC float64
	VolumeCap float64
	PriceChange1HUSD float64
	PriceChange24HUSD float64
	PriceChange7DHUSD float64
	PriceChange1HBTC float64
	PriceChange24HBTC float64
	PriceChange7DHBTC float64
}

