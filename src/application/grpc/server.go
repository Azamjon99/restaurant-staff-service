package grpc

import (
	"context"

	pb "github.com/Azamjon99/restaurant-staff-service/src/application/protos/restaurant_staff"
	"github.com/Azamjon99/restaurant-staff-service/src/application/services"
)

type Server struct {
	pb.StaffServiceServer
	managerApp services.ManagerAppService
}

func NewServer(managerApp services.ManagerAppService) *Server {
	return &Server{
		managerApp: managerApp,
	}
}

func (s *Server) RegisterManager(ctx context.Context, r *pb.RegisterManagerRequest) (*pb.RegisterManagerResponse, error) {

	return s.managerApp.RegisterManager(ctx, r)

}

func (s *Server) ChangeManagerPassword(ctx context.Context, r *pb.ChangeManagerPasswordRequest) (*pb.ChangeManagerPasswordResponse, error) {
	return s.managerApp.ChangeManagerPassword(ctx, r)
}

func (s *Server) LoginManager(ctx context.Context, r *pb.LoginManagerRequest) (*pb.LoginManagerResponse, error) {
	return s.managerApp.LoginManager(ctx, r)
}
func (s *Server) GetManagerProfile(ctx context.Context, r *pb.GetManagerProfileRequest) (*pb.GetManagerProfileResponse, error) {
	return s.managerApp.GetManagerProfile(ctx, r)
}
func (s *Server) UpdateManagerProfile(ctx context.Context, r *pb.UpdateManagerProfileRequest) (*pb.UpdateManagerProfileResponse, error) {
	return s.managerApp.UpdateManagerProfile(ctx, r)
}
func (s *Server) CreateManageRestaurantAssignment(ctx context.Context, r *pb.CreateManagerRestaurantAssignmentRequest) (*pb.CreateManagerRestaurantAssignmentResponse, error) {
	return s.managerApp.CreateManagerRestaurantAssignment(ctx, r)
}
func (s *Server) RemoveManagerRestaurantAssignment(ctx context.Context, r *pb.RemoveManagerRestaurantAssignmentRequest) (*pb.RemoveManagerRestaurantAssignmentResponse, error) {
	return s.managerApp.RemoveManagerRestaurantAssignment(ctx, r)
}
func (s *Server) GetManagerRestaurant(ctx context.Context, r *pb.GetManagerRestaurantRequest) (*pb.GetManagerRestaurantResponse, error) {
	return s.managerApp.GetManagerRestaurant(ctx, r)
}
