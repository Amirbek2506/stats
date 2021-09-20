package stats

import "github.com/Amirbek2506/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var summa types.Money = 0
	var count types.Money = 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail{
			summa+=payment.Amount
			count++
		}
	}
	if count==0{
		return 0
	}
	return summa/count
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var summa types.Money = 0
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail{
			summa+=payment.Amount
		}
	}
	return summa
}