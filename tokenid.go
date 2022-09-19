// Copyright 2019-2021 Weald Technology Trading
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
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
)

// DeriveTokenIdFromENSDomain derive token_id from the ENS domain
func DeriveTokenIdFromENSDomain(domain string) (string, error) {
	domain, err := NormaliseDomain(domain)
	if err != nil {
		return "", err
	}
	domain, err = DomainPart(domain, 1)
	if err != nil {
		return "", err
	}
	labelHash, err := LabelHash(domain)
	hash := fmt.Sprintf("0x%x", labelHash)
	tokenId, ok := math.ParseBig256(hash)
	if !ok {
		return "", err
	}
	return tokenId.String(), nil

}
