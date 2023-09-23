package service

import (
	"context"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	posgtesRepo repository.Postgres
}

func NewAuthService(repo *repository.Repository) AuthService {
	return &authService{}
}

func (s *authService) Login(ctx context.Context, creds data.UserSignInRequest) (resp data.Tokens, err error) {
	return resp, apperror.NewErrorInfo(ctx, errcodes.IncorrectCredsError, "HAHA")
}

// hashAndSalt - hashes the password with salt. Function takes password as []byte and returns the hash as string and error.
func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash - checks if the password matches the hash. Function takes password and has as string, and returns true if they matched and false otherwise.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
