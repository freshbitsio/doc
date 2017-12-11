//-----------------------------------------------------------------------------
// Interface to the PLOS publication archive
//-----------------------------------------------------------------------------

package plos

type PLOS struct {
	Args [][]string
	Id string
	Query string
}

// Retrieve publication or data resource.
func (p PLOS) Get (id string) ([]byte, error) {
	return nil, nil
}

// Get resource metadata.
func (p PLOS) Info (id string) ([]byte, error) {
	return nil, nil
}

// Search for resources.
func (p PLOS) Search (query string, args [][]string) ([]byte, error) {
	return nil, nil
}