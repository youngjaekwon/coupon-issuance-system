package campaign

import (
	"context"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/campaign"
	testdb "couponIssuanceSystem/tests/infra/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strconv"
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

func TestListCampaigns_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaign1 := newTestCampaign()
	campaign2 := newTestCampaign()
	campaign2.Name = "Another Campaign"

	err := repository.Create(context.Background(), campaign1)
	assert.NoError(t, err)
	err = repository.Create(context.Background(), campaign2)
	assert.NoError(t, err)

	campaigns, err := repository.List(context.Background(), 0, 10)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 2)
}

func TestListCampaigns_Empty(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaigns, err := repository.List(context.Background(), 0, 10)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 0)
}

func TestListCampaigns_Pagination(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	for i := 0; i < 25; i++ {
		campaign := newTestCampaign()
		campaign.Name = "Campaign " + strconv.Itoa(i)
		err := repository.Create(context.Background(), campaign)
		assert.NoError(t, err)
	}

	campaigns, err := repository.List(context.Background(), 0, 10)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 10)

	campaigns, err = repository.List(context.Background(), 1, 10)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 10)

	campaigns, err = repository.List(context.Background(), 2, 10)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 5)
}

func TestListCampaigns_InvalidPage(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	campaigns, err := repository.List(context.Background(), -1, 10)
	assert.Error(t, err)
	assert.Nil(t, campaigns)

	campaigns, err = repository.List(context.Background(), 0, -1)
	assert.Error(t, err)
	assert.Nil(t, campaigns)
}

func TestListCampaigns_NoLimit(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)

	for i := 0; i < 25; i++ {
		campaign := newTestCampaign()
		campaign.Name = "Campaign " + strconv.Itoa(i)
		err := repository.Create(context.Background(), campaign)
		assert.NoError(t, err)
	}

	campaigns, err := repository.List(context.Background(), 0, 0)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 25)
}
