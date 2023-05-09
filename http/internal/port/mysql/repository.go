package mysql

import (
	"fmt"
	"github.com/sweetie-pie/line-recommendation/internal/model"
)

// InsertRoute
// insert a route by inserting nodes of that route.
func (m *MySQL) InsertRoute(route *model.Route) error {
	return m.conn.Create(route).Error
}

// InsertNode
// insert a new route to system.
func (m *MySQL) InsertNode(node *model.Node) error {
	return m.conn.Create(node).Error
}

// InsertSearch
// insert a new user search entity.
func (m *MySQL) InsertSearch(search *model.Search) error {
	return m.conn.Create(search).Error
}

// GetRoutes of our system.
func (m *MySQL) GetRoutes() ([]*model.Route, error) {
	return nil, nil
}

// GetNode by its id.
func (m *MySQL) GetNode(id uint) (*model.Node, error) {
	node := new(model.Node)

	if err := m.conn.First(&node, id).Error; err != nil {
		return nil, err
	}

	if node.ID != id {
		return nil, fmt.Errorf("node not found")
	}

	return node, nil
}

// GetSearches of our users.
func (m *MySQL) GetSearches() ([]*model.Search, error) {
	return nil, nil
}
