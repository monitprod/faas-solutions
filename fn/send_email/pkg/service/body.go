package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/monitprod/core/pkg/models"

	"github.com/monitprod/send_email/pkg/util"
)

type BodyService interface {
	MountBody(products *[]models.Product) (*string, error)
}

type BodyServiceImp struct {
}

func NewBodyServiceImp() BodyService {
	return &BodyServiceImp{}
}

func (e *BodyServiceImp) MountBody(products *[]models.Product) (*string, error) {

	var templateFile = util.GetRootPath() + "/static/body.html"

	// Read body template HTML
	templateContent, err := ioutil.ReadFile(templateFile)

	if err != nil {
		log.Fatalln("Error while read template file\n", err)
		return nil, err
	}

	templateStr := string(templateContent)

	// Mount HTML Rows of Product
	tableRows, err := mountRows(products)

	if err != nil {
		log.Fatalln("Error while mount rows\n", err)
		return nil, err
	}

	// Add Rows to Tamplate
	bodyStr := strings.Replace(templateStr, "[TABLE_ROWS]", *tableRows, 1)

	return &bodyStr, nil
}

func mountRows(products *[]models.Product) (*string, error) {
	var tableRowsSb strings.Builder

	for _, p := range *products {

		_, err := tableRowsSb.WriteString(
			fmt.Sprintf(
				`<tr>
					<td>%s</td>
					<td>%s</td>
					<td>R$ %.2f</td>
					<td>n/a</td>
				</tr>`,
				p.Title,
				p.Specification,
				float32(p.Price.Value/100),
			),
		)

		if err != nil {
			return nil, err
		}
	}

	tableRows := tableRowsSb.String()

	return &tableRows, nil
}
