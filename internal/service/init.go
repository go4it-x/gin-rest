package service

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
	"home/internal/model"
	"time"
)

var (
	db   *gorm.DB
	Test = NewTestService()
)

//Init instantiate the service
func Init() error {
	// init orm
	//db = gorm.StdConfig("test").Build()
	//db.SingularTable(true)
	//migrate()

	return nil
}

// Migrate data migration
func migrate() {
	if !db.HasTable(&model.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
		// Add a piece of test data
		user := model.User{
			Model: model.Model{
				CreatedTime: time.Now().Unix(),
				UpdatedTime: time.Now().Unix(),
			},
			Username: "test",
			Password: "test Password",
		}

		db.Create(&user)
	}
}
