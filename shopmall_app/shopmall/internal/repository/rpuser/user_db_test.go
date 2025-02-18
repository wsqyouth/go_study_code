package rpuser

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewUserDBRepoImpl(t *testing.T) {
	// 准备测试数据
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	tests := []struct {
		name string
		db   *gorm.DB
		want Repository
	}{
		{
			name: "successful creation",
			db:   gormDB,
			want: &UserDBRepoImpl{db: gormDB},
		},
		{
			name: "with nil db",
			db:   nil,
			want: &UserDBRepoImpl{db: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserDBRepoImpl(tt.db)
			assert.IsType(t, tt.want, got)

			// 检查内部 db 字段
			repo, ok := got.(*UserDBRepoImpl)
			assert.True(t, ok)
			assert.Equal(t, tt.db, repo.db)
		})
	}
}

func TestUserDBRepoImpl_CRUD(t *testing.T) {
	// 准备测试数据
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	repo := NewUserDBRepoImpl(gormDB)
	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		user := &User{
			Username: "test",
			Email:    "test@example.com",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users`").
			WithArgs(user.Username, user.Email).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Create(ctx, user)
		assert.NoError(t, err)
	})

	t.Run("GetByID", func(t *testing.T) {
		mock.ExpectQuery("SELECT \\* FROM `users`").
			WithArgs(uint64(1)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email"}).
				AddRow(1, "test", "test@example.com"))

		user, err := repo.GetByID(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test", user.Username)
	})

	t.Run("GetByID_NotFound", func(t *testing.T) {
		mock.ExpectQuery("SELECT \\* FROM `users`").
			WithArgs(uint64(999)).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := repo.GetByID(ctx, 999)
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("Update", func(t *testing.T) {
		user := &User{
			ID:       1,
			Username: "updated",
			Email:    "updated@example.com",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").
			WithArgs(user.Username, user.Email, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Update(ctx, user)
		assert.NoError(t, err)
	})

	t.Run("List", func(t *testing.T) {
		mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

		mock.ExpectQuery("SELECT \\* FROM `users`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email"}).
				AddRow(1, "user1", "user1@example.com").
				AddRow(2, "user2", "user2@example.com"))

		total, users, err := repo.List(ctx, 0, 10)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, users, 2)
	})

	t.Run("Delete", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `users`").
			WithArgs(uint64(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(ctx, 1)
		assert.NoError(t, err)
	})

	// 确保所有期望的 SQL 都被执行
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
