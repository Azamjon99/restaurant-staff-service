package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/models"
)

type ManagerRepository interface {
	WithTx(ctx context.Context, f func(r ManagerRepository) error) error
	SaveManager(ctx context.Context, manager *models.Manager) error
	UpdateManager(ctx context.Context, manager *models.Manager) error
	DeleteManager(ctx context.Context, managerID string) error
	GetManager(ctx context.Context, managerID string) (*models.Manager, error)
	GetManagerByEmail(ctx context.Context, email string) (*models.Manager, error)
	SaveManagerProfile(ctx context.Context, profile *models.ManagerProfile) error
	UpdateManagerProfile(ctx context.Context, profile *models.ManagerProfile) error
	GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error)
	SaveManagerRestaurantAssignment(ctx context.Context, assignment *models.ManagerRestaurantAssignment) error
	DeleteManagerRestaurantAssignment(ctx context.Context, assignmentID int) error
	GetManagerRestaurant(ctx context.Context, managerID string) (string, error)
}
