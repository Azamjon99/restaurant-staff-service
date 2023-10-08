package dtos

import (
	"errors"
	"time"

	pb "github.com/Azamjon99/restaurant-staff-service/src/application/protos/restaurant_staff"
	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/models"
)

func ToManagerProfilePB(profile *models.ManagerProfile) *pb.ManagerProfile {
	return &pb.ManagerProfile{
		ManagerId:   profile.ManagerID,
		EntityId:    profile.EntityID,
		Name:        profile.Name,
		ImageUrl:    profile.ImageUrl,
		PhoneNumber: profile.PhoneNumber,
		CreatedAt:   profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   profile.UpdatedAt.Format(time.RFC3339),
	}
}
func ValidateRegisterManagerRequestPB(req *pb.RegisterManagerRequest) error {

	if req.GetEntityId() == "" {
		return errors.New("Invalid or missing entity_id")
	}
	if req.GetRestaurantId() == "" {
		return errors.New("Invalid or missing restaurant_id")
	}
	if req.GetPassword() == "" {
		return errors.New("Invalid or missing password")
	}
	if req.GetEmail() == "" {
		return errors.New("Invalid or missing email")
	}
	return nil
}
