package datetime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	tm := time.Now()
	expectedTime := fmt.Sprintf("%d %s %d", tm.Day(), tm.Weekday(), tm.Year())
	assert.Equal(t, expectedTime, TimeToString(tm, "02 Jan 2006"))
}
