package gmm

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createMySQLClient(port int, dbName string) (sqlDB *sql.DB, gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf("root@tcp(127.0.0.1:%d)/%s", port, dbName)
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open sql mysql client: %w", err)
	}

	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open gorm mysql client: %w", err)
	}
	return sqlDB, gormDB, nil
}
