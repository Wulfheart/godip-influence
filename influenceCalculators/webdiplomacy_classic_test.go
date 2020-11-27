package influenceCalculators

import (
	"fmt"
	"github.com/zond/godip/variants"
	"testing"
)

func TestClassic(t *testing.T){
	var variantName = "Classical"
	variant, found := variants.Variants[variantName]
	if !found {
		t.Errorf("Variant %q not found", variantName)
	}
	state, err := variant.Start()
	a := state.Graph().Provinces()
	fmt.Println(variant, state, err, a)
}