package serializer

import (
	"fmt"
	"testing"

	"github.com/betNevS/code-examples/protobuf/pcbook/sample"
)

func TestProbufToJson(t *testing.T) {
	laptop := sample.NewLaptop()
	fmt.Println(ProbufToJson(laptop))
}
