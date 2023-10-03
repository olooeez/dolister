package task

import (
	"encoding/json"
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask(1, "Test Task", false)

	if task.ID != 1 {
		t.Errorf("Expected ID to be 1, but got %d", task.ID)
	}

	if task.Title != "Test Task" {
		t.Errorf("Expected Description to be 'Test Task', but got '%s'", task.Title)
	}

	if task.Completed != false {
		t.Errorf("Expected Completed to be false, but got %v", task.Completed)
	}
}

func TestTaskString(t *testing.T) {
	// Test when task is not completed
	task := NewTask(1, "Test Task", false)
	expectedStr := "[ ] Test Task"

	if str := task.String(); str != expectedStr {
		t.Errorf("Expected String() to return '%s', but got '%s'", expectedStr, str)
	}

	// Test when task is completed
	task.Completed = true
	expectedStr = "[x] Test Task"

	if str := task.String(); str != expectedStr {
		t.Errorf("Expected String() to return '%s', but got '%s'", expectedStr, str)
	}
}

func TestTaskJSONSerialization(t *testing.T) {
	task := NewTask(1, "Test Task", false)
	expectedJSON := `{"id":1,"title":"Test Task","completed":false}`

	jsonData, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Error marshaling task to JSON: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON serialization to be '%s', but got '%s'", expectedJSON, string(jsonData))
	}
}
