package coupon

import (
	"connectrpc.com/connect"
	"context"
	couponv1 "couponIssuanceSystem/gen/coupon/v1"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/service/coupon"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	service coupon.Service
}

func NewHandler(service coupon.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) IssueCoupon(
	ctx context.Context,
	req *connect.Request[couponv1.IssueCouponRequest],
) (*connect.Response[couponv1.IssueCouponResponse], error) {
	msg := req.Msg

	campaignID, err := uuid.Parse(msg.GetCampaignId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	userID := msg.GetUserId()
	if userID == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, apperrors.ErrInvalidUserID)
	}

	result, err := h.service.IssueCoupon(ctx, campaignID, userID)
	if err != nil {
		switch err {
		case apperrors.ErrCampaignNotFound:
			return nil, connect.NewError(connect.CodeNotFound, err)
		case apperrors.ErrCampaignNotStarted, apperrors.ErrCampaignEnded:
			return nil, connect.NewError(connect.CodeFailedPrecondition, err)
		case apperrors.ErrUserAlreadyIssued:
			return nil, connect.NewError(connect.CodeAlreadyExists, err)
		case apperrors.ErrCampaignSoldOut:
			return nil, connect.NewError(connect.CodeResourceExhausted, err)
		case apperrors.ErrCouponCodeConflict:
			return nil, connect.NewError(connect.CodeAborted, err)
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	resp := &couponv1.IssueCouponResponse{
		Code:       result.Code,
		CampaignId: result.CampaignID,
		UserId:     result.UserID,
		IssuedAt:   timestamppb.New(result.IssuedAt),
	}

	return connect.NewResponse(resp), nil
}
