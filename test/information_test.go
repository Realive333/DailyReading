package test

import (
	"testing"

	"github.com/Realive333/DailyReading/internal/information"
)

func Test_CreateInfoScruct(t *testing.T) {
	assert, err := information.CreateInfoStruct()
	if err != nil {
		t.Errorf("Test CreateInfoScruct Error: %v", err)
	}

	t.Log(assert)
}

func Test_GetMessageStr(t *testing.T) {
	testcase, err := information.CreateInfoStruct()
	if err != nil {
		t.Errorf("Test GetMessageStr Error: %v", err)
	}

	assert := testcase.GetMessageStr()

	t.Log(assert)
}
