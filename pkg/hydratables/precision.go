package hydratables

type Precision int

const Exact = Precision(-1 << 31)
const Precisionless = Exact + 1

const (
	Trillions Precision = iota - 12
	HundredBillions
	TenBillions
	Billions
	HundredMillions
	TenMillions
	Millions
	HundredThousands
	TenThousands
	Thousands
	Hundreds
	Tens
	Oneth
	Tenth
	Hundredth
	Thousandth
	TenThousandth
	HundredThousandth
	Millionth
	TenMillionth
	HundredMillionth
	Billionth
	TenBillionth
	HundredBillionth
	Trillionth
)
