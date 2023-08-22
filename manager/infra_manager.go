package manager

import (
	"database/sql"
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/config"
	_ "github.com/lib/pq"
)

// // akan otomatis terpanggil jika package dia di import
// func init() {}

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
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

	db, err := sql.Open(i.cfg.Driver, dsn)
	if err != nil {
		return err
	}

	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}

	if err := conn.initDb(); err != nil {
		return nil, err
	}
	return conn, nil
}
