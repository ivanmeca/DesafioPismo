package database

import (
	"Hemoce/application/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.Clinic{})
	db.AutoMigrate(model.Worker{})
	db.AutoMigrate(model.Procedure{})

	//db.AutoMigrate(model.Schedules{})
	db.AutoMigrate(model.ClinicWorkers{})
	db.AutoMigrate(model.Scale{})
	db.AutoMigrate(model.Step{})
	db.AutoMigrate(model.ProcedureStep{})
	db.AutoMigrate(model.Patient{})

	db.AutoMigrate(model.FormField{})
	db.AutoMigrate(model.FormGroup{})
	db.AutoMigrate(model.Form{})
	db.AutoMigrate(model.FieldsIntoGroup{})
}
