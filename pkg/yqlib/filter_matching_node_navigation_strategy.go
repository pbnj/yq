package yqlib

func FilterMatchingNodesNavigationStrategy(value string) NavigationStrategy {
	return &NavigationStrategyImpl{
		visitedNodes: []*NodeContext{},
		followAlias: func(nodeContext NodeContext) bool {
			return true
		},
		autoCreateMap: func(nodeContext NodeContext) bool {
			return false
		},
		visit: func(nodeContext NodeContext) error {
			return nil
		},
		shouldVisitExtraFn: func(nodeContext NodeContext) bool {
			log.Debug("does %v match %v ? %v", nodeContext.Node.Value, value, nodeContext.Node.Value == value)
			return nodeContext.Node.Value == value
		},
	}
}
