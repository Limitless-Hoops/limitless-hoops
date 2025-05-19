package admin_test

import (
	"encoding/json"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAdmins_ReturnsAllWithDependentCounts(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	dob := time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC)
	admin := models.Admin{
		FirstName:    "Alice",
		LastName:     "Anderson",
		Email:        "alice@example.com",
		PhoneNumber:  "1112223333",
		PasswordHash: "hashed", // not validated in this test
		Role:         "admin",
		DateOfBirth:  &dob,
	}
	require.NoError(t, db.Create(&admin).Error)

	// Seed 2 valid dependents
	childDOB := time.Date(2010, time.June, 1, 0, 0, 0, 0, time.UTC)
	email1 := "child1@example.com"
	email2 := "child2@example.com"
	phone1 := "9999999991"
	phone2 := "9999999992"

	dependents := []models.Dependent{
		{
			FirstName:    "Child1",
			LastName:     "Anderson",
			Email:        &email1,
			PhoneNumber:  &phone1,
			PasswordHash: "hashed",
			DateOfBirth:  &childDOB,
			AdminID:      &admin.ID,
		},
		{
			FirstName:    "Child2",
			LastName:     "Anderson",
			Email:        &email2,
			PhoneNumber:  &phone2,
			PasswordHash: "hashed",
			DateOfBirth:  &childDOB,
			AdminID:      &admin.ID,
		},
	}
	require.NoError(t, db.Create(&dependents).Error)

	// Perform request
	req := httptest.NewRequest(http.MethodGet, "/api/v1/admins", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode and assert response
	var body []map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Len(t, body, 1)

	adminResp := body[0]
	assert.Equal(t, float64(2), adminResp["dependent_count"])
	assert.Equal(t, "Alice", adminResp["first_name"])
	assert.Equal(t, "alice@example.com", adminResp["email"])
}

func TestGetAdmins_EmptyList(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() {
		require.NoError(t, testutil.ResetTestDB(db))
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/admins", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body []map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.Len(t, body, 0)
}
