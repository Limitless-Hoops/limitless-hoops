package admin_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUpdateAdmin_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	dob := time.Date(1980, 5, 1, 0, 0, 0, 0, time.UTC)
	admin := models.Admin{
		FirstName:    "OldFirst",
		LastName:     "OldLast",
		Email:        "old@example.com",
		PhoneNumber:  "5550001234",
		PasswordHash: "hashed",
		Role:         "staff",
		DateOfBirth:  &dob,
	}
	require.NoError(t, db.Create(&admin).Error)

	newFirst := "NewFirst"
	newEmail := "new@example.com"

	payload := dto.UpdateAdminDTO{
		FirstName: &newFirst,
		Email:     &newEmail,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/admins/%d", admin.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Re-fetch to confirm an update
	var updated models.Admin
	require.NoError(t, db.First(&updated, admin.ID).Error)
	assert.Equal(t, newFirst, updated.FirstName)
	assert.Equal(t, newEmail, updated.Email)
}

func TestUpdateAdmin_NotFound(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	payload := dto.UpdateAdminDTO{
		FirstName: testutil.StrPtr("Ghost"),
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/admins/9999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)
	assert.Contains(t, response["error"], "record not found")
}

func TestUpdateAdmin_InvalidID(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	payload := dto.UpdateAdminDTO{
		FirstName: testutil.StrPtr("Invalid"),
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/admins/abc", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)
	assert.Contains(t, response["error"], "Invalid admin ID")
}

func TestUpdateAdmin_InvalidPayload(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	admin := models.Admin{
		FirstName:    "Valid",
		LastName:     "Admin",
		Email:        "valid@example.com",
		PhoneNumber:  "+11234567890",
		PasswordHash: "hashed",
		Role:         "staff",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-30, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	invalid := dto.UpdateAdminDTO{
		FirstName: testutil.StrPtr("x"), // too short, fails validation
	}
	body, _ := json.Marshal(invalid)

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/admins/%d", admin.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)
	assert.Contains(t, response, "validation_error")
}
