package cloud

import "testing"

func TestGetService(t *testing.T) {
	s:=GetService("sina", TestConfig)
	s.Verity("")
}