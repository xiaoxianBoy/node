/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

// Registry holds of all plugable services
type Registry struct {
	factories map[string]ServiceFactory
}

// NewRegistry creates registry of plugable services
func NewRegistry() *Registry {
	return &Registry{
		factories: make(map[string]ServiceFactory),
	}
}

// Register new plugable service
func (registry *Registry) Register(serviceType string, creator ServiceFactory) {
	registry.factories[serviceType] = creator
}

// CreateConnection create plugable service
func (registry *Registry) Create(options Options) (Service, error) {
	createConnection, exists := registry.factories[options.Type]
	if !exists {
		return nil, ErrUnsupportedServiceType
	}

	return createConnection(options)
}
