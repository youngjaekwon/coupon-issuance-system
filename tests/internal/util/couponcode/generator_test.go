package couponcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCode_Success(t *testing.T) {
	gen := couponcode.NewGenerator()

	code := gen.Generate()
	assert.NotEmpty(t, code)
	assert.LessOrEqual(t, len(code), 10)
}
