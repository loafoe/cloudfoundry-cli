package ccv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
)

// ServiceInstance represents a Cloud Controller Service Instance.
type ServiceInstance struct {
	// GUID is the unique service instance identifier.
	GUID string

	// Name is the name given to the service instance.
	Name string

	// SpaceGUID is the unique identifier of the space that this service instance
	// belongs to.
	SpaceGUID string

	// ServiceGUID is the unique identifier of the service that this service
	// instance belongs to.
	ServiceGUID string

	// ServicePlanGUID is the unique identifier of the service plan that this
	// service instance belongs to.
	ServicePlanGUID string

	// Type is the type of service instance.
	Type constant.ServiceInstanceType

	// Tags is a list of all tags for the service instance.
	Tags []string

	// DashboardURL is the service-broker provided URL to access administrative
	// features of the service instance.
	DashboardURL string

	// RouteServiceURL is the URL of the user-provided service to which requests
	// for bound routes will be forwarded.
	RouteServiceURL string

	// LastOperation is the status of the last operation requested on the service
	// instance.
	LastOperation LastOperation

	// Arbitrary parameters to pass along to the service instance.
	Parameters map[string]interface{}
}

// Managed returns true if the Service Instance is a managed service.
func (serviceInstance ServiceInstance) Managed() bool {
	return serviceInstance.Type == constant.ManagedService
}

// UnmarshalJSON helps unmarshal a Cloud Controller Service Instance response.
func (serviceInstance *ServiceInstance) UnmarshalJSON(data []byte) error {
	var ccServiceInstance struct {
		Metadata internal.Metadata
		Entity   struct {
			Name            string                 `json:"name"`
			SpaceGUID       string                 `json:"space_guid"`
			ServiceGUID     string                 `json:"service_guid"`
			ServicePlanGUID string                 `json:"service_plan_guid"`
			Type            string                 `json:"type"`
			Tags            []string               `json:"tags"`
			DashboardURL    string                 `json:"dashboard_url"`
			RouteServiceURL string                 `json:"route_service_url"`
			LastOperation   LastOperation          `json:"last_operation"`
			Parameters      map[string]interface{} `json:"parameters"`
		}
	}
	err := cloudcontroller.DecodeJSON(data, &ccServiceInstance)
	if err != nil {
		return err
	}

	serviceInstance.GUID = ccServiceInstance.Metadata.GUID
	serviceInstance.Name = ccServiceInstance.Entity.Name
	serviceInstance.SpaceGUID = ccServiceInstance.Entity.SpaceGUID
	serviceInstance.ServiceGUID = ccServiceInstance.Entity.ServiceGUID
	serviceInstance.ServicePlanGUID = ccServiceInstance.Entity.ServicePlanGUID
	serviceInstance.Type = constant.ServiceInstanceType(ccServiceInstance.Entity.Type)
	serviceInstance.Tags = ccServiceInstance.Entity.Tags
	serviceInstance.DashboardURL = ccServiceInstance.Entity.DashboardURL
	serviceInstance.RouteServiceURL = ccServiceInstance.Entity.RouteServiceURL
	serviceInstance.LastOperation = ccServiceInstance.Entity.LastOperation
	serviceInstance.Parameters = ccServiceInstance.Entity.Parameters
	return nil
}

// MarshalJSON converts an user provided service instance into a Cloud Controller user provided service instance.
func (serviceInstance ServiceInstance) MarshalJSON() ([]byte, error) {
	ccObj := struct {
		Name            string                 `json:"name,omitempty"`
		SpaceGUID       string                 `json:"space_guid,omitempty"`
		ServicePlanGUID string                 `json:"service_plan_guid,omitempty"`
		Tags            []string               `json:"tags,omitempty"`
		Parameters      map[string]interface{} `json:"parameters,omitempty"`
	}{
		Name:            serviceInstance.Name,
		SpaceGUID:       serviceInstance.SpaceGUID,
		ServicePlanGUID: serviceInstance.ServicePlanGUID,
		Tags:            serviceInstance.Tags,
		Parameters:      serviceInstance.Parameters,
	}

	return json.Marshal(ccObj)
}

// UserProvided returns true if the Service Instance is a user provided
// service.
func (serviceInstance ServiceInstance) UserProvided() bool {
	return serviceInstance.Type == constant.UserProvidedService
}

type createServiceInstanceRequestBody struct {
	Name            string                 `json:"name"`
	ServicePlanGUID string                 `json:"service_plan_guid"`
	SpaceGUID       string                 `json:"space_guid"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
}

// CreateServiceInstance posts a service instance resource with the provided
// attributes to the api and returns the result.
func (client *Client) CreateServiceInstance(spaceGUID, servicePlanGUID, serviceInstance string, parameters map[string]interface{}, tags []string) (ServiceInstance, Warnings, error) {
	requestBody := createServiceInstanceRequestBody{
		Name:            serviceInstance,
		ServicePlanGUID: servicePlanGUID,
		SpaceGUID:       spaceGUID,
		Parameters:      parameters,
		Tags:            tags,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostServiceInstancesRequest,
		Body:        bytes.NewReader(bodyBytes),
		Query:       url.Values{"accepts_incomplete": {"true"}},
	})
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	var instance ServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &instance,
	}

	err = client.connection.Make(request, &response)
	return instance, response.Warnings, err
}

// CreateServiceInstance posts a service instance resource with the provided
// attributes to the api and returns the result.
func (client *Client) CreateServiceInstanceFromObject(serviceInstance ServiceInstance) (ServiceInstance, Warnings, error) {
	body, err := json.Marshal(serviceInstance)
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostServiceInstancesRequest,
		Body:        bytes.NewReader(body),
		Query:       url.Values{"accepts_incomplete": {"true"}},
	})
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	var instance ServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &instance,
	}

	err = client.connection.Make(request, &response)
	return instance, response.Warnings, err
}

// GetServiceInstance returns the service instance with the given GUID. This
// service can be either a managed or user provided.
func (client *Client) GetServiceInstance(serviceInstanceGUID string) (ServiceInstance, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceInstanceRequest,
		URIParams:   Params{"service_instance_guid": serviceInstanceGUID},
	})
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	var serviceInstance ServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &serviceInstance,
	}

	err = client.connection.Make(request, &response)
	return serviceInstance, response.Warnings, err
}

// GetServiceInstances returns back a list of *managed* Service Instances based
// off of the provided filters.
func (client *Client) GetServiceInstances(filters ...Filter) ([]ServiceInstance, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceInstancesRequest,
		Query:       ConvertFilterParameters(filters),
	})
	if err != nil {
		return nil, nil, err
	}

	var fullInstancesList []ServiceInstance
	warnings, err := client.paginate(request, ServiceInstance{}, func(item interface{}) error {
		if instance, ok := item.(ServiceInstance); ok {
			fullInstancesList = append(fullInstancesList, instance)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceInstance{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullInstancesList, warnings, err
}

// GetSpaceServiceInstances returns back a list of Service Instances based off
// of the space and filters provided. User provided services will be included
// if includeUserProvidedServices is set to true.
func (client *Client) GetSpaceServiceInstances(spaceGUID string, includeUserProvidedServices bool, filters ...Filter) ([]ServiceInstance, Warnings, error) {
	query := ConvertFilterParameters(filters)

	if includeUserProvidedServices {
		query.Add("return_user_provided_service_instances", "true")
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetSpaceServiceInstancesRequest,
		URIParams:   map[string]string{"guid": spaceGUID},
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullInstancesList []ServiceInstance
	warnings, err := client.paginate(request, ServiceInstance{}, func(item interface{}) error {
		if instance, ok := item.(ServiceInstance); ok {
			fullInstancesList = append(fullInstancesList, instance)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceInstance{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullInstancesList, warnings, err
}

// UpdateServiceInstance updates the service instance with the given GUID.
func (client *Client) UpdateServiceInstance(serviceInstance ServiceInstance) (ServiceInstance, Warnings, error) {
	body, err := json.Marshal(serviceInstance)
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PutServiceInstanceRequest,
		URIParams:   Params{"service_instance_guid": serviceInstance.GUID},
		Body:        bytes.NewReader(body),
		Query:       url.Values{"accepts_incomplete": {"true"}},
	})
	if err != nil {
		return ServiceInstance{}, nil, err
	}

	var updatedObj ServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &updatedObj,
	}

	err = client.connection.Make(request, &response)
	return updatedObj, response.Warnings, err
}

// DeleteServiceInstance delete a service instance
func (client *Client) DeleteServiceInstance(guid string, recursive bool, purge bool) (async bool, warn Warnings, err error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteServiceInstanceRequest,
		URIParams: Params{
			"service_instance_guid": guid,
		},
		Query: url.Values{
			"accepts_incomplete": {"true"},
			"recursive":          {fmt.Sprintf("%t", recursive)},
			"purge":              {fmt.Sprintf("%t", purge)},
		},
	})
	if err != nil {
		return false, nil, err
	}

	response := cloudcontroller.Response{}

	err = client.connection.Make(request, &response)
	return response.HTTPResponse.StatusCode == 202, response.Warnings, err
}
