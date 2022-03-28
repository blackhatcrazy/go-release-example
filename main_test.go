package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPrintfFixHappy(t *testing.T) {
	typeString := "fix"
	testString := "fix: test\n"

	getStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintfFix("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = getStdout

	output := string(out)
	ty := output[:strings.IndexByte(output, ':')]

	if ty != "fix" {
		t.Errorf("Expected '%s', got '%s'", typeString, ty)
	}

	if output != testString {
		t.Errorf("Expected '%s', got '%s'", testString, out)
	}
}

func TestPrintfFeatureHappy(t *testing.T) {
	typeString := "feat"
	testString := "feat: test\n"

	getStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintfFeature("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = getStdout

	output := string(out)
	ty := output[:strings.IndexByte(output, ':')]

	if ty != typeString {
		t.Errorf("Expected '%s', got '%s'", typeString, ty)
	}

	if output != testString {
		t.Errorf("Expected '%s', got '%s'", testString, out)
	}
}

func TestPrintfBreakingChangeHappy(t *testing.T) {
	testString := "feat!: test\n"

	getStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintfBreakingChange("feat", "test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = getStdout

	output := string(out)
	ty := output[:strings.IndexByte(output, ':')]

	if !strings.Contains(ty, "!") {
		t.Errorf("Expected '%s' to contain '!'", out)
	}

	if output != testString {
		t.Errorf("Expected '%s', got '%s'", testString, out)
	}
}

func TestPrintfFixSad(t *testing.T) {
	testString := "fixing: testings"

	getStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintfFix("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = getStdout

	if string(out) == testString {
		t.Errorf("Expected '%s' to not equal '%s'", testString, out)
	}
}

func TestPrintfFeatureSad(t *testing.T) {
	testString := "feature: testings"

	getStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintfFix("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = getStdout

	if string(out) == testString {
		t.Errorf("Expected '%s' to not equal '%s'", testString, out)
	}
}
