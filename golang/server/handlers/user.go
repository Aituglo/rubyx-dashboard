package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/aituglo/rubyx-dashboard/golang/db"
	"github.com/aituglo/rubyx-dashboard/golang/env"
	"github.com/aituglo/rubyx-dashboard/golang/errors"
	"github.com/aituglo/rubyx-dashboard/golang/server/write"
	"golang.org/x/crypto/bcrypt"
)

var src = rand.NewSource(time.Now().UnixNano())

type ChangePasswordForm struct {
	PreviousPass string `json:"previousPass"`
	NewPass      string `json:"newPass"`
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// generateRandomString can be used to create verification codes or something this
// this implementation comes from stackoverflow
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func generateRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func hashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

func checkPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

func UpdatePassword(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user == nil {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	changeForm := &ChangePasswordForm{}
	err := decoder.Decode(&changeForm)
	if err != nil || changeForm == nil {
		return write.Error(errors.NoJSONBody)
	}

	currentUser, err := env.DB().FindUserByID(r.Context(), user.ID)
	if err != nil {
		return write.Error(errors.UserNotFound)
	}

	if !checkPasswordHash(changeForm.PreviousPass, currentUser.Salt, currentUser.Pass) {
		return write.Error(errors.PreviousPassIncorrect)
	}

	// salt and hash it
	salt := generateRandomString(32)
	pass, err := hashPassword(changeForm.NewPass, salt)
	if err != nil {
		return write.Error(err)
	}

	err = env.DB().UpdateUserPassword(r.Context(), db.UpdateUserPasswordParams{
		ID:   user.ID,
		Pass: pass,
		Salt: salt,
	})

	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}

func UpdateEmail(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user == nil {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	u := &db.User{}
	err := decoder.Decode(&u)
	if err != nil || u == nil {
		return write.Error(errors.NoJSONBody)
	}

	currentUser, err := env.DB().FindUserByID(r.Context(), user.ID)
	if err != nil {
		return write.Error(errors.UserNotFound)
	}

	if !checkPasswordHash(u.Pass, currentUser.Salt, currentUser.Pass) {
		return write.Error(errors.PreviousPassIncorrect)
	}

	err = env.DB().UpdateUserEmail(r.Context(), db.UpdateUserEmailParams{
		ID:    user.ID,
		Email: u.Email,
	})

	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}

func Whoami(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSON(user)
}
