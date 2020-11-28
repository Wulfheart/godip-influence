package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"testing"
)

func assertAllInArray(t *testing.T, arr []godip.Province, expected bool, g godip.Graph,  f func(graph godip.Graph, n godip.Province) bool){
	for _, province := range arr {
		assert.Equal(t, expected, f(g, province), fmt.Sprintf("Expected %s province to be %v", province, expected))
	}
}
func scaffoldVariant(t *testing.T, variantName string) (s *state.State){
	variant, found := variants.Variants[variantName]
	if !found {
		t.Fatal("Variant", variantName, "not found")
	}
	s, err := variant.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
	return
}

