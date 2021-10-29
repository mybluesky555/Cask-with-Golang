package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.AdminDTO) entity.User
	Profile(userID string) entity.User
	AllUsers(info dto.AllDataDTO) ([]entity.User, int64)
	DeleteUser(id int) error
}
type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.AdminDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	return service.userRepository.UpdateUser(userToUpdate)
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) AllUsers(info dto.AllDataDTO) ([]entity.User, int64) {
	return service.userRepository.AllUsers(info)
}

func (service *userService) DeleteUser(id int) error {
	return service.userRepository.DeleteUser(id)
}
