package service

import (
	"fmt"
	"testing"

	"github.com/monitprod/core/pkg/mock"
)

func TestMountBody(t *testing.T) {

	var bodyService BodyService = NewBodyServiceImp()

	str, _ := bodyService.MountBody(&mock.Products)

	fmt.Println(str)

}
