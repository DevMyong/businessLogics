package price

import (
	"businessLogics/model/price"
	"businessLogics/services/exchanges/upbit"
	cmap "github.com/orcaman/concurrent-map"
)

func init() {
	price.PriceBasedUSDT = cmap.New()
}

var priceFetchFunctions = map[string]func(){
	"upbit": upbit.GetPrices,
}

func Begin() {
	for _, v := range priceFetchFunctions {
		go v()
	}
}
