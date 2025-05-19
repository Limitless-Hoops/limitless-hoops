package admin_test

import (
	"bytes"
	"encoding/json"
	"github.com/Limitless-Hoops/limitless-hoops/dto"
	"github.com/Limitless-Hoops/limitless-hoops/models"
	"github.com/Limitless-Hoops/limitless-hoops/tests/testutil"
	"github.com/Limitless-Hoops/limitless-hoops/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAdminLogin_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	rawPassword := "password123"
	hashed, _ := utilities.HashPassword(rawPassword)
	admin := models.Admin{
		FirstName:    "Jane",
		LastName:     "Doe",
		Email:        "admin@example.com",
		PhoneNumber:  "1234567890",
		PasswordHash: hashed,
		Role:         "admin",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-30, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	payload, _ := json.Marshal(dto.LoginDTO{
		Email:    admin.Email,
		Password: rawPassword,
		Role:     admin.Role,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.NotEmpty(t, body["token"])
}

func TestAdminLogin_InvalidPassword(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	hashed, _ := utilities.HashPassword("correct-password")
	admin := models.Admin{
		Email:        "admin2@example.com",
		PasswordHash: hashed,
		FirstName:    "Wrong",
		LastName:     "Pass",
		PhoneNumber:  "0000000000",
		Role:         "admin",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-25, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	payload, _ := json.Marshal(dto.LoginDTO{
		Email:    admin.Email,
		Password: "wrong-password",
		Role:     "admin",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAdminLogin_EmailNotFound(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	payload, _ := json.Marshal(dto.LoginDTO{
		Email:    "ghost@example.com",
		Password: "irrelevant",
		Role:     "admin",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAdminMe_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	rawPassword := "me-password"
	hashed, _ := utilities.HashPassword(rawPassword)
	admin := models.Admin{
		FirstName:    "Token",
		LastName:     "User",
		Email:        "me@example.com",
		PhoneNumber:  "9998887777",
		PasswordHash: hashed,
		Role:         "admin",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-35, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	token, err := utilities.GenerateJWT(admin.ID, "admin")
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var profile map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&profile))
	assert.Equal(t, "me@example.com", profile["email"])
}

func TestAdminMe_MissingToken(t *testing.T) {
	app := testutil.NewTestApp()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAdminUpdatePassword_Success(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	original := "oldpassword"
	updated := "newsecurepassword"
	hashed, _ := utilities.HashPassword(original)

	admin := models.Admin{
		Email:        "changepass@example.com",
		PasswordHash: hashed,
		FirstName:    "Change",
		LastName:     "Pass",
		PhoneNumber:  "1112223333",
		Role:         "admin",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-29, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	token, err := utilities.GenerateJWT(admin.ID, "admin")
	require.NoError(t, err)

	payload, _ := json.Marshal(dto.UpdatePasswordDTO{
		OldPassword: original,
		NewPassword: updated,
	})

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/auth/password", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Confirm login with a new password
	loginPayload, _ := json.Marshal(dto.LoginDTO{
		Email:    admin.Email,
		Password: updated,
		Role:     "admin",
	})
	req = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(loginPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAdminUpdatePassword_InvalidOldPassword(t *testing.T) {
	app := testutil.NewTestApp()
	db := testutil.ConnectTestDB()
	t.Cleanup(func() { require.NoError(t, testutil.ResetTestDB(db)) })

	hashed, _ := utilities.HashPassword("correct-old")

	admin := models.Admin{
		Email:        "failchange@example.com",
		PasswordHash: hashed,
		FirstName:    "Wrong",
		LastName:     "Old",
		PhoneNumber:  "4445556666",
		Role:         "admin",
		DateOfBirth:  testutil.DatePtr(time.Now().AddDate(-33, 0, 0)),
	}
	require.NoError(t, db.Create(&admin).Error)

	token, _ := utilities.GenerateJWT(admin.ID, "admin")

	payload, _ := json.Marshal(dto.UpdatePasswordDTO{
		OldPassword: "incorrect",
		NewPassword: "doesntmatter",
	})

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/auth/password", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
