package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	repo "couponIssuanceSystem/internal/repository/campaign"
	testdb "couponIssuanceSystem/tests/infra/db"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateCampaign_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	input := &svc.CreateCampaignInput{
		Name:       "Test Campaign",
		TotalCount: 100,
		StartAt:    time.Now().Add(1 * time.Hour),
	}

	campaign, err := service.CreateCampaign(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, input.Name, campaign.Name)
	assert.Equal(t, input.TotalCount, campaign.TotalCount)
}

func TestCreateCampaign_InvalidInput(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	input := &svc.CreateCampaignInput{
		Name: "Test Campaign",
	}

	_, err := service.CreateCampaign(context.Background(), input)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrInvalidCampaignInput))
}
