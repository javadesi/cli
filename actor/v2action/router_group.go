package v2action

import (
	"errors"

	"code.cloudfoundry.org/cli/api/router"
)

type RouterGroup router.RouterGroup

func (actor Actor) GetRouterGroupByName(routerGroupName string, client RouterClient) (RouterGroup, error) {
	client.GetRouterGroups()
	return RouterGroup{}, errors.New("Not a real router group")
}
