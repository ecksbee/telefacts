package types

type Precision int

const MININT32 = -1 << 31

const (
	Exact Precision = iota + MININT32
	Precisionless
	RuleDriven
)

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
