package analysis

type Analysis struct {
	MatchedAmount int
}

func NewAnalysis() *Analysis {
	return &Analysis{
		MatchedAmount: 0,
	}
}
