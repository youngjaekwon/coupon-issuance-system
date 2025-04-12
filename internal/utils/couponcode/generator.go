package couponcode

import (
	"couponIssuanceSystem/internal/config"
	"github.com/bwmarrin/snowflake"
	"log"
	"math"
)

var (
	couponCharset []rune
	base          int64
	maxSize       int64
)

func Init() {
	couponCharset = []rune(config.AppConfig.CouponCodeRuneSet)
	base = int64(len(couponCharset))
	// 나올수 있는 전체 문자셋의 가짓수
	maxSize = int64(math.Pow(float64(base), 10.0)) // 24^10
}

type Generator interface {
	Generate() string
}

type generator struct {
	node *snowflake.Node
}

func NewGenerator() Generator {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("failed to create snowflake node: %v", err)
	}
	return &generator{node: node}
}

func (g *generator) Generate() string {
	id := g.node.Generate()
	encoded := EncodeSnowflakeToHangulNumeric(id)
	return encoded
}
