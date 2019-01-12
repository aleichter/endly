package s3

import (

	"github.com/viant/endly"
	"github.com/viant/endly/system/cloud/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

const (
	//ServiceID aws s3 service id.
	ServiceID = "aws/s3"
)


//no operation service
type service struct {
	*endly.AbstractService

}


func (s *service) registerRoutes() {
	client := &s3.S3{}
	routes, err := aws.BuildRoutes(client, getClient)
	if err != nil {
		log.Printf("unable register service %v actions: %v\n", ServiceID, err)
		return
	}
	for _, route := range routes {
		route.OnRawRequest = setClient
		s.Register(route)
	}
}





//New creates a new AWS S3 service.
func New() endly.Service {
	var result = &service{
		AbstractService: endly.NewAbstractService(ServiceID),
	}
	result.AbstractService.Service = result
	result.registerRoutes()
	return result
}
