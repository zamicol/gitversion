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
	if err := os.Remove("VERSION"); err != nil {
		fmt.Println(err)
	}
}

func TestGet(t *testing.T) {
	v, d, err := Get("./VERSION_example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Version: %s\nDate: %s\n", v, d)
}

func ExampleGetJSON() {
	j, err := GetJSON("./VERSION_example") // Uncommitted version
	if err != nil {
		fmt.Println(err)
	}
	j2, err := GetJSON("./VERSION_example2") // Committed version
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Uncommitted JSON: %s\nCommitted JSON:%s\n", j, j2)
	// Output:
	// Uncommitted JSON: {"tag":"0.0.1","hash":"EF8F94357058CE9CBA81909016B138E6D54C0381","committed":"uncommitted","build_date":"2017-02-28T19:49:11-0700"}
	// Committed JSON:{"tag":"0.0.1","hash":"EF8F94357058CE9CBA81909016B138E6D54C0381","build_date":"2017-02-28T19:49:11-0700"}
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
