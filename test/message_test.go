package test

import (
	"testing"

	"github.com/Realive333/DailyReading/internal/infomation"
)

func Test_CreateInfoScruct(t *testing.T) {
	assert, err := infomation.CreateInfoStruct()
	if err != nil {
		t.Errorf("Test CreateInfoScruct Error: %v", err)
	}

	t.Log(assert)
}
