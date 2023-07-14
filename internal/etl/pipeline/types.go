package pipeline

type ActivityState uint8

const (
	INACTIVE ActivityState = iota
	ACTIVE
	CRASHED
	TERMINATED
)

func (as ActivityState) String() string {
	switch as {
	case INACTIVE:
		return "inactive"

	case ACTIVE:
		return "active"

	case CRASHED:
		return "crashed"

	case TERMINATED:
		return "terminated"
	}

	return "unknown"
}

const (
	// EtlStore error constants
	couldNotCastErr = "could not cast component initializer function to %s constructor type"
	pIDNotFoundErr  = "could not find pipeline ID for %s"
	uuidNotFoundErr = "could not find matching UUID for pipeline entry"

	// ComponentGraph error constants
	cUUIDNotFoundErr = "component with ID %s does not exist within component graph"
	cUUIDExistsErr   = "component with ID %s already exists in component graph"
	edgeExistsErr    = "edge already exists from (%s) to (%s) in component graph"

	emptyPipelineError = "pipeline must contain at least one component"
	// Manager error constants
	unknownCompType = "unknown component type %s provided"

	noAggregatorErr = "aggregator component has yet to be implemented"
)
