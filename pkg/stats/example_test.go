package stats

import (
	"fmt"
	"github.com/Amirbek2506/bank/pkg/types"
)

func ExampleAvg()  {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "pay",
		},
		{
			ID:       2,
			Amount:   200,
			Category: "pay",
		},
	}

		fmt.Println(Avg(payments))

	// Output:
	// 150
}


func ExampleTotalInCategory()  {

	fmt.Println(TotalInCategory([]types.Payment{
		{
			ID:       1,
			Amount:   3,
			Category: "pay",
		},
		{
			ID:       2,
			Amount:   200,
			Category: "fee",
		},
		{
			ID:       3,
			Amount:   4,
			Category: "pay",
		},
	} ,
		types.Category("fee")))

	// Output:
	// 200
}


func ExampleTotalInCategory_empty()  {
	category := types.Category("")
	payments := []types.Payment{
			{
				ID:       1,
				Amount:   300,
				Category: "pay",
			},
			{
			ID:       2,
				Amount:   200,
				Category: "fee",
			},
			{
			ID:       3,
				Amount:   200,
				Category: "pay",
			},
	}

	fmt.Println(TotalInCategory(payments,category))

	// Output:
	// 0
}