package routing

type routingGroup struct {
	Post  post
	Topic topic
}

var Group = new(routingGroup)
