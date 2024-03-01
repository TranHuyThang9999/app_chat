package broker

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Cache là struct đại diện cho tầng cache
type Cache struct {
	client *redis.Client
}

// NewCache là hàm tạo mới một tầng cache với các tham số cấu hình
func NewCache() *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Địa chỉ và cổng Redis
		Password: "",               // Mật khẩu Redis (nếu cần)
		DB:       0,                // Số thứ tự cơ sở dữ liệu Redis
	})

	cache := &Cache{
		client: client,
	}

	return cache
}

// GetCache là phương thức để lấy dữ liệu từ cache dựa trên một key
func (c *Cache) GetCache(key string) (interface{}, error) {
	ctx := context.Background()
	value, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Không tìm thấy key trong cache
	} else if err != nil {
		return nil, err // Xảy ra lỗi khi truy vấn Redis
	}
	return value, nil
}

// SetCache là phương thức để đặt dữ liệu vào cache dựa trên một key và thời gian tồn tại
func (c *Cache) SetCache(key string, value interface{}) error {
	ctx := context.Background()
	err := c.client.Set(ctx, key, value, 0).Err()
	return err
}

// DeleteCache là phương thức để xóa dữ liệu khỏi cache dựa trên một key
func (c *Cache) DeleteCache(key string) error {
	ctx := context.Background()
	err := c.client.Del(ctx, key).Err()
	return err
}

// Close là phương thức để đóng cache và giải phóng kết nối Redis
func (c *Cache) Close() error {
	return c.client.Close()
}
