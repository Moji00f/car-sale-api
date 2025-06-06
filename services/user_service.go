package services

import (
	"github.com/Moji00f/SimpleProject/api/dto"
	"github.com/Moji00f/SimpleProject/common"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:        cfg,
		database:   db.GetDb(),
		logger:     logger,
		otpService: NewOtpService(cfg),
	}
}

func (u *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := u.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	
	return nil
}
