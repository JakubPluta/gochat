package user

import (
	"context"
	"server/internal/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const secretKey = "secret" // todo: generate and place in env

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repo Repository) Service {
	return &service{
		Repository: repo,
		timeout:    5 * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil

}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(ctx context.Context, request *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if err := utils.VerifyPassword(request.Password, u.Password); err != nil {
		return nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}
	return &LoginUserRes{accessToken: ss, ID: strconv.Itoa(int(u.ID)), Username: u.Username}, nil
}
