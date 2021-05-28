package samplelib

import (
	"fmt"

	"github.com/rgao-wish/samplelib/types"
)

func Hello() string {
	return fmt.Sprintf("%s", types.Lion{})
}
