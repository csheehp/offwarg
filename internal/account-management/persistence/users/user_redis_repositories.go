package users

import (
	"context"
	"time"

	"github.com/neel4os/warg/internal/common/util"
)

func (r *UserConcreteRepository) CreateUserVerficationCache(userId string) (string, error) {
	key := "user:" + userId + ":verification_cache"
	value := util.GenerataeNCharacterRandomString(64)
	// FIXME: the expiration shall be configurable
	_, err := r.redisClient.GetClient().Set(context.Background(), key, value, time.Duration(time.Minute*1440)).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
