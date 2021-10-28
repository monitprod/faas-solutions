package user

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"

	s "github.com/monitprod/user_api/pkg/service"
)

func HandleUserRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	user, err := models.UserFromJson(request.Body)
	if err != nil {
		return nil, err
	}

	err = sign(ctx, user)
	if err != nil {
		return nil, errors.New("error while sign user")
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func sign(ctx context.Context, user *models.User) error {

	userRepo := r.NewUserRepositoryMongoDB()
	userService := s.NewUserServiceImp(userRepo)

	return userService.SignUser(ctx, user)
}
