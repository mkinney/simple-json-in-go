package main

import (
	"time"
	"log"
	"os"
	"io/ioutil"
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

func TestDecodeMissingFile(t *testing.T) {
	// do not output expected errors
	log.SetOutput(ioutil.Discard)

	// missing.json should not exist
	_, err := decode("missing.json")
	if !ErrorContains(err, "no such file or directory") {
		t.Errorf("unexpected error: %v", err)
	}

	// if you want to "reset" output
	log.SetOutput(os.Stdout)
}

func TestDecodeInvalidJson(t *testing.T) {
	// do not output expected errors
	log.SetOutput(ioutil.Discard)
	// invalid.json should exist
	_, err := decode("invalid.json")
	if !ErrorContains(err, "invalid character") {
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

func TestUnmarshalMissingFile(t *testing.T) {
	// do not output expected errors
	log.SetOutput(ioutil.Discard)
	// missing.json should not exist
	_, err := unmarshal("missing.json")
	if !ErrorContains(err, "no such file or directory") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUnmarshalInvalidJson(t *testing.T) {
	// do not output expected errors
	log.SetOutput(ioutil.Discard)
	_, err := unmarshal("invalid.json")
	if !ErrorContains(err, "invalid character") {
		t.Errorf("unexpected error: %v", err)
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
