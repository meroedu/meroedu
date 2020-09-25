package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	tagHTTP "github.com/meroedu/meroedu/internal/tag/delivery/http"
)

func TestGetAll(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)
	mockUCase := new(mocks.TagUseCase)
	mockList := make([]domain.Tag, 0)
	mockList = append(mockList, mockTag)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)

	mockUCase := new(mocks.TagUseCase)

	num := int(mockTag.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockTag, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tags/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateTag(t *testing.T) {
	mockTag := domain.Tag{
		Name:      "Title",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tempmockTag := mockTag
	tempmockTag.ID = 0
	mockUCase := new(mocks.TagUseCase)

	j, err := json.Marshal(tempmockTag)
	assert.NoError(t, err)

	mockUCase.On("CreateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/tags", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags")

	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.CreateTag(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateTag(t *testing.T) {
	mockTag := domain.Tag{
		ID:        124,
		Name:      "tag1",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tempmockTag := mockTag
	mockUCase := new(mocks.TagUseCase)
	j, err := json.Marshal(tempmockTag)
	assert.NoError(t, err)
	mockUCase.On("UpdateTag", mock.Anything, mock.AnythingOfType("*domain.Tag"), mock.AnythingOfType("int64")).Return(nil)
	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/tags/124", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags/:id")
	c.SetParamNames("id")
	c.SetParamValues("124")

	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.UpdateTag(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteTag(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)

	mockUCase := new(mocks.TagUseCase)

	num := int(mockTag.ID)

	mockUCase.On("DeleteTag", mock.Anything, int64(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/tags/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tags/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.DeleteTag(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetCourseTags(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)
	mockUCase := new(mocks.TagUseCase)
	mockList := make([]domain.Tag, 0)
	mockList = append(mockList, mockTag)
	courseID := 10
	mockUCase.On("GetCourseTags", mock.Anything, mock.AnythingOfType("int64")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags/course/"+strconv.Itoa(courseID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tags/course/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(courseID))
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetCourseTags(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateCourseTag(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	mockUCase.On("CreateCourseTag", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/tags/course/:course_id/:tag_id", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags/course/:course_id/:tag_id")
	c.SetParamNames("course_id", "tag_id")
	c.SetParamValues("1", "4")
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	if assert.NoError(t, handler.CreateCourseTag(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	mockUCase.AssertExpectations(t)
}

func TestDeleteCourseTag(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	mockUCase.On("DeleteCourseTag", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/tags/course/:course_id/:tag_id", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags/course/:course_id/:tag_id")
	c.SetParamNames("course_id", "tag_id")
	c.SetParamValues("1", "4")
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	if assert.NoError(t, handler.DeleteCourseTag(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	mockUCase.AssertExpectations(t)
}

func TestGetLessonTags(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)
	mockUCase := new(mocks.TagUseCase)
	mockList := make([]domain.Tag, 0)
	mockList = append(mockList, mockTag)
	LessonID := 10
	mockUCase.On("GetLessonTags", mock.Anything, mock.AnythingOfType("int64")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags/Lesson/"+strconv.Itoa(LessonID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tags/lesson/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(LessonID))
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetLessonTags(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateLessonTag(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	mockUCase.On("CreateLessonTag", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/tags/lesson/:lesson_id/:tag_id", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags/lesson/:lesson_id/:tag_id")
	c.SetParamNames("lesson_id", "tag_id")
	c.SetParamValues("1", "4")
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	if assert.NoError(t, handler.CreateLessonTag(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	mockUCase.AssertExpectations(t)
}

func TestDeleteLessonTag(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	mockUCase.On("DeleteLessonTag", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/tags/lesson/:lesson_id/:tag_id", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags/lesson/:lesson_id/:tag_id")
	c.SetParamNames("lesson_id", "tag_id")
	c.SetParamValues("1", "4")
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	if assert.NoError(t, handler.DeleteLessonTag(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	mockUCase.AssertExpectations(t)
}
