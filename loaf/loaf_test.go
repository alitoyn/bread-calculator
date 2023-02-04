package loaf

import (
	"fmt"
    "testing"
)

func TestLoaf(t *testing.T) {
	calculatedLoaf := CalculteRatioFromFlour(200, BasicLoafRatio)
	fmt.Printf("%+v", calculatedLoaf)
}
