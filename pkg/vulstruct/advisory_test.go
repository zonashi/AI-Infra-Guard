package vulstruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvisoryEngine(t *testing.T) {
	dir := "data/advisories"
	ad, err := NewAdvisoryEngine(dir)
	assert.NoError(t, err)
	results, err := ad.GetAdvisories("mlflow", "2.13")
	assert.NoError(t, err)
	for _, result := range results {
		t.Log(result)
	}
}
