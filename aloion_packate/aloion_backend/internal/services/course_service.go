package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/arnob17/aloion_backend/internal/database"
	"github.com/arnob17/aloion_backend/internal/dto"
	"github.com/arnob17/aloion_backend/internal/models"
	// "gorm.io/gorm"
)

type CourseService struct{}

func NewCourseService() *CourseService {
	return &CourseService{}
}

func (s *CourseService) Create(req dto.CreateCourseRequest, teacherID uint) (*models.Course, error) {
	course := models.Course{
		Title:       req.Title,
		Description: req.Description,
		Subject:     req.Subject,
		Class:       req.Class,
		Level:       models.CourseLevel(req.Level),
		Duration:    req.Duration,
		Price:       req.Price,
		IsFree:      req.IsFree,
		TeacherID:   teacherID,
		Status:      models.CourseStatusDraft,
		Language:    req.Language,
	}

	if course.Language == "" {
		course.Language = "bn"
	}

	if err := database.DB.Create(&course).Error; err != nil {
		return nil, fmt.Errorf("failed to create course: %w", err)
	}

	return &course, nil
}

func (s *CourseService) GetByID(id uint) (*models.Course, error) {
	var course models.Course
	if err := database.DB.Preload("Teacher").First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *CourseService) GetAll(limit, offset int, filters map[string]interface{}) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	query := database.DB.Model(&models.Course{})

	// Apply filters
	if subject, ok := filters["subject"].(string); ok && subject != "" {
		query = query.Where("subject = ?", subject)
	}
	if class, ok := filters["class"].(int); ok && class > 0 {
		query = query.Where("class = ?", class)
	}
	if level, ok := filters["level"].(string); ok && level != "" {
		query = query.Where("level = ?", level)
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if isFree, ok := filters["is_free"].(bool); ok {
		query = query.Where("is_free = ?", isFree)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("Teacher").Limit(limit).Offset(offset).Order("created_at DESC").Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

func (s *CourseService) Update(id uint, teacherID uint, req dto.UpdateCourseRequest) (*models.Course, error) {
	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		return nil, err
	}

	// Check if user is the teacher or admin
	if course.TeacherID != teacherID {
		// Check if user is admin
		var user models.User
		if err := database.DB.First(&user, teacherID).Error; err != nil {
			return nil, errors.New("unauthorized")
		}
		if user.Role != models.RoleAdmin {
			return nil, errors.New("unauthorized: only course teacher or admin can update")
		}
	}

	updates := make(map[string]interface{})
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.Discount != nil {
		updates["discount"] = *req.Discount
	}

	if err := database.DB.Model(&course).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (s *CourseService) Delete(id uint, userID uint) error {
	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		return err
	}

	// Check authorization
	if course.TeacherID != userID {
		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			return errors.New("unauthorized")
		}
		if user.Role != models.RoleAdmin {
			return errors.New("unauthorized")
		}
	}

	return database.DB.Delete(&course).Error
}

func (s *CourseService) Enroll(userID, courseID uint) (*models.Enrollment, error) {
	// Check if already enrolled
	var existing models.Enrollment
	if err := database.DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existing).Error; err == nil {
		return nil, errors.New("already enrolled in this course")
	}

	// Get course
	var course models.Course
	if err := database.DB.First(&course, courseID).Error; err != nil {
		return nil, errors.New("course not found")
	}

	// Check if course is published
	if course.Status != models.CourseStatusPublished {
		return nil, errors.New("course is not available for enrollment")
	}

	// Create enrollment
	enrollment := models.Enrollment{
		UserID:     userID,
		CourseID:   courseID,
		Status:     models.EnrollmentStatusActive,
		EnrolledAt: time.Now(),
	}

	if err := database.DB.Create(&enrollment).Error; err != nil {
		return nil, fmt.Errorf("failed to enroll: %w", err)
	}

	return &enrollment, nil
}

func (s *CourseService) GetUserEnrollments(userID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	if err := database.DB.Where("user_id = ?", userID).
		Preload("Course").Preload("Course.Teacher").
		Order("created_at DESC").
		Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}
