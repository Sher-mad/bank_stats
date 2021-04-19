package stats_test

import (
	"fmt"

	"github.com/Sher-mad/bank_stats/pkg/pank/stats"
	"github.com/Sher-mad/bank_stats/pkg/pank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "Караоке",
		},
		{
			ID:       2,
			Amount:   300,
			Category: "Караоке",
		},
		{
			ID:       3,
			Amount:   400,
			Category: "Караоке",
		},
	}
	fmt.Println(stats.Avg(payments))
	// Output:
	// 300
}
