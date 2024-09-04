package gmm

import (
	"context"
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/vitess/go/mysql"
)

func createMySQLServer(dbName string, port int) (*server.Server, error) {
	// create a new database
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	// create a new server engine
	pro := memory.NewDBProvider(db)
	engine := sqle.NewDefault(pro)
	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("127.0.0.1:%d", port),
	}

	// create a new server
	s, err := server.NewServer(config, engine, func(ctx context.Context, c *mysql.Conn, addr string) (sql.Session, error) {
		session := memory.NewSession(sql.NewBaseSession(), pro)
		return session, nil
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return s, nil
}
