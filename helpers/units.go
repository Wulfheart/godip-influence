package helpers

import (
	"github.com/zond/godip"
)

func CanHostArmy(graph godip.Graph, province godip.Province) bool{
	return IsCoast(graph, province) || IsLand(graph, province)
}

func CanHostFleet(graph godip.Graph, province godip.Province) bool{
	return IsCoast(graph, province) || IsSea(graph, province)
}

