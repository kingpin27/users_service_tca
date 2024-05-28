package model

import (
	"errors"
	"log"
	"sync"
)

type InMemoryUserStore struct {
	Users map[uint32]User
	sync.RWMutex
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		Users: make(map[uint32]User),
	}
}

func (store *InMemoryUserStore) GetUserById(id uint32) (User, error) {
	store.RLock()
	defer store.RUnlock()

	user, ok := store.Users[id]
	if !ok {
		return User{}, errors.New("User not found")
	}
	return user, nil
}

func (store *InMemoryUserStore) GetUsersListByIds(ids []uint32) ([]User, error) {
	store.RLock()
	defer store.RUnlock()

	var users []User
	for _, id := range ids {
		user, ok := store.Users[id]
		if !ok {
			return nil, errors.New("User not found")
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *InMemoryUserStore) DeleteUser(id uint32) error {
	store.Lock()
	defer store.Unlock()
	_, ok := store.Users[id]
	if !ok {
		return errors.New("User not found")
	}
	delete(store.Users, id)
	return nil
}

func (store *InMemoryUserStore) UpdateUser(user User) error {
	store.Lock()
	defer store.Unlock()
	_, ok := store.Users[user.Id]
	if !ok {
		return errors.New("User not found")
	}
	if user.FName == "" {
		user.FName = store.Users[user.Id].FName
	}
	if user.City == "" {
		user.City = store.Users[user.Id].City
	}
	if user.Phone == 0 {
		user.Phone = store.Users[user.Id].Phone
	}
	if user.Height == 0 {
		user.Height = store.Users[user.Id].Height
	}
	if user.Married == MARITAL_STATUS_NOT_SET {
		user.Married = store.Users[user.Id].Married
	}
	store.Users[user.Id] = user
	return nil
}

func (store *InMemoryUserStore) CreateUser(user User) (*User, error) {
	store.Lock()
	defer store.Unlock()
	if user.Id != 0 {
		_, ok := store.Users[user.Id]
		if ok {
			return nil, errors.New("User with this Id already exists")
		}
	} else {
		user.Id = uint32(len(store.Users) + 1)
	}
	store.Users[user.Id] = user
	return &user, nil
}

func (store *InMemoryUserStore) Search(queryUser User) ([]User, error) {
	log.Println((queryUser))
	store.RLock()
	defer store.RUnlock()
	var users []User
	for _, user := range store.Users {
		log.Println(user)
		if queryUser.Id != 0 && queryUser.Id != user.Id {
			continue
		}
		if queryUser.FName != "" && queryUser.FName != user.FName {
			continue
		}
		if queryUser.City != "" && queryUser.City != user.City {
			continue
		}
		if queryUser.Phone != 0 && queryUser.Phone != user.Phone {
			continue
		}
		if queryUser.Height != 0 && queryUser.Height != user.Height {
			continue
		}
		if queryUser.Married != MARITAL_STATUS_NOT_SET && queryUser.Married != user.Married {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
