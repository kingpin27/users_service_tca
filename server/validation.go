package server

import (
	"errors"
	"fmt"

	"github.com/kingpin27/users_service_tca/pb"
)

func ValidateCreateUserRequestMessage(req *pb.CreateUserRequest) error {
	if req.Name == "" {
		return errors.New("empty name")
	}
	if req.City == "" {
		return errors.New("empty city")
	}
	phoneString := fmt.Sprint(req.Phone)
	if len(phoneString) != 10 {
		return errors.New("invalid phone number")
	}
	if req.Height < 24 || req.Height > 96 {
		return errors.New("invalid height")
	}
	return nil
}

func ValidateGetUserByIdRequestMessage(req *pb.GetUserByIdRequest) error {
	if req.Id == 0 {
		return errors.New("id cant be 0")
	}
	return nil
}

func ValidateGetUsersListByIdsRequestMessage(req *pb.GetUsersListByIdsRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("empty ids")
	}
	for _, id := range req.Ids {
		if id == 0 {
			return errors.New("id cant be 0")
		}
	}
	return nil
}

func ValidateDeleteUserByIdRequestMessage(req *pb.DeleteUserByIdRequest) error {
	if req.Id == 0 {
		return errors.New("id cant be 0")
	}
	return nil
}
