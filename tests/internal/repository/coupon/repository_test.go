package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/coupon"
	testdb "couponIssuanceSystem/tests/infra/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestCoupon(code string, campaignID uuid.UUID, userID string) *models.Coupon {
	return &models.Coupon{
		Code:       code,
		CampaignID: campaignID,
		UserID:     userID,
	}
}

func TestCreateCoupon_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	coupon := newTestCoupon(
		"TESTCODE",
		uuid.New(),
		"testuser",
	)

	err := repository.Create(context.Background(), coupon)
	assert.NoError(t, err)

	var found models.Coupon
	err = db.WithContext(context.Background()).First(&found, "code = ?", coupon.Code).Error
	assert.NoError(t, err)
	assert.Equal(t, coupon.Code, found.Code)
}

func TestCreateCoupon_DuplicateCode(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	coupon := newTestCoupon(
		"TESTCODE",
		uuid.New(),
		"testuser",
	)

	err := repository.Create(context.Background(), coupon)
	assert.NoError(t, err)

	duplicateCoupon := *coupon
	duplicateCoupon.Code = coupon.Code

	err = repository.Create(context.Background(), &duplicateCoupon)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE")
}

func TestCreateCoupon_DuplicateUserIDInCampaign(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaignID := uuid.New()

	coupon1 := newTestCoupon(
		"TESTCODE1",
		campaignID,
		"testuser",
	)

	err := repository.Create(context.Background(), coupon1)
	assert.NoError(t, err)

	coupon2 := newTestCoupon(
		"TESTCODE2",
		campaignID,
		"testuser",
	)

	err = repository.Create(context.Background(), coupon2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE")
}
