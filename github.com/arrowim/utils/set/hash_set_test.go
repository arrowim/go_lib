package set

import (
	"fmt"
	//"runtime/debug"
	//"strings"
	"testing"
)

//testSetString(t, func() Set { return NewHashSet() }, "HashSet")
func TestToString(t *testing.T) {
	newSet := NewHashSet()
	newSet.Add(1)
	newSet.Add(2)
	t.Logf("Starting Test %s String...", "HashSet")
	fmt.Println(newSet.ToString(","))
}
