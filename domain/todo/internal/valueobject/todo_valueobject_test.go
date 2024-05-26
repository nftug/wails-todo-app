package valueobject

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/domain/todo/internal/enum"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var Now time.Time
var Tz = time.Local

func TestMain(t *testing.M) {
	// Setup
	Now = time.Now().In(Tz)
	flextime.Fix(Now)

	t.Run()

	// Teardown
}

func TestTodoTitle(t *testing.T) {
	maxLength := 150

	maxLengthValue := lo.RandomString(maxLength, lo.AllCharset)
	testsValid := []struct {
		name string
		in   string
		want string
	}{
		{"通常", "Test", "Test"},
		{"スペースをトリム", "  Hello  world　　", "Hello  world"},
		{"最大文字数", maxLengthValue, maxLengthValue},
		{"最大文字数_スペースをトリム", "  " + maxLengthValue + " ", maxLengthValue},
	}
	for _, tt := range testsValid {
		t.Run("正常系_"+tt.name, func(t *testing.T) {
			actual, err := NewTitle(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, actual.Value())
		})
	}

	errEmpty := interfaces.NewInvalidArgError("title", "タイトルを設定してください")
	errTooLong := interfaces.NewInvalidArgError("title", "%d文字以内で入力してください", maxLength)
	testsErr := []struct {
		name    string
		in      string
		wantErr error
	}{
		{"空欄", "", errEmpty},
		{"空欄_スペースをトリム", "　 	", errEmpty},
		{"文字数オーバー", lo.RandomString(maxLength+1, lo.AllCharset), errTooLong},
	}
	for _, tt := range testsErr {
		t.Run("異常系_"+tt.name, func(t *testing.T) {
			_, err := NewTitle(tt.in)
			assert.EqualValues(t, err, tt.wantErr)
		})
	}
}

func TestTodoDescription(t *testing.T) {
	maxLength := 200

	maxLengthValue := lo.RandomString(maxLength, lo.AllCharset)
	testsValid := []struct {
		name string
		in   string
		want string
	}{
		{"通常", "Description", "Description"},
		{"スペースをトリム", "	Description 　", "Description"},
		{"最大文字数", maxLengthValue, maxLengthValue},
		{"最大文字数_スペースをトリム", " " + maxLengthValue + "  ", maxLengthValue},
	}
	for _, tt := range testsValid {
		t.Run("正常系_"+tt.name, func(t *testing.T) {
			actual, err := NewDescription(lo.ToPtr(tt.in))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, lo.FromPtr(actual.Value()))
		})
	}

	testsNilValid := []struct {
		name string
		in   *string
	}{
		{"空欄", nil},
		{"空欄_空白", lo.ToPtr("")},
		{"空欄_スペースをトリム", lo.ToPtr(" 　")},
	}
	for _, tt := range testsNilValid {
		t.Run("正常系_"+tt.name, func(t *testing.T) {
			actual, err := NewDescription(tt.in)
			assert.NoError(t, err)
			assert.Nil(t, actual.Value())
		})
	}

	errTooLong := interfaces.NewInvalidArgError("description", "%d文字以内で入力してください", maxLength)
	testsErr := []struct {
		name    string
		in      string
		wantErr error
	}{
		{"文字数オーバー", lo.RandomString(maxLength+1, lo.AllCharset), errTooLong},
	}
	for _, tt := range testsErr {
		t.Run("異常系_"+tt.name, func(t *testing.T) {
			_, err := NewDescription(lo.ToPtr(tt.in))
			assert.EqualValues(t, err, tt.wantErr)
		})
	}
}

func TestTodoStatus(t *testing.T) {
	testsValid := []struct {
		name string
		in   *enum.StatusValue
		want enum.StatusValue
	}{
		{"Backlog", lo.ToPtr(enum.StatusBacklog), enum.StatusBacklog},
		{"Todo", lo.ToPtr(enum.StatusTodo), enum.StatusTodo},
		{"Doing", lo.ToPtr(enum.StatusDoing), enum.StatusDoing},
		{"Done", lo.ToPtr(enum.StatusDone), enum.StatusDone},
		{"Default (Todo)", nil, enum.StatusTodo},
	}
	for _, tt := range testsValid {
		t.Run("正常系_新規作成_"+tt.name, func(t *testing.T) {
			actual, err := NewStatus(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, actual.Value())
			assert.Equal(t, Now.UTC(), actual.UpdatedAt())
		})
	}

	current := statusImpl{enum.StatusTodo, Now.UTC().AddDate(0, 0, -1)}
	testsChangeValid := []struct {
		name string
		in   enum.StatusValue
		want enum.StatusValue
	}{
		{"Backlog", enum.StatusBacklog, enum.StatusBacklog},
		{"Doing", enum.StatusDoing, enum.StatusDoing},
		{"Done", enum.StatusDone, enum.StatusDone},
	}
	for _, tt := range testsChangeValid {
		t.Run("正常系_更新_"+tt.name, func(t *testing.T) {
			actual, err := current.ChangeStatus(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, actual.Value())
			assert.Equal(t, Now.UTC(), actual.UpdatedAt())
		})
	}

	errInvalid := interfaces.NewInvalidArgError("status", "不正なステータスです")
	testsErr := []struct {
		name    string
		in      enum.StatusValue
		wantErr error
	}{
		{"存在しない値", "Hoge", errInvalid},
	}
	for _, tt := range testsErr {
		t.Run("異常系_新規作成"+tt.name, func(t *testing.T) {
			_, err := NewStatus(lo.ToPtr(tt.in))
			assert.EqualValues(t, err, tt.wantErr)
		})
	}

	errNotChanged := interfaces.NewInvalidArgError("status", "現在と異なるステータスを設定してください")
	current = statusImpl{enum.StatusTodo, Now.UTC().AddDate(0, 0, -1)}
	testsChangeErr := []struct {
		name    string
		in      enum.StatusValue
		wantErr error
	}{
		{"存在しない値", "Hoge", errInvalid},
		{"ステータス変化なし", current.value, errNotChanged},
	}
	for _, tt := range testsChangeErr {
		t.Run("異常系_更新_"+tt.name, func(t *testing.T) {
			_, err := current.ChangeStatus(tt.in)
			assert.EqualValues(t, err, tt.wantErr)
		})
	}
}

func TestNewTodoDueDate(t *testing.T) {
	testsValid := []struct {
		name string
		in   *time.Time
		want time.Time
	}{
		{"通常", lo.ToPtr(Now.AddDate(0, 0, 1)), Now.AddDate(0, 0, 1).UTC()},
	}
	for _, tt := range testsValid {
		t.Run("正常系_"+tt.name, func(t *testing.T) {
			actual, err := NewDueDate(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, lo.FromPtr(actual.Value()))
		})
	}

	t.Run("正常系_空欄", func(t *testing.T) {
		actual, err := NewDueDate(nil)
		assert.NoError(t, err)
		assert.Nil(t, actual.Value())
	})

	errPastDate := interfaces.NewInvalidArgError("dueDate", "過去の日付は指定できません。")
	testsErr := []struct {
		name    string
		in      *time.Time
		wantErr error
	}{
		{"過去の日付", lo.ToPtr(Now.AddDate(0, 0, -1)), errPastDate},
	}
	for _, tt := range testsErr {
		t.Run("異常系_"+tt.name, func(t *testing.T) {
			_, err := NewDueDate(tt.in)
			assert.EqualValues(t, err, tt.wantErr)
		})
	}
}
