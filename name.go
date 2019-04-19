// Copyright 2017-2019 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Name represents an ENS name
type Name struct {
	client *ethclient.Client
	// Name is the full name of an ENS domain e.g. foo.bar.eth
	Name string
	// Domain is the domain of an ENS domain e.g. bar.eth
	Domain string
	// Label is the name part of an ENS domain e.g. foo
	Label string
}

// NewName creates an ENS name
func NewName(client *ethclient.Client, name string) (*Name, error) {
	name = NormaliseDomain(name)
	domain := Domain(name)
	label, err := DomainPart(name, 1)
	if err != nil {
		return nil, err
	}
	return &Name{
		client: client,
		Name:   name,
		Domain: domain,
		Label:  label,
	}, nil
}
