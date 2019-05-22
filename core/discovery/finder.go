/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
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

package discovery

import (
	"github.com/mysteriumnetwork/node/market"
)

// finder implements ProposalFinder, which finds proposals from local storage
type finder struct {
	storage *ProposalStorage
}

// NewFinder creates instance of local storage finder
func NewFinder(storage *ProposalStorage) *finder {
	return &finder{storage: storage}
}

// GetProposal fetches service proposal from discovery by exact ID
func (finder *finder) GetProposal(id market.ProposalID) (*market.ServiceProposal, error) {
	for _, proposal := range finder.storage.Proposals() {
		if proposal.UniqueID() == id {
			return &proposal, nil
		}
	}

	return nil, nil
}

// FindProposals fetches currently active service proposals from discovery
func (finder *finder) FindProposals(filter market.ProposalFilter) ([]market.ServiceProposal, error) {
	proposalsFiltered := make([]market.ServiceProposal, 0)
	for _, proposal := range finder.storage.Proposals() {

		if filter.ProviderID != "" && filter.ProviderID != proposal.ProviderID {
			continue
		}
		if filter.ServiceType != "" && filter.ServiceType != proposal.ServiceType {
			continue
		}
		if filter.Location != emptyFilterLocation && !filterByLocation(proposal, filter.Location) {
			continue
		}
		if filter.AccessPolicy != emptyFilterAccessPolicy && !filterByAccessPolicy(proposal, filter.AccessPolicy) {
			continue
		}

		proposalsFiltered = append(proposalsFiltered, proposal)
	}
	return proposalsFiltered, nil
}

var (
	emptyFilterLocation = market.LocationFilter{}
)

func filterByLocation(proposal market.ServiceProposal, filter market.LocationFilter) bool {
	location := proposal.ServiceDefinition.GetLocation()
	if filter.NodeType != "" && filter.NodeType != location.NodeType {
		return false
	}

	return true
}

var (
	emptyFilterAccessPolicy = market.AccessPolicyFilter{}
)

func filterByAccessPolicy(proposal market.ServiceProposal, filter market.AccessPolicyFilter) bool {
	// These proposals accepts all access lists
	if proposal.AccessPolicies == nil {
		return false
	}

	var match bool
	for _, policy := range *proposal.AccessPolicies {
		if filter.ID != "" {
			match = filter.ID == policy.ID
		}
		if filter.Source != "" {
			match = match && filter.Source == policy.Source
		}
		if match == true {
			break
		}
	}
	return match
}
