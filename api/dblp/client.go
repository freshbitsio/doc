//-----------------------------------------------------------------------------
// Interface to the DBLP publication archive
//-----------------------------------------------------------------------------

package dblp

type DBLP struct {
	Args [][]string
	Id string
	Query string
}

// Retrieve publication or data resource.
func (d DBLP) Get (id string) ([]byte, error) {
	return nil, nil
}

// Get resource metadata.
func (d DBLP) Info (id string) ([]byte, error) {
	return nil, nil
}

// Search for resources.
func (d DBLP) Search (query string, args [][]string) ([]byte, error) {
	return nil, nil
}