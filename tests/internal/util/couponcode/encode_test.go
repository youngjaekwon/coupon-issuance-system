package couponcode

import (
	"couponIssuanceSystem/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeSnowflakeToHangulNumeric(t *testing.T) {
	noe, err := snowflake.NewNode(1)
	assert.NoError(t, err)

	id := noe.Generate()
	encoded := couponcode.EncodeSnowflakeToHangulNumeric(id)

	assert.NotEmpty(t, encoded)
	assert.LessOrEqual(t, len(encoded), 10)

	for _, r := range encoded {
		assert.Contains(t, config.AppConfig.CouponCodeRuneSet, string(r))
	}
}
