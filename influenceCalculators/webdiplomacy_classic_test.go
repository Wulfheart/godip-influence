package influenceCalculators

import (
	"github.com/stretchr/testify/assert"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"testing"
)

func TestClassic(t *testing.T) {
	variantName := "Classical"
	variant, found := variants.Variants[variantName]
	if !found {
		t.Error("Variant", variantName, "not found")
		return
	}
	state, err := variant.Start()
	if err != nil {
		t.Error(err.Error())
		return
	}

	state.SetUnit("bel", godip.Unit{
		Type:   godip.Army,
		Nation: godip.France,
	})

	var influence = WebdiplomacyClassic(defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), state)
	assert.Equal(t, godip.Neutral, influence["bel"])

	advance(t, state, 5)

	influence = WebdiplomacyClassic(defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), state)
	assert.Equal(t, godip.France, influence["bel"])

}

func advance(t *testing.T, state *state.State, times int) {
	for i := 0; i < times; i++ {
		err := state.Next()
		if err != nil {
			t.Error(err.Error())
			return
		}
	}
}
