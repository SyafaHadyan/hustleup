// Package usecase handles the logic for each user request
package usecase

import (
	"context"
	"mime/multipart"

	"github.com/SyafaHadyan/worku/internal/app/ai/repository"
	"github.com/SyafaHadyan/worku/internal/domain/dto"
	"github.com/SyafaHadyan/worku/internal/domain/entity"
	aiitf "github.com/SyafaHadyan/worku/internal/infra/ai"
	"github.com/google/uuid"
)

type AIUseCaseItf interface {
	AnalyzeCV(analyzeCV dto.AnalyzeCV, file multipart.FileHeader) (dto.ResponseAnalyzeCV, error)
}

type AIUseCase struct {
	aiRepo    repository.AIDBItf
	ai        aiitf.AIItf
	aiContext context.Context
}

func NewAIUseCase(
	aiRepo repository.AIDBItf, ai aiitf.AIItf,
) AIUseCaseItf {
	return &AIUseCase{
		aiRepo:    aiRepo,
		ai:        ai,
		aiContext: context.Background(),
	}
}

func (u *AIUseCase) AnalyzeCV(analyzeCV dto.AnalyzeCV, file multipart.FileHeader) (dto.ResponseAnalyzeCV, error) {
	response, err := u.ai.AnalyzeCV(u.aiContext, analyzeCV, &file)
	if err != nil {
		return dto.ResponseAnalyzeCV{}, err
	}

	responseAnalyzeCV := entity.ResponseAnalyzeCV{
		ID:       uuid.New(),
		UserID:   analyzeCV.UserID,
		Response: response,
	}

	err = u.aiRepo.ResponseAnalyzeCV(&responseAnalyzeCV)

	return responseAnalyzeCV.ParseToDTOResponseAnalyzeCV(), err
}
