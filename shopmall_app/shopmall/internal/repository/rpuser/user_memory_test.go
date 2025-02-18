package rpuser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserMemRepoImpl(t *testing.T) {
	repo := NewUserMemRepoImpl()
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.users)
	assert.Equal(t, uint64(0), repo.idSeq)
}

func TestUserMemRepoImpl_CRUD(t *testing.T) {
	repo := NewUserMemRepoImpl()
	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		user := &User{
			Username: "test",
			Email:    "test@example.com",
		}

		err := repo.Create(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), user.ID)
		assert.Len(t, repo.users, 1)
	})

	t.Run("GetByID", func(t *testing.T) {
		// 获取存在的用户
		user, err := repo.GetByID(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test", user.Username)

		// 获取不存在的用户
		user, err = repo.GetByID(ctx, 999)
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("Update", func(t *testing.T) {
		// 更新存在的用户
		user := &User{
			ID:       1,
			Username: "updated",
			Email:    "updated@example.com",
		}
		err := repo.Update(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, "updated", repo.users[1].Username)

		// 更新不存在的用户
		nonExistUser := &User{
			ID:       999,
			Username: "nonexist",
		}
		err = repo.Update(ctx, nonExistUser)
		assert.Error(t, err)
	})

	t.Run("List", func(t *testing.T) {
		// 创建更多测试数据
		repo.Create(ctx, &User{Username: "user2", Email: "user2@example.com"})
		repo.Create(ctx, &User{Username: "user3", Email: "user3@example.com"})

		// 测试正常分页
		total, users, err := repo.List(ctx, 0, 2)
		assert.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, users, 2)

		// 测试超出范围的偏移量
		total, users, err = repo.List(ctx, 10, 2)
		assert.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, users, 0)
	})

	t.Run("Delete", func(t *testing.T) {
		// 删除存在的用户
		err := repo.Delete(ctx, 1)
		assert.NoError(t, err)
		assert.Len(t, repo.users, 2)

		// 删除不存在的用户
		err = repo.Delete(ctx, 999)
		assert.Error(t, err)
	})

	t.Run("Concurrent Operations", func(t *testing.T) {
		// 测试并发操作
		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func() {
				user := &User{
					Username: "concurrent",
					Email:    "concurrent@example.com",
				}
				repo.Create(ctx, user)
				done <- true
			}()
		}

		// 等待所有并发操作完成
		for i := 0; i < 10; i++ {
			<-done
		}

		total, _, err := repo.List(ctx, 0, 100)
		assert.NoError(t, err)
		assert.Equal(t, int64(12), total) // 2 (remaining from previous tests) + 10 new
	})
}
