package todo_test

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

var Now time.Time
var Tz *time.Location = time.Local

func TestMain(m *testing.M) {
	// Setup
	Now = time.Now().In(Tz)
	flextime.Fix(Now)

	m.Run()

	// Teardown
}
