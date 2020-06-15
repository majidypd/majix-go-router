package majix

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"time"
)

type Provider interface {
	Set(key string, content map[string]interface{}) error
	Get(key string) (map[string]interface{}, error)
	Delete(key string) error
}

type RedisProvider struct {
	Driver interface{}
}

func (p *RedisProvider) Set(key string, content map[string]interface{}) error {
	rdb := p.Driver.(*redis.Client)
	c, err := json.Marshal(content)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, key, string(c), MaxAge).Err()
	if err != nil {
		return err
	}
	return nil
}

func (p *RedisProvider) Get(key string) (map[string]interface{}, error) {
	rdb := p.Driver.(*redis.Client)
	content, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	json.Unmarshal([]byte(content), &result)
	return result, nil
}

func (p *RedisProvider) Delete(key string) error {
	rdb := p.Driver.(*redis.Client)
	rdb.Del(ctx, key).Err()
	return nil
}

type GormProvider struct {
	Driver interface{}
}

type SessionEntity struct {
	Value      string    `gorm:"column:value"`
	Content    string    `gorm:"column:content"`
	Expiration time.Time `gorm:"column:expiration"`
}

func (*SessionEntity) TableName() string {
	return "session"
}

func (p *GormProvider) Set(key string, c map[string]interface{}) error {
	rdb := p.Driver.(*gorm.DB)
	content, err := json.Marshal(c)
	if err != nil {
		return err
	}
	s := SessionEntity{
		Value:      key,
		Content:    string(content),
		Expiration: time.Now(),
	}
	rdb.FirstOrCreate(&s, SessionEntity{Value: key})
	return nil
}

func (p *GormProvider) Get(key string) (map[string]interface{}, error) {
	rdb := p.Driver.(*gorm.DB)
	s := SessionEntity{}
	err := rdb.Where("value = ?", key).First(&s)
	if err.Error != nil {
		return nil, err.Error
	}
	result := map[string]interface{}{}
	json.Unmarshal([]byte(s.Content), &result)
	return result, nil
}

func (p *GormProvider) Delete(key string) error {
	rdb := p.Driver.(*gorm.DB)
	rdb.Where("value = ?", key).Delete(SessionEntity{})
	return nil
}
