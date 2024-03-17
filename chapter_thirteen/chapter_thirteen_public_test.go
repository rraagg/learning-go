package chapter_thirteen_test

import (
	"fmt"
	"testing"

	chapter_thirteen "github.com/rraagg/learning-go/chapter_thirteen"
)

func TestAddNumbers(t *testing.T) {
	fmt.Println("TestAddNumbers")
	result := chapter_thirteen.AddNumbers(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
