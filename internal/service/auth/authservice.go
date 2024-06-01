package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"server/internal/config"
	"server/internal/model"
	"server/internal/pkg"
	"server/internal/pkg/mailer"
	"server/internal/repository/auth"
	"time"
)

type AuthService struct {
	repo      auth.IAuthRepo
	ms        *mailer.Mailer
	redis     *redis.Client
	secretKey string
}

type IAuthService interface {
	CreateNewUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	GetUserById(id int) (model.User, error)
	DeleteUserByEmail(email string) error
	CheckUserCreds(creds model.User) (model.User, error)
	JwtAuthorization(user model.User) (string, error)
	SendEmailCode(email string, ctx context.Context) error
	CheckCode(email, code string, ctx context.Context) error
	UpdateUserByEmail(email string, req model.UpdateUserRequest) error
	SendMsg(to, msg string) error
}

func NewAuthService(repo auth.IAuthRepo, secretKey string, mailCfg config.Mailer, client *redis.Client) *AuthService {
	return &AuthService{
		redis:     client,
		repo:      repo,
		secretKey: secretKey,
		ms:        mailer.NewMailer(mailCfg),
	}
}

const (
	VerificationCodeSubject = "Verification code"
)

func validateUserData(user model.User) error {
	if !pkg.IsValid(user) {
		return model.ErrEmptiness
	}

	if !pkg.IsEmailValid(user.Email) {
		return model.ErrInvalidEmail
	}

	if !pkg.IsNameValid(user.FirstName, user.LastName) {
		return model.ErrInvalidName
	}

	if !pkg.IsPasswordStrong(user.Password) {
		return model.ErrInvalidPassword
	}

	return nil
}

func (a AuthService) CreateNewUser(user model.User) (err error) {
	if err = validateUserData(user); err != nil {
		return err
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return a.repo.CreateUser(user)
}

func (a AuthService) GetUserByEmail(email string) (model.User, error) {
	return a.repo.GetUserByEmail(email)
}

func (a AuthService) DeleteUserByEmail(email string) error {
	return a.repo.DeleteUserByEmail(email)
}

func (a AuthService) CheckUserCreds(creds model.User) (model.User, error) {
	user, err := a.repo.GetUserByEmail(creds.Email)
	if err != nil {
		return model.User{}, err
	}

	if !pkg.CheckPasswordHash(creds.Password, user.Password) {
		return model.User{}, model.ErrIncorrectEmailOrPassword
	}

	return user, nil
}

func (a AuthService) JwtAuthorization(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
			ID:        pkg.RandStringBytesMaskImpr(40),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a AuthService) SendEmailCode(email string, ctx context.Context) error {
	code, _ := pkg.GenerateCode(6)

	if err := a.redis.Set(ctx, email, code, 10*time.Minute).Err(); err != nil {
		return err
	}

	msg := fmt.Sprintf("Ваш код подтверждения: %s, если вы не должны были получать данный код, проигнорируйте данное сообщение", code)

	if err := a.ms.SendMessage(email, msg); err != nil {
		return err
	}

	return nil
}

func (a AuthService) CheckCode(email, code string, ctx context.Context) error {
	stringCmd := a.redis.Get(ctx, email)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}

	if stringCmd.Val() != code {
		return model.ErrIncorrectCode
	}

	return nil
}

func (a AuthService) GetUserById(id int) (model.User, error) {
	return a.repo.GetUserById(id)
}

func (a AuthService) UpdateUserByEmail(email string, req model.UpdateUserRequest) error {
	u, err := a.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if req.HasPassword() {
		req = req.SetPassword(req.Password)
	}

	if err := a.repo.UpdateUser(req.SetId(u.Id)); err != nil {
		return err
	}

	return nil
}

func (a AuthService) SendMsg(to, msg string) error {
	return a.ms.SendMessage(to, msg)
}
