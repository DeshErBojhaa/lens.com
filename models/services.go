package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServicesConfig func(*Services) error

func WithGorm(conInfo string) ServicesConfig {
	return func(s *Services) error {
		db, err := gorm.Open(postgres.Open(conInfo), &gorm.Config{})
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithUser(pepper, hmacKey string) ServicesConfig {
	return func(s *Services) error {
		s.User = NewUserService(s.db, pepper, hmacKey)
		return nil
	}
}

func WithGallery() ServicesConfig {
	return func(s *Services) error {
		s.Gallery = NewGalleryService(s.db)
		return nil
	}
}

func WithImage() ServicesConfig {
	return func(s *Services) error {
		s.Image = NewImageService()
		return nil
	}
}

func WithLog(prefix string) ServicesConfig {
	return func(s *Services) error {
		s.Log = log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
		return nil
	}
}

func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

type Services struct {
	Gallery GalleryService
	User    UserService
	Image   ImageService
	db      *gorm.DB
	Log     *log.Logger
}

func (s *Services) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}

func (s *Services) DestructiveReset() error {
	err := s.db.Migrator().DropTable(&User{}, &Gallery{})
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

func (s *Services) AutoMigrate() error {
	return s.db.Migrator().AutoMigrate(&User{}, &Gallery{})
}
