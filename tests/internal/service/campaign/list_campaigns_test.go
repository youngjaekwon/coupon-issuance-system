package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/campaign"
	svc "couponIssuanceSystem/internal/service/campaign"
	testdb "couponIssuanceSystem/tests/infra/db"
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func createCampaign(repository repo.Repository, name string, totalCount int, startAt time.Time) (*models.Campaign, error) {
	campaign := &models.Campaign{
		Name:       name,
		TotalCount: totalCount,
		StartAt:    startAt,
	}
	err := repository.Create(context.Background(), campaign)
	return campaign, err
}

func TestListCampaigns_Success(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	campaigns := make([]*models.Campaign, 0)

	for i := 0; i < 5; i++ {
		campaign, err := createCampaign(repository, "Campaign "+strconv.Itoa(i), 100, time.Now().Add(1*time.Hour))
		if err != nil {
			t.Fatalf("Failed to create campaign: %v", err)
		}
		campaigns = append(campaigns, campaign)
	}

	gotCampaigns, err := service.ListCampaigns(context.Background(), 1, 5)

	assert.NoError(t, err)
	assert.Len(t, gotCampaigns, 5)
	for i, campaign := range gotCampaigns {
		assert.Equal(t, campaigns[i].Name, campaign.Name)
		assert.Equal(t, 100, campaign.TotalCount)
		assert.True(t, campaign.StartAt.After(time.Now()))
	}
}

func TestListCampaigns_Empty(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	gotCampaigns, err := service.ListCampaigns(context.Background(), 1, 5)

	assert.NoError(t, err)
	assert.Empty(t, gotCampaigns)
}

func TestListCampaigns_InvalidPageOrLimit(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	_, err := service.ListCampaigns(context.Background(), -1, 5)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrInvalidPage))

	_, err = service.ListCampaigns(context.Background(), 1, -5)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrInvalidPageSize))
}

func TestListCampaigns_Pagination(t *testing.T) {
	db := testdb.NewTestDB()
	repository := repo.New(db)
	service := svc.New(repository)

	for i := 0; i < 25; i++ {
		_, err := createCampaign(repository, "Campaign "+strconv.Itoa(i), 100, time.Now().Add(1*time.Hour))
		if err != nil {
			t.Fatalf("Failed to create campaign: %v", err)
		}
	}

	gotCampaigns, err := service.ListCampaigns(context.Background(), 1, 10)
	assert.NoError(t, err)
	assert.Len(t, gotCampaigns, 10)

	gotCampaigns, err = service.ListCampaigns(context.Background(), 2, 10)
	assert.NoError(t, err)
	assert.Len(t, gotCampaigns, 10)

	gotCampaigns, err = service.ListCampaigns(context.Background(), 3, 10)
	assert.NoError(t, err)
	assert.Len(t, gotCampaigns, 5)
}
