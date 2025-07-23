package services

import (
	"gopay-clone/config"
	"gopay-clone/models"
)

type DriverService struct {
	db *config.Database
}

func NewDriverService(db *config.Database) *DriverService {
	return &DriverService{db: db}
}

func (s *DriverService) CreateDriverProfile(driver *models.DriverProfile) error {
	return s.db.Create(driver).Error
}

func (s *DriverService) GetDriverByID(id uint) (*models.DriverProfile, error) {
	var driver models.DriverProfile
	result := s.db.Preload("User").First(&driver, id)
	return &driver, result.Error
}

func (s *DriverService) GetDriverByUserID(userID uint) (*models.DriverProfile, error) {
	var driver models.DriverProfile
	result := s.db.Preload("User").Where("user_id = ?", userID).First(&driver)
	return &driver, result.Error
}

func (s *DriverService) GetAllDrivers() ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	result := s.db.Preload("User").Find(&drivers)
	return drivers, result.Error
}

func (s *DriverService) GetAvailableDrivers() ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	result := s.db.Preload("User").Where("status = ? AND is_verified = ?", "online", true).Find(&drivers)
	return drivers, result.Error
}

func (s *DriverService) GetDriversByVehicleType(vehicleType models.VehicleType) ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	result := s.db.Preload("User").Where("vehicle_type = ? AND status = ? AND is_verified = ?", vehicleType, "online", true).Find(&drivers)
	return drivers, result.Error
}

func (s *DriverService) UpdateDriver(id uint, updates map[string]any) error {
	return s.db.Model(&models.DriverProfile{}).Where("id = ?", id).Updates(updates).Error
}

func (s *DriverService) UpdateDriverStatus(userID uint, status string) error {
	return s.db.Model(&models.DriverProfile{}).Where("user_id = ?", userID).Update("status", status).Error
}

func (s *DriverService) UpdateDriverLocation(userID uint, location string) error {
	return s.db.Model(&models.DriverProfile{}).Where("user_id = ?", userID).Update("current_location", location).Error
}

func (s *DriverService) DeleteDriver(id uint) error {
	return s.db.Delete(&models.DriverProfile{}, id).Error
}

func (s *DriverService) VerifyDriver(id uint) error {
	return s.db.Model(&models.DriverProfile{}).Where("id = ?", id).Update("is_verified", true).Error
}
