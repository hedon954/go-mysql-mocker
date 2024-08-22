package gmm

import (
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
)

func createMySQLServer(dbName string, port int) (*server.Server, error) {
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	pro := memory.NewDBProvider(db)
	engine := sqle.NewDefault(pro)

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("127.0.0.1:%d", port),
	}
	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return s, nil
}
