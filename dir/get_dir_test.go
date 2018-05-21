package dir

import (
	"testing"
	"fmt"
)

func TestGetTargetFiles(t *testing.T) {
	fl := GetTargetFiles([]string{"/Users/hyh/go-work/src/hcp/action"})
	for _, f:=range fl{
		fmt.Println(f)
	}
	if len(fl) != 3{
		t.Fail()
	}
}

func TestGetTargetFiles2(t *testing.T) {
	fl := GetTargetFiles([]string{"/Users/hyh/go-work/src/hcp/action/*.go"})
	for _, f:=range fl{
		fmt.Println(f)
	}
	if len(fl) != 3{
		t.Fail()
	}
}
