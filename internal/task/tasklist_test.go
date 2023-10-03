package task

import (
	"os"
	"testing"
)

func TestNewTaskList(t *testing.T) {
	taskList := NewTaskList()

	if taskList == nil {
		t.Error("NewTaskList should return a non-nil TaskList")
	}

	if len(taskList.Tasks) != 0 {
		t.Errorf("Newly created TaskList should have an empty tasks slice, but got length %d", len(taskList.Tasks))
	}
}

func TestAddTask(t *testing.T) {
	taskList := NewTaskList()

	taskList.AddTask("Test Task 1")

	if len(taskList.Tasks) != 1 {
		t.Errorf("Expected 1 task in the list after adding one, but got %d", len(taskList.Tasks))
	}

	if taskList.Tasks[0].Title != "Test Task 1" {
		t.Errorf("Expected task description to be 'Test Task 1', but got '%s'", taskList.Tasks[0].Title)
	}
}

func TestRemoveTask(t *testing.T) {
	taskList := NewTaskList()

	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")
	taskList.AddTask("Test Task 3")

	taskList.RemoveTask(2)

	if len(taskList.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the list after removing one, but got %d", len(taskList.Tasks))
	}

	for _, task := range taskList.Tasks {
		if task.ID == 2 {
			t.Errorf("Task with ID 2 should be removed, but it still exists in the list")
		}
	}
}

func TestGetTaskByID(t *testing.T) {
	taskList := NewTaskList()

	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")

	task := taskList.GetTaskByID(2)

	if task == nil {
		t.Error("Expected to get a task, but got nil")
	}

	if task.ID != 2 {
		t.Errorf("Expected task with ID 2, but got task with ID %d", task.ID)
	}

	nonExistentTask := taskList.GetTaskByID(3)

	if nonExistentTask != nil {
		t.Errorf("Expected to get nil for a non-existent task, but got task with ID %d", nonExistentTask.ID)
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	taskList := NewTaskList()

	taskList.AddTask("Test Task 1")

	taskList.UpdateTaskDescription(1, "Updated Task 1")

	task := taskList.GetTaskByID(1)

	if task == nil {
		t.Error("Expected to get a task, but got nil")
	}

	if task.Title != "Updated Task 1" {
		t.Errorf("Expected updated task description to be 'Updated Task 1', but got '%s'", task.Title)
	}

	taskList.UpdateTaskDescription(2, "Update Non-existent Task")

	if len(taskList.Tasks) != 1 {
		t.Errorf("Expected 1 task in the list, but got %d", len(taskList.Tasks))
	}
}

func TestCompleteTask(t *testing.T) {
	taskList := NewTaskList()

	taskList.AddTask("Test Task 1")

	taskList.CompleteTask(1)

	task := taskList.GetTaskByID(1)

	if task == nil {
		t.Error("Expected to get a task, but got nil")
	}

	if !task.Completed {
		t.Errorf("Expected the task to be completed, but it's not")
	}

	taskList.CompleteTask(2)

	if len(taskList.Tasks) != 1 {
		t.Errorf("Expected 1 task in the list, but got %d", len(taskList.Tasks))
	}
}

func TestTaskListSerialization(t *testing.T) {
	taskList := NewTaskList()
	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")

	filename := "test_tasklist.json"

	err := taskList.SaveToFile(filename)
	if err != nil {
		t.Fatalf("Error saving task list to file: %v", err)
	}

	loadedTaskList := NewTaskList()
	err = loadedTaskList.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("Error loading task list from file: %v", err)
	}

	if len(loadedTaskList.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the loaded task list, but got %d", len(loadedTaskList.Tasks))
	}

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Error cleaning up test file: %v", err)
	}
}

func TestTaskListString(t *testing.T) {
	taskList := NewTaskList()
	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")

	expectedString := "[ ] Test Task 1\n[ ] Test Task 2\n"

	if str := taskList.String(); str != expectedString {
		t.Errorf("Expected String() to return '%s', but got '%s'", expectedString, str)
	}
}

func TestTaskListLength(t *testing.T) {
	taskList := NewTaskList()
	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")

	if length := taskList.Length(); length != 2 {
		t.Errorf("Expected Length() to return 2, but got %d", length)
	}
}

func TestTaskListGetTaskByIndex(t *testing.T) {
	taskList := NewTaskList()
	taskList.AddTask("Test Task 1")
	taskList.AddTask("Test Task 2")

	task1 := taskList.GetTaskByIndex(0)
	task2 := taskList.GetTaskByIndex(1)

	if task1 == nil || task1.ID != 1 || task1.Title != "Test Task 1" {
		t.Errorf("Expected to get task with ID 1 and description 'Test Task 1', but got %+v", task1)
	}

	if task2 == nil || task2.ID != 2 || task2.Title != "Test Task 2" {
		t.Errorf("Expected to get task with ID 2 and description 'Test Task 2', but got %+v", task2)
	}

	taskOutOfRange := taskList.GetTaskByIndex(2)
	if taskOutOfRange != nil {
		t.Errorf("Expected to get nil for an out-of-range index, but got %+v", taskOutOfRange)
	}
}
