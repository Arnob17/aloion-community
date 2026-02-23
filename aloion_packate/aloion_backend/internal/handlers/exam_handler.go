package handlers

import (
	"net/http"
	"strconv"

	"github.com/arnob17/aloion_backend/internal/dto"
	"github.com/arnob17/aloion_backend/internal/models"
	"github.com/arnob17/aloion_backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	examService *services.ExamService
}

func NewExamHandler(examService *services.ExamService) *ExamHandler {
	return &ExamHandler{
		examService: examService,
	}
}

func (h *ExamHandler) CreateExam(c *gin.Context) {
	var req dto.CreateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := h.examService.CreateExam(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ToExamResponse(exam))
}

func (h *ExamHandler) GetExam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam ID"})
		return
	}

	exam, err := h.examService.GetExamByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToExamResponse(exam))
}

func (h *ExamHandler) GetAllExams(c *gin.Context) {
	exams, err := h.examService.GetAllExams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.ExamResponse, len(exams))
	for i, exam := range exams {
		responses[i] = ToExamResponse(&exam)
	}

	c.JSON(http.StatusOK, responses)
}

func (h *ExamHandler) AddResult(c *gin.Context) {
	var req dto.CreateExamResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.examService.AddExamResult(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ToExamResultResponse(result))
}

func (h *ExamHandler) GetMyResults(c *gin.Context) {
	userID, _ := c.Get("user_id")
	results, err := h.examService.GetUserResults(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.ExamResultResponse, len(results))
	for i, result := range results {
		responses[i] = ToExamResultResponse(&result)
	}

	c.JSON(http.StatusOK, responses)
}

func (h *ExamHandler) GetResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid result ID"})
		return
	}

	result, err := h.examService.GetResultByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToExamResultResponse(result))
}

func (h *ExamHandler) GetAllResults(c *gin.Context) {
	results, err := h.examService.GetAllResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.ExamResultResponse, len(results))
	for i, result := range results {
		responses[i] = ToExamResultResponse(&result)
	}

	c.JSON(http.StatusOK, responses)
}

func ToExamResponse(exam *models.Exam) dto.ExamResponse {
	return dto.ExamResponse{
		ID:          exam.ID,
		Title:       exam.Title,
		Description: exam.Description,
		CourseID:    exam.CourseID,
		BatchID:     exam.BatchID,
		MaxScore:    exam.MaxScore,
		ExamDate:    exam.ExamDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   exam.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToExamResultResponse(result *models.ExamResult) dto.ExamResultResponse {
	response := dto.ExamResultResponse{
		ID:         result.ID,
		ExamID:     result.ExamID,
		UserID:     result.UserID,
		Score:      result.Score,
		MaxScore:   result.MaxScore,
		Percentage: result.Percentage,
		Grade:      result.Grade,
		Feedback:   result.Feedback,
		ExamDate:   result.ExamDate.Format("2006-01-02T15:04:05Z07:00"),
		ResultDate: result.ResultDate.Format("2006-01-02T15:04:05Z07:00"),
	}

	if result.Exam.ID != 0 {
		examResp := ToExamResponse(&result.Exam)
		response.Exam = &examResp
	}

	return response
}
