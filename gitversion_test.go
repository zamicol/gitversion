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

func ExampleDir() {
	v := Dir("test_dir")
	fmt.Println(v)
	// Output:
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
}

// Cleans up testing.
func Clean() {
	// delete file
	var err = os.Remove("VERSION")
	if err != nil {
		fmt.Println(err)
	}

}
