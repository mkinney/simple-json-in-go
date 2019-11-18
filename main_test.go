package main

import (
	"time"
	"testing"
	"strings"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id. Was expecting 1 but got ", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content. Was expecting 'Hello World!' but got ", post.Content)
	}
}

// helper function
// https://stackoverflow.com/questions/42035104/how-to-unit-test-go-errors
// which looks to be from https://github.com/Teamwork/test
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestDecodeInvalidFile(t *testing.T) {
	_, err := decode("not_really_a_file.json")
	if !ErrorContains(err, "no such file or directory") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUnmarshal(t *testing.T) {
	post, err := unmarshal("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id. Was expecting 1 but got ", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content. Was expecting 'Hello World!' but got ", post.Content)
	}
}


func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
