package servicediscovery

import (
	"encoding/json"
	"errors"

	routing "github.com/qiangxue/fasthttp-routing"
)

func ValidateServiceGroupBody(c *routing.Context) (ServiceGroup, error) {
	var s ServiceGroup
	err := json.Unmarshal(c.Request.Body(), &s)

	if s.Name == "" {
		return ServiceGroup{}, errors.New(`{"error" : true, "msg": "Missing body parameters."}`)
	}
	if err != nil {
		return ServiceGroup{}, errors.New(`{"error" : true, "msg": "Error parsing body."}`)
	}
	
	return s, nil
}

func ValidateServiceBody(c *routing.Context) (Service, error) {
	var s Service
	err := json.Unmarshal(c.Request.Body(), &s)

	if s.Name == "" || len(s.Hosts) == 0 || s.MatchingURI == "" || s.ToURI == "" || s.APIDocumentation == "" {
		return Service{}, errors.New(`{"error" : true, "msg": "Missing body parameters."}`)
	}
	if err != nil {
		return Service{}, errors.New(`{"error" : true, "msg": "Error parsing body."}`)
	}

	s.NormalizeService()
	
	return s, nil
}

func ValidateServiceExists(s Service) (Service, error) {
	service, err := sd.FindService(s)
	
	if err != nil {
		return Service{}, errors.New(`{"error":true, "msg":"Resource not found"}`)
	}

	return service, nil
}
