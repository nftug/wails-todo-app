package todo

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/interfaces/enums"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
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

func Arrange[T any](base T, f func(c *T)) T {
	c := base
	f(&c)
	return c
}

func TestNewTodo(t *testing.T) {
	var cmdValid = CreateCommand{
		Title:         "Test!",
		Description:   lo.ToPtr("Description"),
		InitialStatus: nil,
		DueDate:       lo.ToPtr(time.Now().In(Tz)),
	}

	t.Run("正常系", func(t *testing.T) {
		actual, err := NewTodo(cmdValid)

		dueDateWant := lo.TernaryF(cmdValid.DueDate == nil,
			func() *time.Time { return nil }, func() *time.Time { return lo.ToPtr(cmdValid.DueDate.UTC()) })
		assert.NoError(t, err)
		assert.Equal(t, cmdValid.Title, actual.Title())
		assert.Equal(t, cmdValid.Description, actual.Description())
		assert.Equal(t, lo.FromPtrOr(cmdValid.InitialStatus, "Todo"), actual.Status())
		assert.Equal(t, Now.UTC(), actual.StatusUpdatedAt())
		assert.Equal(t, dueDateWant, actual.DueDate())
		assert.Equal(t, Now.UTC(), actual.CreatedAt())
		assert.Nil(t, actual.UpdatedAt())
	})

	// 異常系
	testsErr := []struct {
		name  string
		value CreateCommand
	}{
		{"タイトルが不正", Arrange(cmdValid, func(c *CreateCommand) { c.Title = "" })},
		{"説明欄が不正",
			Arrange(cmdValid, func(c *CreateCommand) { c.Description = lo.ToPtr(lo.RandomString(300, lo.LettersCharset)) }),
		},
		{"ステータスが不正", Arrange(cmdValid, func(c *CreateCommand) { c.InitialStatus = lo.ToPtr(enums.StatusValue("Hoge")) })},
		{"期限が不正", Arrange(cmdValid, func(c *CreateCommand) { c.DueDate = lo.ToPtr(Now.AddDate(0, 0, -1)) })},
	}
	for _, tt := range testsErr {
		t.Run("異常系_"+tt.name, func(t *testing.T) {
			_, err := NewTodo(tt.value)
			assert.Error(t, err)
		})
	}
}
