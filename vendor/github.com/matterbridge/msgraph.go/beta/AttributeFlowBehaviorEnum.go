// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttributeFlowBehavior undocumented
type AttributeFlowBehavior int

const (
	// AttributeFlowBehaviorVFlowWhenChanged undocumented
	AttributeFlowBehaviorVFlowWhenChanged AttributeFlowBehavior = 0
	// AttributeFlowBehaviorVFlowAlways undocumented
	AttributeFlowBehaviorVFlowAlways AttributeFlowBehavior = 1
)

// AttributeFlowBehaviorPFlowWhenChanged returns a pointer to AttributeFlowBehaviorVFlowWhenChanged
func AttributeFlowBehaviorPFlowWhenChanged() *AttributeFlowBehavior {
	v := AttributeFlowBehaviorVFlowWhenChanged
	return &v
}

// AttributeFlowBehaviorPFlowAlways returns a pointer to AttributeFlowBehaviorVFlowAlways
func AttributeFlowBehaviorPFlowAlways() *AttributeFlowBehavior {
	v := AttributeFlowBehaviorVFlowAlways
	return &v
}