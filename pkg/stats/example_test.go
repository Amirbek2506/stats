package stats

import (
	"fmt"
	"github.com/Amirbek2506/bank/v2/pkg/types"
)

func ExampleAvg()  {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "pay",
			Status: types.StatusOk,
		},
		{
			ID:       2,
			Amount:   200,
			Category: "pay",
			Status: types.StatusFail,
		},
	}

		fmt.Println(Avg(payments))

	// Output:
	// 100
}


func ExampleTotalInCategory()  {

	fmt.Println(TotalInCategory([]types.Payment{
		{
			ID:       1,
			Amount:   300,
			Category: "pay",
			Status: types.StatusOk,
		},
		{
			ID:       2,
			Amount:   200,
			Category: "fee",
			Status: types.StatusInProgress,
		},
		{
			ID:       3,
			Amount:   444444,
			Category: "pay",
			Status: types.StatusFail,
		},
	} ,
		types.Category("pay")))

	// Output:
	// 300
}


func ExampleTotalInCategory_empty()  {
	category := types.Category("")
	payments := []types.Payment{
			{
				ID:       1,
				Amount:   300,
				Category: "pay",
				Status: types.StatusFail,
			},
			{
				ID:       2,
				Amount:   200,
				Category: "fee",
				Status: types.StatusOk,
			},
			{
			ID:       3,
				Amount:   200,
				Category: "pay",
				Status: types.StatusInProgress,
			},
	}

	fmt.Println(TotalInCategory(payments,category))

	// Output:
	// 0
}