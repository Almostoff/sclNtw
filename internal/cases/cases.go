package cases

import (
	"fmt"
	"lesson30/internal/entity"
)

type (
	Usecase interface {
		CreateUser(*entity.User) (int, error)
		DeleteUser(int) (string, error)
		DeleteUserFromFL(int, int) error
		NewAge(int, int) error
		GetFriendList(int) (string, error)
		AddFriend(int, *entity.User) (string, error)
		SpecUser(int) (string, error)
		AllUsers() (string, error)
	}

	Repository interface {
		CreateUser(*entity.User) (int, error)
		DeleteUser(int) (string, error)
		DeleteUserFromFL(int, int) error
		NewAge(int, int) error
		GetFriendList(int) (*entity.User, error)
		AddFriend(int, *entity.User) (string, error)
		SpecUser(int) (string, error)
		AllUsers() (string, error)
	}
)

func (u *usecase) friendList(user *entity.User) string {
	return fmt.Sprintf("Friendlist`s %s: %s", user.Name, user.Friends)
}

type usecase struct {
	repository Repository
}

func (u *usecase) DeleteUserFromFL(id int, idt int) error {
	err := u.repository.DeleteUserFromFL(id, idt)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) NewAge(id int, age int) error {
	error := u.repository.NewAge(id, age)
	return error
}

func NewUseCase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) DeleteUser(id int) (string, error) {
	response, err := u.repository.DeleteUser(id)
	if err != nil {
		return "", err
	}
	return response, nil
}

func (u *usecase) AllUsers() (string, error) {
	response, error := u.repository.AllUsers()
	return response, error
}

func (u *usecase) SpecUser(id int) (string, error) {
	response, error := u.repository.SpecUser(id)
	return response, error
}

func (u *usecase) CreateUser(user *entity.User) (int, error) {
	uid, error := u.repository.CreateUser(user)
	return uid, error
}

func (u usecase) GetFriendList(id int) (string, error) {
	response, err := u.repository.GetFriendList(id)
	if err != nil {
		return "", err
	}
	return u.friendList(response), nil
}

func (u usecase) AddFriend(id int, v *entity.User) (string, error) {
	response, err := u.repository.AddFriend(id, v)
	if err != nil {
		return "", err
	}
	return response, nil
}
