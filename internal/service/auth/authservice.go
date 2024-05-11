package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mailersend/mailersend-go"
	"github.com/redis/go-redis/v9"
	"net/http"
	"server/internal/model"
	"server/internal/pkg"
	"server/internal/repository/auth"
	"time"
)

type AuthService struct {
	repo      auth.IAuthRepo
	ms        *mailersend.Mailersend
	redis     *redis.Client
	secretKey string
}

type IAuthService interface {
	CreateNewUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	DeleteUserByEmail(email string) error
	CheckUserCreds(creds model.User) (model.User, error)
	JwtAuthorization(user model.User) (string, error)
	SendEmailCode(email string, ctx context.Context) error
	CheckCode(email, code string, ctx context.Context) error
	//Login(user models.User) (string, error)
}

func NewAuthService(repo auth.IAuthRepo, secretKey string, mailsenderKey string, client *redis.Client) *AuthService {
	ms := mailersend.NewMailersend(mailsenderKey)

	return &AuthService{
		redis:     client,
		repo:      repo,
		secretKey: secretKey,
		ms:        ms,
	}
}

const (
	VerificationCodeSubject = "Verification code"
)

func validateUserData(user model.User) error {
	if !pkg.IsValid(user) {
		return model.ErrEmptyness
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

	if err := a.sendEmail(email, code, ctx); err != nil {
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

func (a AuthService) sendEmail(email, text string, ctx context.Context) error {
	from := mailersend.From{
		Name:  "Kamal",
		Email: "foreverwantlive@gmail.com",
	}

	to := []mailersend.Recipient{
		{
			Name:  "Client",
			Email: email,
		},
	}

	//sendAt := time.Now().Add(time.Second * 30).Unix()

	msg := a.ms.Email.NewMessage()

	msg.SetFrom(from)
	msg.SetRecipients(to)
	msg.SetSubject(VerificationCodeSubject)
	msg.SetText(text)

	res, err := a.ms.Email.Send(ctx, msg)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return errors.New("mailsender gave an error")
	}

	return nil
}
