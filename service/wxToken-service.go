package service

import (
	"context"
	"fox/dto"
	"fox/entity"
	"fox/repository"
	"github.com/go-redis/redis/v9"
	"github.com/mashingan/smapping"
	"log"
)

type WxTokenService interface {
	Insert(wxToken dto.WxTokenCreateDTO) entity.WxToken
	FindByAccessToken() entity.WxToken
	GetRedisAccessToken() (token string)
}
type tokenService struct {
	wxToken     repository.WxTokenRepository
	redisClient *redis.Client
}

func NewTokenService(token repository.WxTokenRepository, client *redis.Client) WxTokenService {
	return &tokenService{wxToken: token, redisClient: client}
}

func (s tokenService) Insert(token dto.WxTokenCreateDTO) entity.WxToken {
	tokenCreate := entity.WxToken{}
	err := smapping.FillStruct(&tokenCreate, smapping.MapFields(&token))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	var createToken = s.wxToken.InsertWxToken(tokenCreate)
	return createToken
}
func (s tokenService) FindByAccessToken() entity.WxToken {
	return s.wxToken.FindByAccessToken()
}
func (t *tokenService) GetRedisAccessToken() (token string) {
	ctx := context.Background()
	accessToken, err := t.redisClient.HGet(ctx, "wxAccessToken", "access_token").Result()
	if err != nil {
		log.Fatal(err)
	}
	return accessToken
}
