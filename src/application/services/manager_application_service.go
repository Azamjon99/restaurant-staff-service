package services

import (
	"context"
	"errors"

	"github.com/Azamjon99/restaurant-staff-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-staff-service/src/application/protos/restaurant_staff"
	managersvc "github.com/Azamjon99/restaurant-staff-service/src/domain/manager/services"
	"github.com/Azamjon99/restaurant-staff-service/src/infrastructure/jwt"
)

type ManagerAppService interface {
	RegisterManager(ctx context.Context, req *pb.RegisterManagerRequest) (*pb.RegisterManagerResponse, error)
	ChangeManagerPassword(ctx context.Context, req *pb.ChangeManagerPasswordRequest) (*pb.ChangeManagerPasswordResponse, error)
	LoginManager(ctx context.Context, req *pb.LoginManagerRequest) (*pb.LoginManagerResponse, error)
	GetManagerProfile(ctx context.Context, req *pb.GetManagerProfileRequest) (*pb.GetManagerProfileResponse, error)
	UpdateManagerProfile(ctx context.Context, req *pb.UpdateManagerProfileRequest) (*pb.UpdateManagerProfileResponse, error)
	CreateManagerRestaurantAssignment(ctx context.Context, req *pb.CreateManagerRestaurantAssignmentRequest) (*pb.CreateManagerRestaurantAssignmentResponse, error)
	RemoveManagerRestaurantAssignment(ctx context.Context, req *pb.RemoveManagerRestaurantAssignmentRequest) (*pb.RemoveManagerRestaurantAssignmentResponse, error)
	GetManagerRestaurant(ctx context.Context, req *pb.GetManagerRestaurantRequest) (*pb.GetManagerRestaurantResponse, error)
}

type managerAppSvcImpl struct {
	managerSvc managersvc.ManagerService
	tokenSvc   jwt.Service
}

func NewManagerAppService(
	managerSvc managersvc.ManagerService,
	tokenSvc jwt.Service,
) ManagerAppService {
	return &managerAppSvcImpl{
		managerSvc: managerSvc,
		tokenSvc:   tokenSvc,
	}
}
func (s *managerAppSvcImpl) RegisterManager(ctx context.Context, req *pb.RegisterManagerRequest) (*pb.RegisterManagerResponse, error) {

	if dtos.ValidateRegisterManagerRequestPB(req) != nil {
		return nil, dtos.ValidateRegisterManagerRequestPB(req)
	}
	managerID, err := s.managerSvc.RegisterManager(ctx, req.GetEntityId(), req.GetRestaurantId(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.RegisterManagerResponse{
		ManagerId: managerID,
	}, nil
}

func (s *managerAppSvcImpl) ChangeManagerPassword(ctx context.Context, req *pb.ChangeManagerPasswordRequest) (*pb.ChangeManagerPasswordResponse, error) {

	if req.GetManagerId() == "" {
		return nil, errors.New("Invalid or missing manager_id")
	}
	if req.GetCurrentPassword() == "" {
		return nil, errors.New("Invalid or missing current_password")
	}
	if req.GetNewPassword() == "" {
		return nil, errors.New("Invalid or missing new_password")
	}

	err := s.managerSvc.ChangeManagerPassword(ctx, req.GetManagerId(), req.GetCurrentPassword(), req.GetNewPassword())

	if err != nil {
		return nil, err
	}
	return &pb.ChangeManagerPasswordResponse{}, nil

}
func (s *managerAppSvcImpl) LoginManager(ctx context.Context, req *pb.LoginManagerRequest) (*pb.LoginManagerResponse, error) {

	if req.GetEmail() == "" {
		return nil, errors.New("Invalid or missing email")
	}
	if req.GetPassword() == "" {
		return nil, errors.New("Invalid or missing password")
	}

	profile, err := s.managerSvc.LoginManager(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, err
	}
	token, err := s.tokenSvc.CreateToken(ctx, profile.ManagerID)
	if err != nil {
		return nil, err
	}
	return &pb.LoginManagerResponse{
		Profile: dtos.ToManagerProfilePB(profile),
		Token:   token,
	}, nil

}
func (s *managerAppSvcImpl) GetManagerProfile(ctx context.Context, req *pb.GetManagerProfileRequest) (*pb.GetManagerProfileResponse, error) {

	if req.GetManagerId() == "" {
		return nil, errors.New("Invalid or missing manager_id")
	}

	manager_profile, err := s.managerSvc.GetManagerProfile(ctx, req.GetManagerId())

	if err != nil {
		return nil, err
	}
	return &pb.GetManagerProfileResponse{
		Profile: dtos.ToManagerProfilePB(manager_profile),
	}, nil
}
func (s *managerAppSvcImpl) UpdateManagerProfile(ctx context.Context, req *pb.UpdateManagerProfileRequest) (*pb.UpdateManagerProfileResponse, error) {

	if req.GetManagerId() == "" {
		return nil, errors.New("Invalid or missing manager_id")
	}
	if req.GetName() == "" {
		return nil, errors.New("Invalid or missing name")
	}

	profile, err := s.managerSvc.UpdateManagerProfile(ctx, req.GetManagerId(), req.GetName(), req.GetPhoneNumber(), req.GetEmail(), req.GetImageUrl())

	if err != nil {
		return nil, err
	}
	return &pb.UpdateManagerProfileResponse{
		Profile: dtos.ToManagerProfilePB(profile),
	}, nil
}
func (s *managerAppSvcImpl) CreateManagerRestaurantAssignment(ctx context.Context, req *pb.CreateManagerRestaurantAssignmentRequest) (*pb.CreateManagerRestaurantAssignmentResponse, error) {

	if req.GetManagerId() == "" {
		return nil, errors.New("Invalid or missing manager_id")
	}
	if req.GetRestaurantId() == "" {
		return nil, errors.New("Invalid or missing restaurant_id")
	}
	assignmentID, err := s.managerSvc.CreateManageRestaurantAssignment(ctx, req.GetManagerId(), req.GetRestaurantId())
	if err != nil {
		return nil, err
	}
	return &pb.CreateManagerRestaurantAssignmentResponse{
		AssigmentId: assignmentID,
	}, nil
}
func (s *managerAppSvcImpl) RemoveManagerRestaurantAssignment(ctx context.Context, req *pb.RemoveManagerRestaurantAssignmentRequest) (*pb.RemoveManagerRestaurantAssignmentResponse, error) {

	if req.GetAssigmentId() == 0 {
		return nil, errors.New("Invalid or missing assigment_id")
	}
	err := s.managerSvc.RemoveManagerRestaurantAssignment(ctx, int(req.GetAssigmentId()))

	if err != nil {
		return nil, err
	}
	return &pb.RemoveManagerRestaurantAssignmentResponse{}, nil
}
func (s *managerAppSvcImpl) GetManagerRestaurant(ctx context.Context, req *pb.GetManagerRestaurantRequest) (*pb.GetManagerRestaurantResponse, error) {

	if req.GetManagerId() == "" {
		return nil, errors.New("Invalid or missing manager_id")
	}

	resID, err := s.managerSvc.GetManagerRestaurant(ctx, req.GetManagerId())

	if err != nil {
		return nil, err
	}
	return &pb.GetManagerRestaurantResponse{
		RestaurantId: resID,
	}, nil
}
