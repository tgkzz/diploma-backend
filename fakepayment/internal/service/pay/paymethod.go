package pay

import (
	"encoding/json"
	"errors"
	"fakepayment/internal/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"net/url"
)

func (p *PayService) BuyCourse(input model.ClientInput) error {
	token, err := jwt.ParseWithClaims(input.JwtToken, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing method")
		}

		return []byte(p.secretKey), nil
	})
	//error is here
	if err != nil {
		return err
	}

	var email string
	if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
		email = claims.Email
	}
	serviceURL := "http://localhost:8181/auth/getUserByEmail"
	queryParams := url.Values{}
	queryParams.Set("email", email)
	serviceURLWithParams := fmt.Sprintf("%s?%s", serviceURL, queryParams.Encode())

	resp, err := http.Get(serviceURLWithParams)
	if err != nil {
		return fmt.Errorf("error making request to auth service: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var userResponse model.UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}

	//-------------------

	contentUrl := fmt.Sprintf("http://localhost:8282/course/name/%s", url.PathEscape(input.CourseName))
	respContent, err := http.Get(contentUrl)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(respContent.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var courseResponse model.CourseResponse
	if err := json.Unmarshal(body, &courseResponse); err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}

	tr := model.Transaction{
		CourseId: courseResponse.Course.Id,
		UserId:   userResponse.Email.Id,
		Cost:     courseResponse.Course.Cost,
	}

	return p.repo.CreateNewTransaction(tr)
}
