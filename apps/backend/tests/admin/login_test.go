package testadmin

import (
	"bytes"
	"encoding/json"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testsetup"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminLogin_Success(t *testing.T) {
	// Setup
	app := testsetup.NewTestApp()
	db := testsetup.ConnectTestDB()
	require.NoError(t, testsetup.ResetTestDB(db))

	// Create test admin directly in DB
	admin := models.Admin{
		Email:    "admin@example.com",
		Password: utils.UUIDv4(), // raw password for login request
	}
	hashed, _ := models.HashPassword(admin.Password)
	admin.Password = hashed
	require.NoError(t, db.Create(&admin).Error)

	// Prepare login payload
	payload, _ := json.Marshal(map[string]string{
		"email":    "admin@example.com",
		"password": admin.Password, // use unhashed here
	})

	// Perform request
	req := httptest.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response body
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	_, ok := body["token"]
	assert.True(t, ok, "response should contain a token")
}
