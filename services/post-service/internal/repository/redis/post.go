package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/1111000110/go_service/services/post-service/internal/model"
	"github.com/go-redis/redis/v8"
	"time"
)

func CacheDeletePostByPid(ctx context.Context, pid int64) error {
	rdb := GetDefaultCollection()
	key := fmt.Sprintf("post:%d", pid)
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
func CacheSetPost(ctx context.Context, post *model.Post) error {
	rdb := GetDefaultCollection()
	// 将 post 对象转换为 JSON 字符串
	postJSON, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("failed to marshal post: %v", err)
	}
	// 设置 Redis 键值对，键为 "post:<post.Id>"，值为 JSON 字符串
	key := fmt.Sprintf("post:%d", post.Pid)
	err = rdb.Set(ctx, key, postJSON, 90*time.Second).Err() // 设置 0 表示没有过期时间，直到手动删除
	if err != nil {
		return fmt.Errorf("failed to set post in Redis: %v", err)
	}
	return nil
}
func CacheSetPosts(ctx context.Context, posts *[]model.Post) error {
	for _, post := range *posts {
		err := CacheSetPost(ctx, &post)
		if err != nil {
			return err
		}
	}
	return nil
}
func CacheGetPostsByPids(ctx context.Context, pids []int64) (*[]model.Post, error) {
	// 用于存储获取到的帖子
	var posts []model.Post
	rdb := GetDefaultCollection()
	// 循环遍历所有的 pids
	for _, pid := range pids {
		// 生成 Redis 键
		key := fmt.Sprintf("post:%d", pid)
		// 从 Redis 中获取帖子数据
		val, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			// 如果没有找到数据
			continue //查看下一个
		} else if err != nil {
			// 处理其他错误
			return nil, fmt.Errorf("failed to get post from Redis: %v", err)
		}
		// 反序列化 JSON 数据为 Post 对象
		var post model.Post
		err = json.Unmarshal([]byte(val), &post)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal post data: %v", err)
		}
		// 将获取的帖子添加到结果切片中
		posts = append(posts, post)
	}
	// 返回获取到的帖子数据
	return &posts, nil
}
