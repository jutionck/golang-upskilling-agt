package manager

import (
	"database/sql"
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/config/db"
	_ "github.com/lib/pq"
)

// // akan otomatis terpanggil jika package dia di import
// func init() {}

type InfraManager interface {
	Conn() *db.Queries
}

type infraManager struct {
	db  *db.Queries
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	// buat url koneksi ke db postgres
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host,
		i.cfg.Port,
		i.cfg.User,
		i.cfg.Password,
		i.cfg.Name,
	)

	dbConn, err := sql.Open(i.cfg.Driver, dsn)
	if err != nil {
		return err
	}

	i.db = db.New(dbConn)
	return nil
}

func (i *infraManager) Conn() *db.Queries {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}

	if err := conn.initDb(); err != nil {
		return nil, err
	}
	return conn, nil
}
