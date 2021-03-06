package handler

import (
	"context"

	log "github.com/sirupsen/logrus"

	m "github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"

	s "github.com/monitprod/send_email/pkg/service"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

func sendEmailHandler(ctx context.Context, payload f.EventPayload) error {

	products, err := getProducts(ctx, payload)
	if err != nil {
		return err
	}

	body, err := mountBody(products)
	if err != nil {
		return err
	}

	usersInfo, err := getUsersAndCount(ctx, payload)
	if err != nil {
		return err
	}

	err = sendEmailForAll(usersInfo.Users, body)
	if err != nil {
		return err
	}

	if fetchMore(payload, *usersInfo.Count) {
		newPayload := iterateExecution(payload)
		runNewFunction(ctx, newPayload)
	}

	return nil
}

func runNewFunction(ctx context.Context, payload f.EventPayload) error {
	fnService := s.NewSendEmailFnServiceImp()

	return fnService.RunNewFunction(ctx, payload)
}

func iterateExecution(payload f.EventPayload) f.EventPayload {
	payload.Execution++
	return payload
}

func fetchMore(p f.EventPayload, countUsers int64) bool {
	return p.UsersPerExecution*p.Execution < countUsers
}

func sendEmailForAll(users *[]m.User, body *string) error {
	emailService := s.NewEmailService()

	err := emailService.SendToMany(*users, s.EmailOptions{Body: *body})

	if err != nil {
		return err
	}

	return nil
}

func mountBody(products *[]m.Product) (*string, error) {
	bodyService := s.NewBodyServiceImp()

	body, err := bodyService.MountBody(products)

	if err != nil {
		log.Errorln("Error on mount body from body service", err)
		return nil, err
	}

	return body, nil
}

func getProducts(ctx context.Context, payload f.EventPayload) (*[]m.Product, error) {
	productRepo := r.NewProductRepositoryMongoDB()
	productService := s.NewProductServiceImp(productRepo, payload)

	products, err := productService.GetProducts(ctx)

	if err != nil {
		log.Errorln("Error while get products from product service", err)
		return nil, err
	}

	return products, nil
}

type UsersInfo struct {
	Users *[]m.User
	Count *int64
}

func getUsersAndCount(ctx context.Context, payload f.EventPayload) (*UsersInfo, error) {
	userRepo := r.NewUserRepositoryMongoDB()
	userService := s.NewUserServiceImp(userRepo, payload)

	users, err := userService.GetUsers(ctx)
	if err != nil {
		log.Errorln("Error while get users from user service", err)
		return nil, err
	}

	count, err := userService.CountUsers(ctx)
	if err != nil {
		log.Errorln("Error while get count from user service", err)
		return nil, err
	}

	return &UsersInfo{users, count}, nil
}
