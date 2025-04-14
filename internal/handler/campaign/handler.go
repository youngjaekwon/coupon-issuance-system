package campaign

import (
	"connectrpc.com/connect"
	"context"
	campaignv1 "couponIssuanceSystem/gen/campaign/v1"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/service/campaign"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	service campaign.Service
}

func NewHandler(service campaign.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.CreateCampaignRequest],
) (*connect.Response[campaignv1.CreateCampaignResponse], error) {
	msg := req.Msg

	if msg.GetName() == "" || msg.GetTotalCount() <= 0 || msg.GetStartAt() == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, apperrors.ErrInvalidCampaignInput)
	}

	input := &campaign.CreateCampaignInput{
		Name:       msg.GetName(),
		TotalCount: int(msg.GetTotalCount()),
		StartAt:    msg.GetStartAt().AsTime(),
	}

	if msg.GetEndAt() != nil {
		endAt := msg.GetEndAt().AsTime()
		input.EndAt = &endAt
	}

	result, err := h.service.CreateCampaign(ctx, input)
	if err != nil {
		switch err {
		case apperrors.ErrInvalidCampaignInput:
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	resp := &campaignv1.CreateCampaignResponse{
		Id:         result.ID,
		Name:       result.Name,
		TotalCount: int32(result.TotalCount),
		StartAt:    timestamppb.New(result.StartAt),
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}

	if result.EndAt != nil {
		resp.EndAt = timestamppb.New(*result.EndAt)
	}

	return connect.NewResponse(resp), nil
}

func (h *Handler) GetCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.GetCampaignRequest],
) (*connect.Response[campaignv1.GetCampaignResponse], error) {
	campaignID, err := uuid.Parse(req.Msg.GetId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	result, err := h.service.FindCampaign(ctx, campaignID)
	if err != nil {
		switch err {
		case apperrors.ErrCampaignNotFound:
			return nil, connect.NewError(connect.CodeNotFound, err)
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	resp := &campaignv1.GetCampaignResponse{
		Id:         result.ID,
		Name:       result.Name,
		TotalCount: int32(result.TotalCount),
		Stock:      int32(result.Stock),
		StartAt:    timestamppb.New(result.StartAt),
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}

	if result.EndAt != nil {
		resp.EndAt = timestamppb.New(*result.EndAt)
	}

	for _, c := range result.Coupons {
		resp.Coupons = append(resp.Coupons, &campaignv1.Coupon{
			Code:     c.Code,
			UserId:   c.UserID,
			IssuedAt: timestamppb.New(c.IssuedAt),
		})
	}

	return connect.NewResponse(resp), nil
}
