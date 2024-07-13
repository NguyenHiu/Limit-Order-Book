package analysis

type Analysis struct {
	AliceMatchedAmount int
	BobMatchedAmount   int
	// MatchedAmount int
}

func NewAnalysis() *Analysis {
	return &Analysis{
		// MatchedAmount: 0,
		AliceMatchedAmount: 0,
		BobMatchedAmount:   0,
	}
}
