package gitversion

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	// Run example.
	Write("VERSION")

	Clean()
}

// Cleans up testing.
func Clean() {
	// delete file
	var err = os.Remove("VERSION")
	if err != nil {
		fmt.Println(err)
	}

}
