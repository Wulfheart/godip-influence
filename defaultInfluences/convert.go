package defaultInfluences

import (
	godipInfluence "github.com/wulfheart/godip-influence"
	"github.com/zond/godip"
)

func ConvertToInfluence(def Default) (influence godipInfluence.Influence){
	var provinces = def.Variant.Graph().Provinces()
	influence = map[godip.Province]godip.Nation{}
	for _, prov := range provinces {
		influence[prov] = ""
	}
	for nation, provs := range def.Influences {
		for _, p := range provs {
			influence[p] = nation
		}
	}
	for key, val := range influence {
		if len(val) <= 0 {
			influence[key] = godip.Neutral
		}
	}
	return
}
