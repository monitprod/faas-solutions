package handler

import (
	"context"
	"log"
	"os"

	m "github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"
	coreService "github.com/monitprod/core/pkg/service"
	"github.com/monitprod/core/pkg/util/local"

	c "github.com/monitprod/core/pkg/constant"
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

func iterateExecution(payload f.EventPayload) f.EventPayload {
	payload.Execution++
	return payload
}

func runNewFunction(ctx context.Context, payload f.EventPayload) error {
	isLocal, _ := ctx.Value(c.IsLocal).(bool)
	localMainFunc, _ := ctx.Value(c.LocalMainFunc).(local.LocalFunc)

	builder := coreService.FunctionBuilder{
		IsLocal:   isLocal,
		LocalFunc: localMainFunc,
		Payload:   payload.ToMap(),
	}

	if !isLocal {
		builder.ServiceOptions = &coreService.ServiceOptions{
			Region:       os.Getenv("SE_FUNCTION_REGION"),
			FunctionName: os.Getenv("SE_FUNCTION_NAME"),
			Credentials: coreService.Credentials{
				AccessKey: os.Getenv("SE_AWS_ACCESS_KEY_ID"),
				Secret:    os.Getenv("SE_AWS_SECRET"),
			},
		}
	}

	funcService := coreService.NewFunctionServiceImp(builder)

	return funcService.Exec()
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
		log.Fatalln("Error on mount body from body service", err)
		return nil, err
	}

	return body, nil
}

func getProducts(ctx context.Context, payload f.EventPayload) (*[]m.Product, error) {
	productRepo := r.NewProductRepositoryMongoDB()
	productService := s.NewProductServiceImp(productRepo, payload)

	products, err := productService.GetProducts(ctx)

	if err != nil {
		log.Fatalln("Error while get products from product service", err)
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
		log.Fatalln("Error while get users from user service", err)
		return nil, err
	}

	count, err := userService.CountUsers(ctx)
	if err != nil {
		log.Fatalln("Error while get count from user service", err)
		return nil, err
	}

	return &UsersInfo{users, count}, nil
}
