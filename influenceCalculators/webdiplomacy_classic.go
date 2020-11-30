package influenceCalculators

import (
	"github.com/stretchr/stew/slice"
	godipinfluence "github.com/wulfheart/godip-influence"
	"github.com/wulfheart/godip-influence/helpers"
	"github.com/zond/godip/state"
)

var WebdiplomacyClassic godipinfluence.InfluenceCalculator = func(old godipinfluence.Influence, adjudicated *state.State) godipinfluence.Influence {
	// on every land/coast province where a unit is located it influences the province
	for province, unit := range adjudicated.Units() {
		if ok := slice.Contains(adjudicated.Graph().AllSCs(), province); ok {
			continue
		}
		if !helpers.CanHostArmy(adjudicated.Graph(), province){
			continue
		}
		old[province] = unit.Nation
	}
	// Set supply center states -> owner also influences SC
	for province, _ := range adjudicated.SupplyCenters() {
		nation, sup, ok := adjudicated.SupplyCenter(province)
		if ok {

			old[sup] = nation
		}
	}
	return old
}
