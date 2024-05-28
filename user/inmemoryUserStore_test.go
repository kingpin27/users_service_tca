package user

import "testing"

func TestCreateUser(t *testing.T) {
	store := NewInMemoryUserStore()
	user := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	err := store.CreateUser(user)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	store := NewInMemoryUserStore()
	user := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	store.CreateUser(user)
	user, err := store.GetUserById(1)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	if user.Id != 1 {
		t.Errorf("Expected user id to be 1, got %d", user.Id)
	}
	if user.FName != "John" {
		t.Errorf("Expected user first name to be John, got %s", user.FName)
	}
	if user.City != "San Francisco" {
		t.Errorf("Expected user city to be San Francisco, got %s", user.City)
	}
	if user.Phone != 1234567890 {
		t.Errorf("Expected user phone to be 1234567890, got %d", user.Phone)
	}
	if user.Height != 72 {
		t.Errorf("Expected user height to be 72, got %d", user.Height)
	}
	if user.Married != MARITAL_STATUS_TRUE {
		t.Errorf("Expected user married to be true, got %d", user.Married)
	}
}

func TestGetUsersListByIds(t *testing.T) {
	store := NewInMemoryUserStore()
	user1 := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	user2 := User{
		Id:      2,
		FName:   "Jane",
		City:    "New York",
		Phone:   1234567890,
		Height:  68,
		Married: MARITAL_STATUS_FALSE,
	}
	store.CreateUser(user1)
	store.CreateUser(user2)
	users, err := store.GetUsersListByIds([]uint32{1, 2})
	if err != nil {
		t.Errorf("Error getting users: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
	if users[0].Id != 1 {
		t.Errorf("Expected user id to be 1, got %d", users[0].Id)
	}
	if users[1].Id != 2 {
		t.Errorf("Expected user id to be 2, got %d", users[1].Id)
	}
}

func TestDeleteUser(t *testing.T) {
	store := NewInMemoryUserStore()
	user := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	store.CreateUser(user)
	err := store.DeleteUser(1)
	if err != nil {
		t.Errorf("Error deleting user: %v", err)
	}
	_, err = store.GetUserById(1)
	if err == nil {
		t.Errorf("Expected error getting user, got nil")
	}
}

func TestUpdateUser(t *testing.T) {
	store := NewInMemoryUserStore()
	user := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	store.CreateUser(user)
	user.FName = "Jane"
	user.City = "New York"
	user.Phone = 9876543210
	user.Height = 68
	user.Married = MARITAL_STATUS_FALSE
	err := store.UpdateUser(user)
	if err != nil {
		t.Errorf("Error updating user: %v", err)
	}
	user, err = store.GetUserById(1)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	if user.FName != "Jane" {
		t.Errorf("Expected user first name to be Jane, got %s", user.FName)
	}
	if user.City != "New York" {
		t.Errorf("Expected user city to be New York, got %s", user.City)
	}
	if user.Phone != 9876543210 {
		t.Errorf("Expected user phone to be 9876543210, got %d", user.Phone)
	}
	if user.Height != 68 {
		t.Errorf("Expected user height to be 68, got %d", user.Height)
	}
	if user.Married != MARITAL_STATUS_FALSE {
		t.Errorf("Expected user married to be false, got %d", user.Married)
	}
}

func TestSearchUser(t *testing.T) {
	store := NewInMemoryUserStore()
	user1 := User{
		Id:      1,
		FName:   "John",
		City:    "San Francisco",
		Phone:   1234567890,
		Height:  72,
		Married: MARITAL_STATUS_TRUE,
	}
	user2 := User{
		Id:      2,
		FName:   "Jane",
		City:    "New York",
		Phone:   1234567890,
		Height:  68,
		Married: MARITAL_STATUS_FALSE,
	}
	store.CreateUser(user1)
	store.CreateUser(user2)
	query := User{
		FName: "John",
		City:  "San Francisco",
	}
	users, err := store.Search(query)
	if err != nil {
		t.Errorf("Error searching users: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
	if users[0].Id != 1 {
		t.Errorf("Expected user id to be 1, got %d", users[0].Id)
	}
	if users[0].FName != "John" {
		t.Errorf("Expected user first name to be John, got %s", users[0].FName)
	}
	if users[0].City != "San Francisco" {
		t.Errorf("Expected user city to be San Francisco, got %s", users[0].City)
	}
	if users[0].Phone != 1234567890 {
		t.Errorf("Expected user phone to be 1234567890, got %d", users[0].Phone)
	}
	if users[0].Height != 72 {
		t.Errorf("Expected user height to be 72, got %d", users[0].Height)
	}
	if users[0].Married != MARITAL_STATUS_TRUE {
		t.Errorf("Expected user married to be true, got %d", users[0].Married)
	}
}
