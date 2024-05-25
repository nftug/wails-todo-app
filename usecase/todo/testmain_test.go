package todo_test

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

var Now time.Time
var Tz *time.Location = time.Local

func TestMain(t *testing.M) {
	// Setup
	Now = time.Now().In(Tz)
	flextime.Fix(Now)

	t.Run()

	// Teardown
}

func arrange[T any](base T, f func(c *T)) T {
	c := base
	f(&c)
	return c
}
