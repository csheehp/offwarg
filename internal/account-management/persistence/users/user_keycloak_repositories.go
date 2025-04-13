package users

import (
	"errors"
	"strings"

	"github.com/neel4os/warg/internal/common/cache"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/keycloak"
	"github.com/neel4os/warg/internal/common/redis"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"
)

type UserConcreteRepository struct {
	cliet       *resty.Client
	redisClient *redis.RedisCon
}

func NewUserConcreteRepository() *UserConcreteRepository {
	cfg := config.GetConfig()
	cache := cache.NewIMCache(cfg)
	token := cache.GetToken()
	restyClient := resty.New()
	restyClient.SetHeader("Authorization", "Bearer "+token)
	restyClient.SetHeader("Content-Type", "application/json")
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetRetryCount(3)
	restyClient.SetBaseURL(cfg.IdpConfig.Url + "/admin/realms/" + cfg.IdpConfig.RealmName)
	rediscon := redis.GetRedisCon(cfg)
	return &UserConcreteRepository{
		cliet:       restyClient,
		redisClient: rediscon,
	}
}

func (r *UserConcreteRepository) CreateUser(email string, firstname string, lastname string) (string, error) {
	// Create a new user in Keycloak

	resp, err := r.cliet.R().SetBody(keycloak.NewUserRepresentation(email, firstname, lastname)).Post("/users")
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating user")
		return "", err
	}

	if resp.IsError() {
		log.Error().Str("response", string(resp.Bytes())).Caller().Msg("Error while creating user")
		return "", errors.New("Error while creating user")
	}

	location := resp.Header().Get("Location")
	if location == "" {
		return "", errors.New("Location header is missing in the response")
	}

	userId := location[strings.LastIndex(location, "/")+1:]
	return userId, nil
}
