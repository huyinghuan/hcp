package utils

import (
	"testing"
	"fmt"
)

func TestGetRandomString(t *testing.T) {
	i := 20
	for i > 0{
		fmt.Println(GetRandomString(10))
		i--
	}
}
