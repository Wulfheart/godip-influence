package defaultInfluences

import (
	"github.com/zond/godip"
	"github.com/zond/godip/variants/classical/start"
)

import "testing"

func TestClassicalInfluences(t *testing.T) {
	var provinces = start.Graph().Provinces()
	for _, x := range Classical.Influences {
		for _, y := range x {
			if !contains(provinces, y) {
				t.Error(y, "not found in provinces")
			}
		}
	}

}

func contains(s []godip.Province, e godip.Province) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
