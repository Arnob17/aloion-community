package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/arnob17/aloion_backend/internal/database"
	"github.com/arnob17/aloion_backend/internal/dto"
	"github.com/arnob17/aloion_backend/internal/models"
	"gorm.io/gorm"
)

type ExamService struct{}

func NewExamService() *ExamService {
	return &ExamService{}
}

func (s *ExamService) CreateExam(req dto.CreateExamRequest) (*models.Exam, error) {
	exam := models.Exam{
		Title:       req.Title,
		Description: req.Description,
		CourseID:    req.CourseID,
		BatchID:     req.BatchID,
		MaxScore:    req.MaxScore,
		ExamDate:    req.ExamDate,
	}

	if err := database.DB.Create(&exam).Error; err != nil {
		return nil, fmt.Errorf("failed to create exam: %w", err)
	}

	return &exam, nil
}

func (s *ExamService) GetExamByID(id uint) (*models.Exam, error) {
	var exam models.Exam
	if err := database.DB.Preload("Course").First(&exam, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("exam not found")
		}
		return nil, err
	}
	return &exam, nil
}

func (s *ExamService) GetAllExams() ([]models.Exam, error) {
	var exams []models.Exam
	if err := database.DB.Preload("Course").Order("exam_date DESC").Find(&exams).Error; err != nil {
		return nil, err
	}
	return exams, nil
}

func calculateGrade(percentage float64) string {
	if percentage >= 80 {
		return "A+"
	}
	if percentage >= 70 {
		return "A"
	}
	if percentage >= 60 {
		return "A-"
	}
	if percentage >= 50 {
		return "B"
	}
	if percentage >= 40 {
		return "C"
	}
	return "F"
}

func (s *ExamService) AddExamResult(req dto.CreateExamResultRequest) (*models.ExamResult, error) {
	// Verify exam exists
	var exam models.Exam
	if err := database.DB.First(&exam, req.ExamID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("exam not found")
		}
		return nil, err
	}

	// Verify user exists
	var user models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if result already exists
	var existingResult models.ExamResult
	if err := database.DB.Where("exam_id = ? AND user_id = ?", req.ExamID, req.UserID).First(&existingResult).Error; err == nil {
		return nil, errors.New("result already exists for this exam and user")
	}

	// Calculate percentage
	percentage := (req.Score / req.MaxScore) * 100
	grade := calculateGrade(percentage)

	result := models.ExamResult{
		ExamID:     req.ExamID,
		UserID:     req.UserID,
		Score:      req.Score,
		MaxScore:   req.MaxScore,
		Percentage: percentage,
		Grade:      grade,
		Feedback:   req.Feedback,
		ExamDate:   req.ExamDate,
		ResultDate: time.Now(),
	}

	if err := database.DB.Create(&result).Error; err != nil {
		return nil, fmt.Errorf("failed to create exam result: %w", err)
	}

	// Reload with relations
	if err := database.DB.Preload("Exam").Preload("User").First(&result, result.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *ExamService) GetUserResults(userID uint) ([]models.ExamResult, error) {
	var results []models.ExamResult
	if err := database.DB.
		Preload("Exam").
		Preload("User").
		Where("user_id = ?", userID).
		Order("exam_date DESC").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (s *ExamService) GetResultByID(id uint) (*models.ExamResult, error) {
	var result models.ExamResult
	if err := database.DB.Preload("Exam").Preload("User").First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("result not found")
		}
		return nil, err
	}
	return &result, nil
}

func (s *ExamService) GetAllResults() ([]models.ExamResult, error) {
	var results []models.ExamResult
	if err := database.DB.
		Preload("Exam").
		Preload("User").
		Order("exam_date DESC").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
