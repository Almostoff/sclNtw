package repository

import (
	"errors"
	"fmt"
	"lesson30/internal/entity"
)

type repository struct {
	users map[string]*entity.User
}

func (r repository) toString(user *entity.User) string {
	return fmt.Sprintf("Name is %s and age is %d and his ID is %d \n", user.Name, user.Age, user.Id)
}

func (r repository) NewAge(id int, age int) error {
	for _, u := range r.users {
		if u.Id == id {
			u.Age = age
			return nil
		}
	}
	return errors.New("user not found")
}

func (r repository) DeleteUser(id int) (string, error) {
	for _, u := range r.users {
		if u.Id == id && id <= len(r.users) {
			delete(r.users, u.Name)
			return u.Name, nil
		}
	}
	return "", errors.New("user not found")
}

func (r repository) DeleteUserFromFL(id int, idt int) error {
	for _, u := range r.users {
		if u.Id == id && id <= len(r.users) {
			for i, _ := range u.Friends {
				if i+1 == idt {
					u.Friends = append(u.Friends[:i], (u.Friends)[i+1:]...)
					return nil
				}
			}
		}
	}
	return errors.New("user not found")
}

func (r repository) AllUsers() (string, error) {
	response := ""
	for _, user := range r.users {
		response += r.toString(user)
	}
	return response, nil
}

func (r repository) SpecUser(id int) (string, error) {
	for _, u := range r.users {
		if u.Id == id && id <= len(r.users) {
			response := r.toString(u)
			return response, nil
		}
	}
	return "", errors.New("user not found")
}

func NewRepository() *repository {
	return &repository{
		users: make(map[string]*entity.User),
	}
}

func (r repository) CreateUser(user *entity.User) (int, error) {
	r.users[user.Name] = user
	user.Id = len(r.users)
	return user.Id, nil
}

func (r repository) GetFriendList(id int) (*entity.User, error) {
	for _, u := range r.users {
		if u.Id == id && id <= len(r.users) {
			res := u
			return res, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r repository) AddFriend(id int, v *entity.User) (string, error) {
	for _, u := range r.users {
		if u.Id == id && id <= len(r.users) {
			for _, i := range r.users {
				if i.Id == v.Id {
					u.Friends = append(u.Friends, i.Name)
					i.Friends = append(i.Friends, u.Name)
					return u.Name, nil
				}
			}
		}
	}
	return "", errors.New("user not found")
}
