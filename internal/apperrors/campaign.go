package apperrors

import "errors"

var (
	ErrCampaignNotFound     = errors.New("campaign not found")
	ErrInvalidCampaignInput = errors.New("invalid campaign input")
)
