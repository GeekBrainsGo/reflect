package rflags

import (
	"reflect"
	"testing"
)

// source="./data" debug output=out

type OutFlags struct {
	Source string   `rflag:"source,s,src"`
	Debug  bool     `rflag:"debug,d"`
	Port   int      `rflag:"port,p"`
	Tags   []string `rflag:"tags,t"`
	Output string
}

func TestParseFlags(t *testing.T) {
	f := OutFlags{}
	args := []string{`source="./data"`, `debug`, `output=out`, `p=1233`, `t=tag1`, `t=tag2`}
	if err := ParseFlags(&f, args); err != nil {
		t.Error(err)
	}
	if f.Source != "./data" {
		t.Errorf("Source should be: %s, got: %s", "./data", f.Source)
	}
	if f.Debug != true {
		t.Errorf("Debug should be: %v, got: %v", true, f.Debug)
	}
	if f.Output != "out" {
		t.Errorf("Output should be: %s, got: %s", "out", f.Output)
	}
	if f.Port != 1233 {
		t.Errorf("Port should be: %d, got: %d", 1233, f.Port)
	}
	expectedTags := []string{"tag1", "tag2"}
	if !reflect.DeepEqual(expectedTags, f.Tags) {
		t.Errorf("expected Tags=%v, but got %v", expectedTags, f.Tags)
	}
}
