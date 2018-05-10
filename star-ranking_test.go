package rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRank(t *testing.T) {
	stars := []StarItem{StarItem{Count: 60, Type: 5}}
	stars = append(stars, StarItem{Count: 80, Type: 4})
	stars = append(stars, StarItem{Count: 75, Type: 3})
	stars = append(stars, StarItem{Count: 20, Type: 2})
	stars = append(stars, StarItem{Count: 25, Type: 1})

	avg := AvgStars(stars)
	assert.Equal(t, avg, 3.3493715912397968, "avg should be equal")
}
