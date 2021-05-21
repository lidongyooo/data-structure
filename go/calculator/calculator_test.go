package calculator

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println(Run("1+1/1-(1*1)"))
}
