package domain

import (
	"context"
	"interactions/internal/domain/entity"
	"interactions/internal/dto"
)

// InteractionRepository - интерфейс для взаимодействий
type InteractionRepository interface {
	CreateInteraction(ctx context.Context, interaction entity.Interaction) error
	GetInteractions(ctx context.Context, adID string) ([]entity.Interaction, error)
}

type InteractionUsecase interface {
	AddInteraction(ctx context.Context, req *dto.InteractionDTO) (*dto.InteractionResponseDTO, error)
	GetInteraction(ctx context.Context, req *dto.GetInteractionRequestDTO) (*dto.GetInteractionResponseDTO, error)
}
