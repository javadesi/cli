package v2action

type RouterGroups []RouterGroup

type RouterGroup {
	GUID string
}

func (actor Actor) GetRouterGroups() (RouterGroups, error) {
	clientRouterGroups, warnings, err := actor.RouterClient.GetRouterGroups()
	if err != nil {
		return nil, err
	}
	return RouterGroups(clientRouterGroups), err
}