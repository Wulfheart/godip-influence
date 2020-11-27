package influenceCalculators

import (
	"github.com/stretchr/testify/assert"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip"
	"github.com/zond/godip/variants"
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

	// err = state.SetOrder("par", orders.Move("par", "pic"))
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }
	// err = state.Next()
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	state.SetUnit("bel", godip.Unit{
		Type:   godip.Army,
		Nation: godip.France,
	})




	var influence = WebdiplomacyClassic(defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), state)
	assert.Equal(t, godip.Neutral, influence["bel"])

	state.Next()
	state.Next()
	 influence = WebdiplomacyClassic(defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), state)
	assert.Equal(t, godip.France, influence["bel"])

}