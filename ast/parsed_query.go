package ast

// ParsedQuery represents a parsed query where
// bind locations are precompued for fast substitutions.
type ParsedQuery struct {
	Query         string
	bindLocations []bindLocation
}

type bindLocation struct {
	offset, length int
}
