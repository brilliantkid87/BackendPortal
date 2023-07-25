package repositories

import (
	"test/models"

	"gorm.io/gorm"
)

// interface => kontrak
type UserRepository interface {
	FindUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(ID int, user models.User) (models.User, error)
	DeleteUser(ID int, user models.User) (models.User, error)
}

// function connection
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// method query
func (r *repository) FindUsers() ([]models.User, error) {
	var Users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&Users).Error

	return Users, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name, address) VALUES (?, ?)", user.Name, user.Address).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var User models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&User).Error

	return User, err
}

func (r *repository) UpdateUser(ID int, user models.User) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?,address=? WHERE id=?", user.Name, user.Address, ID).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteUser(ID int, user models.User) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
