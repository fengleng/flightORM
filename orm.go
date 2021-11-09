package flightORM

import (
	"context"
	"github.com/fengleng/flight/sqlparser/sqlparser"
	"github.com/pingcap/errors"

	"github.com/fengleng/go-mysql-client/backend"
	"time"
)

type DB struct {
	*Config
	conn  *backend.DB
	Error error
	sqlparser.SQLNode
}

type Config struct {
	dbName   string
	Addr     string
	User     string
	PassWord string
}

func Open(cfg *Config, opts ...Option) (*DB, error) {
	if cfg == nil {
		cfg = new(Config)
	}
	for _, opt := range opts {
		err := opt.Apply(cfg)
		return nil, err
	}
	db := new(DB)
	db.Config = cfg
	openConn, err := backend.Open(cfg.Addr, cfg.User, cfg.PassWord, cfg.dbName)
	if err != nil {
		return nil, errors.Trace(err)
	}
	db.conn = openConn
	return db, nil
}

func (db *DB) Where() (tx *DB) {
	tx = &DB{Config: db.Config, Error: db.Error}
	return tx
}

func (c *Config) Apply(config *Config) error {
	config = c
	return nil
}

type Option interface {
	Apply(*Config) error
}

type Scope struct {
}

// Session  config when create session with Session() method
type Session struct {
	DryRun                   bool
	PrepareStmt              bool
	NewDB                    bool
	SkipHooks                bool
	SkipDefaultTransaction   bool
	DisableNestedTransaction bool
	AllowGlobalUpdate        bool
	FullSaveAssociations     bool
	QueryFields              bool
	Context                  context.Context
	//Logger                   logger.Interface
	NowFunc         func() time.Time
	CreateBatchSize int
}

func (db *DB) Session() {

}
