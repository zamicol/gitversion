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

func ExampleVersion() {
	v, err := version()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(v)
	_ = v

	// Since "version" is not deterministic, this is an example:
	fmt.Println("v0.1.0 C144D080CCD14F38D562924AF69E6D6DA1642E0A uncommitted")
	//Output:
	// v0.1.0 C144D080CCD14F38D562924AF69E6D6DA1642E0A uncommitted
}
