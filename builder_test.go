package gmm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CertificationInfo struct {
	ID       int    `gorm:"column:id;index;primaryKey;autoIncrement"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (receiver CertificationInfo) TableName() string {
	return "certification_info"
}

type UserState struct {
	UID   string `gorm:"primaryKey;column:uid"`
	State string `gorm:"column:state"`
}

func (receiver UserState) TableName() string {
	return "user_state"
}

type UnknownType struct{}

func Test_GMMBuilder_Port(t *testing.T) {
	t.Run("listening on same port should return err", func(t *testing.T) {
		_, _, _, err := Builder().Port(9527).Build()
		assert.Nil(t, err)
		_, _, _, err = Builder().Port(9527).Build()
		assert.NotNil(t, err)
		assert.Equal(t, "failed to create server: Port 127.0.0.1:9527 already in use.", err.Error())
	})
}

func Test_GMMBuilder_SQLFiles(t *testing.T) {
	t.Run("non existing file should return err", func(t *testing.T) {
		_, _, _, err := Builder().SQLFiles("fixtures/not-exist.sql").Build()
		assert.NotNil(t, err)
		assert.Equal(t, "sql file fixtures/not-exist.sql not exist", err.Error())
	})

	t.Run("existing and valid file should ok", func(t *testing.T) {
		_, _, _, err := Builder().SQLFiles("fixtures/sequel_ace.sql").Build()
		assert.Nil(t, err)
	})
}

//nolint:lll
func Test_GMMBuilder_SQLStmts(t *testing.T) {
	t.Run("invalid sql statement should return err", func(t *testing.T) {
		_, _, _, err := Builder().SQLStmts("invalid sql statement").Build()
		assert.NotNil(t, err)
		assert.Equal(t, "failed to exec sql stmt 'invalid sql statement': Error 1105 (HY000): syntax error at position 8 near 'invalid'", err.Error())
	})

	t.Run("valid sql statement should ok", func(t *testing.T) {
		_, _, _, err := Builder().CreateTable(CertificationInfo{}).SQLStmts("insert into certification_info(id, username, password) values (1, 'hedon', 'hedon-pwd');").Build()
		assert.Nil(t, err)
	})

	t.Run("multi valid sql statements should ok", func(t *testing.T) {
		sql := `
CREATE TABLE IF NOT EXISTS users (
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(50) NOT NULL,
	email VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO users (username, email) VALUES
('john_doe', 'john.doe@example.com'),
('jane_smith', 'jane.smith@example.com'),
('alice_jones', 'alice.jones@example.com');`
		_, _, _, err := Builder().SQLStmts(sql).Build()
		assert.Nil(t, err)
	})
}

func Test_GMMBuilder_InitData(t *testing.T) {
	t.Run("invalid data type should return err", func(t *testing.T) {
		_, _, _, err := Builder().InitData(1).Build()
		assert.NotNil(t, err)
		assert.Equal(t, "data should implement gorm schema.Tabler", err.Error())
	})

	t.Run("auto increment, but struct is unaddressable, should panic", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _, _, _ = Builder().InitData(CertificationInfo{}).Build()
		})
	})

	t.Run("struct is unaddressable, but would not be changed, should ok", func(t *testing.T) {
		_, _, _, err := Builder().InitData(UserState{}).Build()
		assert.Nil(t, err)
	})

	t.Run("type not implemented schema.Tabler, should return error", func(t *testing.T) {
		_, _, _, err := Builder().InitData(&UnknownType{}).Build()
		assert.NotNil(t, err)
		assert.Equal(t, "data should implement gorm schema.Tabler", err.Error())
	})

	t.Run("pointer implemented schema.Tabler should ok", func(t *testing.T) {
		_, _, _, err := Builder().InitData(&CertificationInfo{}).Build()
		assert.Nil(t, err)
	})

	t.Run("slice implemented schema.Tabler should ok", func(t *testing.T) {
		_, _, _, err := Builder().InitData([]CertificationInfo{}).Build()
		assert.Nil(t, err)
	})

	t.Run("slice of different types that implemented schema.Tabler should ok", func(t *testing.T) {
		_, _, _, err := Builder().InitData([]interface{}{&CertificationInfo{}, &UserState{}}).Build()
		assert.Nil(t, err)
	})
}

func TestGMMBuilder_CreateTable(t *testing.T) {
	t.Run("valid type should ok", func(t *testing.T) {
		_, _, _, err := Builder().CreateTable(CertificationInfo{}).Build()
		assert.Nil(t, err)
	})

	t.Run("create multi times should ok", func(t *testing.T) {
		_, _, _, err := Builder().CreateTable(CertificationInfo{}).CreateTable(CertificationInfo{}).Build()
		assert.Nil(t, err)
	})
}
