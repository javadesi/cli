package router

import (
	"code.cloudfoundry.org/cli/api/router/internal"
)

// RouterGroup represents router group
type RouterGroup struct {
	GUID string `json:"guid"`
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
	//TODO get router groups
	return fullSpacesList, warnings, err
}
