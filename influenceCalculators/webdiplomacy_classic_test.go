package influenceCalculators

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zond/godip/variants"
	"github.com/zond/godip/variants/classical"
	"net/http"
	"testing"
)

func TestClassic(t *testing.T){
	variantName := "Classical"
	variant, found := variants.Variants[variantName]
	if !found {
		t.Error("Variant", variantName,  "not found")
		return
	}
	state, err := variant.Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
	var phase := classical.NewPhase()
	var influence =
}