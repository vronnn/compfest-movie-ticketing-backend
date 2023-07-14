package repository

import (
	"context"
	"time"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entities.User) (entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) (error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (error) 
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{
		connection: db,
	}
}

func (ur *userRepository) RegisterUser(ctx context.Context, user entities.User) (entities.User, error){
	layout := "01/02/2006"
	birthDate, err := time.Parse(layout, user.TanggalLahir)
	if err != nil {
		return entities.User{}, err
	}

	today, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		return entities.User{}, err
	}

	user.Age = Age(birthDate, today)
	
	if err := ur.connection.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetAllUser(ctx context.Context) ([]entities.User, error){
	var user []entities.User
	if err := ur.connection.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error){
	var user entities.User
	if err := ur.connection.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	if err := ur.connection.Where("email = ?", email).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user entities.User) (error) {
	if err := ur.connection.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, userID uuid.UUID) (error) {
	if err := ur.connection.Delete(&entities.User{}, &userID).Error; err != nil {
		return err
	}
	return nil
}

func Age(birthdata, today time.Time) int {
	today = today.In(birthdata.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)

	by, bm, bd := birthdata.Date()
	birthdata = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)

	if today.Before(birthdata) {
		return 0
	}

	age := ty - by
	anniversary := birthdata.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}

	return age
}