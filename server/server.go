package server

import (
	"context"
	"log"
	"net"

	"github.com/kingpin27/users_service_tca/model"
	"github.com/kingpin27/users_service_tca/pb"
	"google.golang.org/grpc"
)

var userStore *model.InMemoryUserStore

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := ValidateCreateUserRequestMessage(in); err != nil {
		return nil, err
	}
	user := model.User{
		FName:   in.Name,
		City:    in.City,
		Phone:   in.Phone,
		Height:  uint8(in.Height),
		Married: model.GetMarrigeStatusBoolToEnum(in.Married),
	}
	returnedUser, err := userStore.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id:      returnedUser.Id,
		Name:    user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  uint64(user.Height),
		Married: in.Married,
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	if err := ValidateGetUserByIdRequestMessage(in); err != nil {
		return nil, err
	}
	user, err := userStore.GetUserById(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIdResponse{
		Id:      user.Id,
		Name:    user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  uint64(user.Height),
		Married: model.GetMarriedStatusEnumToBool(user.Married),
	}, nil
}

func (s *UserService) GetUsersListByIds(ctx context.Context, in *pb.GetUsersListByIdsRequest) (*pb.GetUsersListByIdsResponse, error) {
	if err := ValidateGetUsersListByIdsRequestMessage(in); err != nil {
		return nil, err
	}
	users, err := userStore.GetUsersListByIds(in.Ids)
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      user.Id,
			Name:    user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  uint64(user.Height),
			Married: model.GetMarriedStatusEnumToBool(user.Married),
		})
	}
	return &pb.GetUsersListByIdsResponse{
		Users: pbUsers,
	}, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, in *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	if err := ValidateDeleteUserByIdRequestMessage(in); err != nil {
		return nil, err
	}
	err := userStore.DeleteUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserByIdResponse{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	updateUser := ParseOptionalStringField(InputFields{
		Name:    in.Name,
		Phone:   in.Phone,
		City:    in.City,
		Height:  in.Height,
		Married: in.Married,
	}, &in.Id)
	err := userStore.UpdateUser(*updateUser)
	if err != nil {
		return &pb.UpdateUserResponse{}, err
	}
	return &pb.UpdateUserResponse{}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, in *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	query := ParseOptionalStringField(InputFields{
		Name:    in.Name,
		Phone:   in.Phone,
		City:    in.City,
		Height:  in.Height,
		Married: in.Married,
	}, nil)
	users, err := userStore.Search(*query)
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      user.Id,
			Name:    user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  uint64(user.Height),
			Married: model.GetMarriedStatusEnumToBool(user.Married),
		})
	}
	return &pb.SearchUsersResponse{
		Users: pbUsers,
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userStore = model.NewInMemoryUserStore()

	serviceRegistrar := grpc.NewServer()
	userService := UserService{}

	pb.RegisterUserServiceServer(serviceRegistrar, &userService)

	if err := serviceRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
