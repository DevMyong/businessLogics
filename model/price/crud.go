package price

import cmap "github.com/orcaman/concurrent-map"

var PriceBasedUSDT cmap.ConcurrentMap

// GetOne get price from in-memory data structure (ConcurrentMap)
func GetOne(exchange, symbol string) Price {
	return Price{}
}
func GetAll(exchange string) Prices {
	prices, exists := PriceBasedUSDT.Get(exchange)
	if !exists {
		return make(Prices, 0)
	}
	return prices.(Prices)
}

// SetOne set price to in-memory data structure (ConcurrentMap)
func SetOne(exchange, symbol, price string) error {
	return nil
}
func SetAll(exchange string, p Prices) error {
	PriceBasedUSDT.Set(exchange, p)
	return nil
}

// LoadOne load price from Database (Redis)
func LoadOne(exchange, symbol string) Price {
	return Price{}
}
func LoadAll(exchange string) Prices {
	return Prices{}
}

// SaveOne save price to Database (Redis)
func SaveOne(exchange string) error {
	return nil
}
func SaveAll(exchange string) error {
	return nil
}
