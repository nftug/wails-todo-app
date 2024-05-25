package todo_test

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

var Now time.Time
var Tz *time.Location

func TestMain(t *testing.M) {
	// Setup
	Tz, _ = time.LoadLocation("Asia/Tokyo")
	Now = time.Now().In(Tz)
	flextime.Fix(Now)

	t.Run()

	// Teardown
}
