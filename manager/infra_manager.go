package manager

import (
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Employee{},
	); err != nil {
		return err
	}

	i.db = db
	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}

	if err := conn.initDb(); err != nil {
		return nil, err
	}
	return conn, nil
}
