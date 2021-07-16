package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tai9/golang_jwt/dto"
	"github.com/tai9/golang_jwt/entity"
	"github.com/tai9/golang_jwt/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser

}
func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
