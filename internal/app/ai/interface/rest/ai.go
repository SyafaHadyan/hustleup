// Package rest receive request from user and return appropriate response based on package usecase
package rest

import (
	"net/http"

	"github.com/SyafaHadyan/worku/internal/app/ai/usecase"
	"github.com/SyafaHadyan/worku/internal/domain/dto"
	"github.com/SyafaHadyan/worku/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

type AIHandler struct {
	Validator  *validator.Validate
	Decoder    *schema.Decoder
	Middleware middleware.MiddlewareItf
	AIUseCase  usecase.AIUseCaseItf
}

func NewAIHandler(
	routerGroup fiber.Router, validator *validator.Validate,
	decoder *schema.Decoder, middleware middleware.MiddlewareItf,
	aiUseCase usecase.AIUseCaseItf,
) {
	aiHandler := AIHandler{
		Validator:  validator,
		Decoder:    decoder,
		Middleware: middleware,
		AIUseCase:  aiUseCase,
	}

	routerGroup = routerGroup.Group("/ai")

	routerGroup.Post("/cv", middleware.Authentication, aiHandler.AnalyzeCV)
}

func (h *AIHandler) AnalyzeCV(ctx *fiber.Ctx) error {
	var analyzeCV dto.AnalyzeCV

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(
			http.StatusUnauthorized,
			"user unauthorized",
		)
	}

	// TODO: parse multipart form

	// multipartForm, err := ctx.MultipartForm()
	// if err != nil {
	// 	return fiber.NewError(
	// 		http.StatusBadRequest,
	// 		"failed to parse multipart form",
	// 	)
	// }

	// err = h.Decoder.Decode(&analyzeCV, multipartForm.Value)
	// if err != nil {
	// 	return fiber.NewError(
	// 		http.StatusBadRequest,
	// 		"failed to decode form fields",
	// 	)
	// }

	analyzeCV.UserID = userID

	// err = h.Validator.Struct(analyzeCV)
	// if err != nil {
	// 	return fiber.NewError(
	// 		http.StatusBadRequest,
	// 		"invalid request form",
	// 	)
	// }

	file, err := ctx.FormFile("document")
	if err != nil {
		return err
	}

	res, err := h.AIUseCase.AnalyzeCV(analyzeCV, *file)
	if err != nil {
		return fiber.NewError(
			http.StatusInternalServerError,
			"failed to analyze cv",
		)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully analyzed cv",
		"payload": res,
	})
}
