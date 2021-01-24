package templates

import (
	"fmt"
	"testing"
)

func TestToGetterName(t *testing.T) {
	result :=ToGetterName("name")
	fmt.Println(result)
}