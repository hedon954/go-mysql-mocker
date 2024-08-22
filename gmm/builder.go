package gmm

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"

	"github.com/dolthub/go-mysql-server/server"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GMMBuilder struct for building and managing the mock MySQL server
type GMMBuilder struct {
	dbName string
	port   int
	server *server.Server
	sqlDB  *sql.DB
	gormDB *gorm.DB
	err    error

	tables   []schema.Tabler
	models   []schema.Tabler
	sqlStmts []string
	sqlFiles []string
}

// Builder initializes a new GMMBuilder instance with db name,
// if db name is not provided, gmm would generate a random db name.
func Builder(db ...string) *GMMBuilder {
	b := &GMMBuilder{
		tables:   make([]schema.Tabler, 0),
		models:   make([]schema.Tabler, 0),
		sqlStmts: make([]string, 0),
		sqlFiles: make([]string, 0),
	}
	dbName := "gmm-test-db-" + uuid.NewString()[:6]
	if len(db) > 0 {
		dbName = db[0]
	}
	b.dbName = dbName
	return b
}

// Port sets the port for the MySQL server,
// if not set, gmm would generate a port start from 19527
func (b *GMMBuilder) Port(port int) *GMMBuilder {
	if b.err != nil {
		return b
	}
	b.port = port
	return b
}

// Build initializes and starts the MySQL server, returns handles to SQL and Gorm DB
func (b *GMMBuilder) Build() (sDB *sql.DB, gDB *gorm.DB, shutdown func(), err error) {
	if b.err != nil {
		return nil, nil, nil, b.err
	}

	// If not specify port, get an unused one form local machine.
	//
	// NOTE: The `getFreePort` method indeed has the limitation
	// that it cannot guarantee the port will remain available.
	// While it can find a currently unused port,
	// there is a possibility that the port might be occupied
	// by another process between the time the port number is
	// retrieved and the moment it is actually used.
	if b.port == 0 {
		var listener net.Listener
		listener, b.port, b.err = getFreePort()
		if b.err != nil {
			return nil, nil, nil, b.err
		}
		_ = listener.Close()
	}

	// Init mysql server
	b.initServer()
	if b.err != nil {
		return nil, nil, nil, b.err
	}

	// Start mysql server
	slog.Info("start go mysql mocker server, listening at 127.0.0.1:" + strconv.Itoa(b.port))
	go func() {
		if err := b.server.Start(); err != nil {
			panic(err)
		}
	}()

	shutdown = func() {
		_ = b.server.Close()
	}

	// Create client and connect to server
	b.sqlDB, b.gormDB, err = createMySQLClient(b.port, b.dbName)
	if err != nil {
		b.err = fmt.Errorf("failed to create sql client: %w", err)
		return nil, nil, nil, b.err
	}

	// Initialize tables and data
	b.initTables()
	b.initWithModels()
	b.initWithStmts()
	b.initWithFiles()
	if b.err != nil {
		return nil, nil, nil, b.err
	}

	return b.sqlDB, b.gormDB, shutdown, nil
}

// initServer initializes the mock MySQL server
func (b *GMMBuilder) initServer() *GMMBuilder {
	if b.err != nil {
		return b
	}
	b.server, b.err = createMySQLServer(b.dbName, b.port)
	return b
}

// CreateTable adds a table to be created upon initialization
func (b *GMMBuilder) CreateTable(table schema.Tabler) *GMMBuilder {
	if b.err != nil {
		return b
	}
	b.tables = append(b.tables, table)
	return b
}

// InitData adds initialization data to the mock database
func (b *GMMBuilder) InitData(data interface{}) *GMMBuilder {
	if b.err != nil {
		return b
	}

	if reflect.TypeOf(data).Kind() == reflect.Slice {
		slice := reflect.ValueOf(data)
		for i := 0; i < slice.Len(); i++ {
			item := slice.Index(i).Interface()
			model, ok := item.(schema.Tabler)
			if !ok {
				b.err = errors.New("every single data should implement gorm schema.Tabler")
				return b
			}
			b.models = append(b.models, model)
		}
		return b
	}

	model, ok := data.(schema.Tabler)
	if !ok {
		b.err = errors.New("data should implement gorm schema.Tabler")
		return b
	}
	b.models = append(b.models, model)
	return b
}

// SQLStmts adds SQL statements to be executed upon initialization
func (b *GMMBuilder) SQLStmts(stmts ...string) *GMMBuilder {
	if b.err != nil {
		return b
	}
	b.sqlStmts = append(b.sqlStmts, stmts...)
	return b
}

// SQLFiles adds SQL files whose contents are to be executed upon initialization
func (b *GMMBuilder) SQLFiles(files ...string) *GMMBuilder {
	if b.err != nil {
		return b
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			b.err = fmt.Errorf("sql file %s not exist", file)
			return b
		}
	}

	b.sqlFiles = append(b.sqlFiles, files...)
	return b
}

func (b *GMMBuilder) initTables() {
	if b.err != nil || len(b.tables) == 0 {
		return
	}

	slog.Info("start to init tables, count = " + strconv.Itoa(len(b.tables)))
	for _, table := range b.tables {
		if err := b.gormDB.AutoMigrate(table); err != nil {
			b.err = fmt.Errorf("failed to auto migrate(type=%T): %w", table, err)
			return
		}
	}
	slog.Info("init tables successfully, count = " + strconv.Itoa(len(b.tables)))
}

func (b *GMMBuilder) initWithModels() {
	if b.err != nil || len(b.models) == 0 {
		return
	}

	slog.Info("start to init data with models, count = " + strconv.Itoa(len(b.models)))
	for _, model := range b.models {
		if err := b.gormDB.AutoMigrate(model); err != nil {
			b.err = fmt.Errorf("failed to auto migrate(type=%T): %w", model, err)
			return
		}
		if err := b.gormDB.Create(model).Error; err != nil {
			b.err = fmt.Errorf("failed to init data(type=%T): %w", model, err)
			return
		}
	}
	slog.Info("init data with models successfully, count = " + strconv.Itoa(len(b.models)))
}

func (b *GMMBuilder) initWithStmts() {
	if b.err != nil || len(b.sqlStmts) == 0 {
		return
	}
	slog.Info("start to init data with sql stmts, count = " + strconv.Itoa(len(b.sqlStmts)))
	for _, stmt := range b.sqlStmts {
		stmts, err := splitSQLStatements(stmt)
		if err != nil {
			b.err = err
			return
		}
		if err = b.executeSQLStatements(stmts); err != nil {
			b.err = err
			return
		}
	}
	slog.Info("init data with sql stmts successfully, count = " + strconv.Itoa(len(b.sqlStmts)))
}

func (b *GMMBuilder) initWithFiles() {
	if b.err != nil || len(b.sqlFiles) == 0 {
		return
	}
	slog.Info("start to init data with sql files, count = " + strconv.Itoa(len(b.sqlFiles)))
	for _, file := range b.sqlFiles {
		stmts, err := splitSQLFile(file)
		if err != nil {
			b.err = fmt.Errorf("failed to split sql file '%s': %w", file, err)
			return
		}
		if err = b.executeSQLStatements(stmts); err != nil {
			b.err = err
			return
		}
	}
	slog.Info("init data with sql files successfully, count = " + strconv.Itoa(len(b.sqlFiles)))
}

func (b *GMMBuilder) executeSQLStatements(stmts []string) error {
	for _, stmt := range stmts {
		_, err := b.sqlDB.Exec(stmt)
		if err != nil {
			return fmt.Errorf("failed to exec sql stmt '%s': %w", stmt, err)
		}
	}
	return nil
}
