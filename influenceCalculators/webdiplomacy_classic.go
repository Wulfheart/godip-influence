package influenceCalculators

import (
	godipinfluence "github.com/wulfheart/godip-influence"
	"github.com/zond/godip"
)

var WebdiplomacyClassic godipinfluence.InfluenceCalculator = func(old godipinfluence.Influence, adjudicated godip.State) godipinfluence.Influence {
	// on every land/coast province where a unit is located it influences the province
	for province, unit := range adjudicated.Units(){
		old[province.Super()] = unit.Nation
	}
	// Set supply center states -> owner also influences SC
	for province, nation := range adjudicated.SupplyCenters(){
		old[province.Super()] = nation
	}

	return old
}
