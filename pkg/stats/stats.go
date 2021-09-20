package stats

import "github.com/Amirbek2506/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var summa types.Money = 0
	for _, payment := range payments {
		summa+=payment.Amount
	}
	return summa/types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var summa types.Money = 0
	for _, payment := range payments {
		if payment.Category == category{
			summa+=payment.Amount
		}
	}
	return summa
}