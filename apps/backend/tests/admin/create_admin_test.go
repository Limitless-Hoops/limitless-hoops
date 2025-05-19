package admin_test

import (
	"bytes"
	"encoding/json"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateAdmin_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	payload := dto.CreateAdminDTO{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "admin@example.com",
		PhoneNumber: "+11234567890",
		Password:    "securepassword",
		Role:        "admin",
		DateOfBirth: time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC),
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admins", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
	assert.Equal(t, payload.Email, response["email"])
	assert.Equal(t, payload.FirstName, response["first_name"])
}

func TestCreateAdmin_MissingRequiredFields(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	// Empty body
	payload := dto.CreateAdminDTO{}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admins", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)
	assert.Contains(t, response, "validation_error")
}

func TestCreateAdmin_InvalidJSON(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	body := []byte(`{invalid_json`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admins", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)
	assert.Contains(t, response, "error")
}
