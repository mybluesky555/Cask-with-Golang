package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string, isAdmin bool) interface{}
	CreateUser(user dto.RegisterDTO) (entity.User, error)
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
	CheckAdminAndActive(userID string) int
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string, isAdmin bool) interface{} {
	res := service.userRepository.VerifyCredential(email, password, isAdmin)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword && v.Active {
			return res
		} else if !v.Active {
			return 1 //Not Active
		}
		return 2 // Wrong Password
	}
	return 3 // Wrong UserName
}

func (service *authService) CreateUser(user dto.RegisterDTO) (entity.User, error) {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err.Error())
		return userToCreate, err
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res, nil
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) CheckAdminAndActive(userID string) int {
	user := service.userRepository.ProfileUser(userID)
	if !user.Active {
		return 1 //Inactive
	} else {
		if user.IsAdmin {
			return 0 // Admin & Active
		} else {
			return 2 // User & Active
		}
	}
}
