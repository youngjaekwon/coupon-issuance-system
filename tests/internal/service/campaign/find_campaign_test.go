package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/campaign"
	svc "couponIssuanceSystem/internal/service/campaign"
	testdb "couponIssuanceSystem/tests/infra/db"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindCampaign_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	campaignID := uuid.New()

	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(1 * time.Hour),
		TotalCount: 100,
	}
	err := repository.Create(context.Background(), campaign)
	assert.NoError(t, err)

	foundCampaign, err := service.FindCampaign(context.Background(), campaignID)
	assert.NoError(t, err)
	assert.Equal(t, campaign.ID.String(), foundCampaign.ID)
	assert.Equal(t, campaign.Name, foundCampaign.Name)
}

func TestFindCampaign_NotFound(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	campaignID := uuid.New()

	_, err := service.FindCampaign(context.Background(), campaignID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCampaignNotFound))
}
