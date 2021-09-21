package managers

import (
	"github.com/biges/hybrone-sentinel-backend/models"
	"github.com/biges/hybrone-sentinel-backend/helpers"
	"gorm.io/gorm"
)

//UserStorage defines the behaviors required to perform CRUD operations on a Device model.
type UserStorage interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	Get(user *models.User) error
	List(user []*models.User)error
}

// UserManager represents Device related CRUD operations.
type UserManager struct {
	db *gorm.DB
}

// NewUserManager returns DeviceManager instance.
func NewUserManager(db *gorm.DB) UserStorage {
	return &UserManager{db: db}
}

// Create creates new User record.
func (m UserManager) Create(user *models.User) error {
	result := m.db.Create(user)
	return helpers.GORMErrConverter(result.Error)
}

// Update updates given User record.
func (m UserManager) Update(user *models.User) error {
	result := m.db.Save(user)
	return helpers.GORMErrConverter(result.Error)
}

// Delete deletes User record.
func (m UserManager) Delete(id uint) error {
	result := m.db.Delete(&models.User{Base: models.Base{ID: id}})
	return result.Error
}

// Get get User by given id.
func (m UserManager) Get(user *models.User) error {
	result := m.db.First(user)
	return helpers.GORMErrConverter(result.Error)
}

// List get Users list.
func (m UserManager) List(users *[]models.User) error {
	result := m.db.Find(users)
	return helpers.GORMErrConverter(result.Error)
}