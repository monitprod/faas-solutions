package service

import (
	"fmt"
	"testing"

	"github.com/monitprod/db_repository/pkg/mock"
)

func TestMountBody(t *testing.T) {

	var bodyService BodyService = newBodyServiceImp()

	str, _ := bodyService.MountBody(&mock.Products)

	fmt.Println(str)

}
