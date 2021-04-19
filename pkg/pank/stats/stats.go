package stats

import "github.com/Sher-mad/bank_stats/pkg/pank/types"

//Avg рассчмиывает среднюю сумма платёжаю
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		sum += payment.Amount
	}

	return sum / types.Money(len(payments))
}
