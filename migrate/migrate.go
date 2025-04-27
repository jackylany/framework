package migrate

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Migrate interface {
	Up()
	Down()
}

type GormMigrate struct {
	Migrations []*gormigrate.Migration
	Db         *gorm.DB
}

func (m GormMigrate) Up() {
	gm := gormigrate.New(m.Db, gormigrate.DefaultOptions, m.Migrations)
	if err := gm.Migrate(); err != nil {
		log.Fatalf("Run up migrations failed, err: %s", err)
	}
	log.Println("Run up migrations success")
}

func (m GormMigrate) Down() {
	gm := gormigrate.New(m.Db, gormigrate.DefaultOptions, m.Migrations)
	ids := make([]string, 0)
	for _, v := range m.Migrations {
		if err := gm.RollbackMigration(v); err != nil {
			log.Fatalf("Run down [%s] migrations failed, err: %s", v.ID, err)
		}
		ids = append(ids, v.ID)
	}
	log.Printf("Run down [%s] migrations success", strings.Join(ids, ","))
}
