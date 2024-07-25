package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func GetKWDesc(balance decimal.Decimal) string {
	var balanceDesc string
	zero := decimal.NewFromInt(0)
	thousand := decimal.NewFromInt(1000)
	tenThousand := decimal.NewFromInt(10000)
	ninetyNineThousand := decimal.NewFromInt(990000)

	if balance.Equal(zero) {
		balanceDesc = "0"
	} else if balance.LessThan(thousand) {
		balanceDesc = balance.StringFixedBank(0)
	} else if balance.LessThan(tenThousand) {
		kPart := balance.DivRound(thousand, 0)
		yusb := balance.Mod(thousand).GreaterThan(zero)
		balanceDesc = fmt.Sprintf("%sk%s", kPart.StringFixedBank(0), IfThenElse(yusb, "+", ""))
	} else if balance.LessThan(ninetyNineThousand) {
		wPart := balance.DivRound(tenThousand, 0)
		yusk := balance.Mod(thousand).GreaterThan(zero)
		balanceDesc = fmt.Sprintf("%sw%s", wPart.StringFixedBank(0), IfThenElse(yusk, "+", ""))
	} else {
		balanceDesc = "99w+"
	}

	return balanceDesc
}

func IfThenElse(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}
