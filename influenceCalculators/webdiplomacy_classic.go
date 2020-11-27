package influenceCalculators

import (
	"fmt"
	"github.com/stretchr/stew/slice"
	godipinfluence "github.com/wulfheart/godip-influence"
	"github.com/zond/godip/state"
)

var WebdiplomacyClassic godipinfluence.InfluenceCalculator = func(old godipinfluence.Influence, adjudicated *state.State) godipinfluence.Influence {
	// on every land/coast province where a unit is located it influences the province
	for province, unit := range adjudicated.Units(){
		if slice.Contains(adjudicated.Graph().AllSCs(), province) {
			continue
		}
		old[province.Super()] = unit.Nation
	}
	// Set supply center states -> owner also influences SC
	for province, nation := range adjudicated.SupplyCenters(){
		old[province.Super()] = nation
	}

	var x = adjudicated.Graph().Edges("ber", false)
	fmt.Println(x)

	return old
}
