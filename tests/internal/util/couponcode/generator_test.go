package couponcode

import (
	"couponIssuanceSystem/internal/config"
	"couponIssuanceSystem/internal/utils/couponcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCode_Success(t *testing.T) {
	config.Init()
	couponcode.Init()
	gen := couponcode.NewGenerator()

	code := gen.Generate()
	assert.NotEmpty(t, code)
	assert.LessOrEqual(t, len([]rune(code)), 10)
}
