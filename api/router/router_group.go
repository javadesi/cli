package router

import (
	"code.cloudfoundry.org/cli/api/router/internal"
)

// RouterGroup represents router group
type RouterGroup struct {
	GUID            string `json:"guid"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	ReservablePorts string `json:"reservable_ports"`
}

// GetRouterGroups returns a list of RouterGroups
func (client *Client) GetRouterGroups() ([]RouterGroup, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetRouterGroups,
	})

	if err != nil {
		return nil, err
	}
	var fullRouterGroupList []RouterGroup

	var response = Response{
		Result: &fullRouterGroupList,
	}
	err = client.connection.Make(request, &response)
	return fullRouterGroupList, err
}
