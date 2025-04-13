package apperrors

import "errors"

var (
	ErrCampaignNotFound     = errors.New("campaign not found")
	ErrInvalidCampaignInput = errors.New("invalid campaign input")
	ErrCampaignNotStarted   = errors.New("campaign not started")
	ErrCampaignEnded        = errors.New("campaign ended")
	ErrCampaignSoldOut      = errors.New("campaign sold out")
)
