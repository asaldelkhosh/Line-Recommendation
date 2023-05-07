package mysql

import (
	"github.com/sweetie-pie/line-recommendation/internal/model"
)

// InsertRoute
// insert a route by inserting nodes of that route.
func (m *MySQL) InsertRoute(route *model.Route) error {
	return nil
}

// InsertNode
// insert a new route to system.
func (m *MySQL) InsertNode(node *model.Node) error {
	return nil
}

// InsertSearch
// insert a new user search entity.
func (m *MySQL) InsertSearch(search *model.Search) error {
	return nil
}

// GetRoutes of our system.
func (m *MySQL) GetRoutes() ([]*model.Route, error) {
	return nil, nil
}

// GetSearches of our users.
func (m *MySQL) GetSearches() ([]*model.Search, error) {
	return nil, nil
}
