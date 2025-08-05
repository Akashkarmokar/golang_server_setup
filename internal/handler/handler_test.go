package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Akashkarmokar/golang_server_setup/internal/handler"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.PostNew()(w, r)

			// Assert

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}

}

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.GetAllNews()(w, r)

			// Assert

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}

}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.GetNewsByID()(w, r)

			// Assert

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}

}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.UpdateNewsByID()(w, r)

			// Assert

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}

}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.DeleteNewsByID()(w, r)

			// Assert

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}

}
