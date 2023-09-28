package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Course interface {
	Add(name string, description string, folderName string) error
	Delete(id uint) (string, error)
	GetByName(name string) ([]core.Course, error)
	GetAll() ([]core.Course, error)
	Get(path string) (core.СourseСontent, error)
}

type Module interface {
	Add(name string, description string, courseName string, folderName string) error

	Delete(id uint) (string, error)
	Get(path string) (core.ModLes, error)
}

type Lesson interface {
	Add(name string, description string, fileName string, courseName string, moduleName string) error
	Delete(id uint) (string, error)
	Get(path string, mdfile []string) (core.LesMd, error)
}

type Repository struct {
	Course
	Module
	Lesson
	//db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Course: NewCoursePostgres(db),
		Module: NewModulePostgres(db),
		Lesson: NewLessonPostgres(db),
	}
}
