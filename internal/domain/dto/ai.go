// Package dto defines standarized struct to be used as data exchange
package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type AnalyzeCV struct {
	ID                       uuid.UUID            `json:"id"`
	UserID                   uuid.UUID            `json:"user_id"`
	JobTitle                 string               `json:"job_title" validate:"required"`
	TargetCompany            string               `json:"target_company" validate:"required"`
	Industry                 string               `json:"industry" validate:"required"`
	WorkExperience           string               `json:"work_experience" validate:"required"`
	HighestEducation         string               `json:"highest_education" validate:"required"`
	EmploymentStatus         string               `json:"employment_status" validate:"required"`
	PrimarySkill             string               `json:"primary_skill" validate:"required"`
	Tools                    string               `json:"tools" validate:"required"`
	SpokenAndWrittenLanguage string               `json:"spoken_and_written_language" validate:"required"`
	PrimaryAnalysisGoals     string               `json:"primary_analysis_goals" validate:"required"`
	JobApplicationsSent      uint32               `json:"job_applications_sent" validate:"required"`
	BiggestConcern           string               `json:"biggest_concern" validate:"required"`
	AdditionalRequest        string               `json:"AdditionalRequest"`
	File                     multipart.FileHeader `json:"file"`
}

type ResponseAnalyzeCV struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Response string    `json:"response"`
}

type PayloadAnalyzeCV struct {
	Model string           `json:"model"`
	Input []InputAnalyzeCV `json:"input"`
}
type ContentAnalyzeCV struct {
	Type   string `json:"type"`
	FileID string `json:"file_id,omitempty"`
	Text   string `json:"text,omitempty"`
}
type InputAnalyzeCV struct {
	Role    string             `json:"role"`
	Content []ContentAnalyzeCV `json:"content"`
}
