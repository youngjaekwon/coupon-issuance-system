package campaign

import (
	"context"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/campaign"
	testdb "couponIssuanceSystem/tests/infra/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func newTestCampaign() *models.Campaign {
	return &models.Campaign{
		ID:         uuid.New(),
		Name:       "Test Campaign",
		TotalCount: 100,
		StartAt:    time.Now(),
	}
}

func TestCreateCampaign_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaign := newTestCampaign()

	err := repository.Create(context.Background(), campaign)
	assert.NoError(t, err)

	var found models.Campaign
	err = db.WithContext(context.Background()).First(&found, "id = ?", campaign.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, campaign.Name, found.Name)
}

func TestCreateCampaign_DuplicateID(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaign := newTestCampaign()

	err := repository.Create(context.Background(), campaign)
	assert.NoError(t, err)

	duplicateCampaign := *campaign
	duplicateCampaign.ID = campaign.ID

	err = repository.Create(context.Background(), &duplicateCampaign)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE")
}

func TestFindByID_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaign := newTestCampaign()

	err := repository.Create(context.Background(), campaign)
	assert.NoError(t, err)

	foundCampaign, err := repository.FindByID(context.Background(), campaign.ID)
	assert.NoError(t, err)
	assert.Equal(t, campaign.ID, foundCampaign.ID)
	assert.Equal(t, campaign.Name, foundCampaign.Name)
}

func TestFindByID_NotFound(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaignID := uuid.New()

	foundCampaign, err := repository.FindByID(context.Background(), campaignID)
	assert.Error(t, err)
	assert.Nil(t, foundCampaign)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
