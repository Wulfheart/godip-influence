package helpers

import (
	"github.com/zond/godip"
)

func IsLand(graph godip.Graph, n godip.Province) (isLand bool) {

	flags := graph.AllFlags(n)
	hasLand := flags[godip.Land]
	hasSea := flags[godip.Sea]
	return hasLand && !hasSea
}

func IsSea(graph godip.Graph, n godip.Province) (isSea bool) {
	flags := graph.AllFlags(n)
	hasLand := flags[godip.Land]
	hasSea := flags[godip.Sea]
	return !hasLand && hasSea
}

func IsCoast(graph godip.Graph, n godip.Province) (isCoast bool) {
	flags := graph.AllFlags(n)
	hasLand := flags[godip.Land]
	hasSea := flags[godip.Sea]
	return hasLand && hasSea
}
