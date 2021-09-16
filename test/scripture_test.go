package test

import (
	"testing"
	"time"

	"github.com/Realive333/DailyReading/internal/scripture"
)

func Test_GetScripture(t *testing.T) {
	assert := "創世記1-2"

	actual, err := scripture.GetScripture(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), "../document/scripture.csv")
	if err != nil {
		t.Errorf("Test GetScripture Error: %v", err)
	}

	if assert != actual {
		t.Errorf("Test GetScripture: should be %s, is %s", assert, actual)
	}

	assert2 := "出埃及記32-34"

	actual2, err := scripture.GetScripture(time.Date(2021, 1, 30, 0, 0, 0, 0, time.UTC), "../document/scripture.csv")
	if err != nil {
		t.Errorf("Test GetScripture Error: %v", err)
	}

	if assert2 != actual2 {
		t.Errorf("Test GetScripture: should be %s, is %s", assert2, actual2)
	}
}
