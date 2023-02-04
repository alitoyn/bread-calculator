package loaf

type loaf struct {
	Flour float64
	Water float64
	Fat float64
	Yeast float64
	Salt float64
	Sugar float64
}

var BasicLoafRatio = loaf{
	Flour: 1,
	Water: 0.6,
	Fat: 0.06,
	Yeast: 0.01,
	Salt: 0.02,
	Sugar: 0.03,
}

func CalculteRatioFromFlour(flour float64, loafRatio loaf) loaf {
	return loaf{
		Flour: flour*loafRatio.Flour,
		Water: flour*loafRatio.Water,
		Fat: flour*loafRatio.Fat,
		Yeast: flour*loafRatio.Yeast,
		Salt: flour*loafRatio.Salt,
		Sugar: flour*loafRatio.Sugar,
	}
}
