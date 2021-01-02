package types

type CalculationWeight int

const (
	Positive CalculationWeight = iota
	Negative
)

type HypercubeContext int

const (
	ClosedSegment HypercubeContext = iota
	ClosedScenario
	Open
)

type HypercubeArcrole int

const (
	All HypercubeArcrole = iota
	NotAll
)
