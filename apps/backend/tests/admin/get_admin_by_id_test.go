package admin_test

import (
	"encoding/json"
	"fmt"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAdminByID_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	dob := time.Date(1985, time.January, 1, 0, 0, 0, 0, time.UTC)
	admin := models.Admin{
		FirstName:    "Bob",
		LastName:     "Builder",
		Email:        "bob@example.com",
		PhoneNumber:  "2223334444",
		PasswordHash: "hashed",
		Role:         "staff",
		DateOfBirth:  &dob,
	}
	require.NoError(t, db.Create(&admin).Error)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/admins/%d", admin.ID), nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.Equal(t, "Bob", body["first_name"])
	assert.Equal(t, "bob@example.com", body["email"])
	assert.Equal(t, float64(admin.ID), body["id"])
}

func TestGetAdminByID_NotFound(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/admins/9999", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var body map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&body)
	assert.Contains(t, body["error"], "not found")
}

func TestGetAdminByID_InvalidID(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/admins/abc", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var body map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&body)
	assert.Contains(t, body["error"], "Invalid admin ID")
}
