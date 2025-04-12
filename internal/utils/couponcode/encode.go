package couponcode

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

func EncodeSnowflakeToHangulNumeric(id snowflake.ID) string {
	/*
		Snowflake로 생성한 int64 ID 값을 '숫자 + 한글' 조합으로 변환하는 함수

		변환 과정:
		1. Snowflake ID를 24진수로 변환 (숫자 + 한글 -> 24진수)
		2. 변환된 각 자릿수는 문자셋에서 인덱싱하여 매핑
		3. 전체 문자열을 뒤집음 -> 진변 변환 과정
		4. 10자리 미만일 경우 랜덤 문자로 패딩 -> 최소 10자리 보장
		5. 10자리 초과일 경우 뒤에서 10자리만 남김
		6. 최종적으로 변환된 문자열 반환
	*/
	num := id.Int64()

	modNum := num % maxSize
	if modNum < 0 {
		modNum = -modNum
	}

	var reversed []rune
	for modNum > 0 {
		remainder := modNum % base
		reversed = append(reversed, couponCharset[remainder])
		modNum /= base
	}

	if len(reversed) == 0 {
		reversed = append(reversed, couponCharset[0])
	}

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	n := len(reversed)
	if n < 10 {
		padCount := 10 - n
		padRunes := make([]rune, padCount)
		for i := 0; i < padCount; i++ {
			randomIndex := rand.Intn(len(couponCharset))
			padRunes[i] = couponCharset[randomIndex]
		}
		reversed = append(reversed, padRunes...)
	} else if n > 10 {
		reversed = reversed[n-10:]
	}

	return string(reversed)
}
