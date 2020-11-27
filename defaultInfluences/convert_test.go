package defaultInfluences

import (
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"testing"
)

func TestConvertToInfluence(t *testing.T) {
	var influence = ConvertToInfluence(Classical)
	assert.Equal(t, godip.Germany, influence["ber"])
	assert.Equal(t, godip.Neutral, influence["bel"])
}
