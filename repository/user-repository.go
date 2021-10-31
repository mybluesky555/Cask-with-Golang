package repository

import (
	"fmt"
	"log"

	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string, isAdmin bool) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	AllUsers(info dto.AllDataDTO) ([]entity.User, int64)
	DeleteUser(id int) error
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
		db.connection.Save(&user)
	} else {
		var tempUser entity.User
		// db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
		db.connection.Omit("password").Save(&user)
	}
	return user
}

func (db *userConnection) VerifyCredential(email string, password string, isAdmin bool) interface{} {
	var user entity.User
	res := db.connection.Model(&entity.User{}).Where("email = ?", email).Where("is_admin = ?", isAdmin).Find(&entity.UserForClient{}).Take(&user)
	fmt.Println(res)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) AllUsers(info dto.AllDataDTO) ([]entity.User, int64) {
	var users []entity.User
	offset := (info.Page - 1) * info.PerPage
	search := "%" + info.Search + "%"
	search_db := db.connection.Where("name Like ? OR email LIKE ?", search, search)
	if info.Type == "all" {
		search_db.Limit(info.PerPage).Offset(offset).Find(&users)
	} else {
		var user_type int
		if info.Type == "admin" {
			user_type = 1
		} else {
			user_type = 0
		}
		search_db.Where("is_admin=?", user_type).Limit(info.PerPage).Offset(offset).Find(&users)
	}
	var total_count int64
	search_db.Count(&total_count)
	return users, total_count
}

func (db *userConnection) DeleteUser(id int) error {
	result := db.connection.Delete(&entity.User{}, id)
	return result.Error
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
