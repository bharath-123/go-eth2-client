// Copyright © 2026 Attestant Limited.
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

package gloas

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethpandaops/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// builderPreferencesJSON is the spec representation of the struct.
type builderPreferencesJSON struct {
	MaxTrustedBid string `json:"max_trusted_bid"`
}

// MarshalJSON implements json.Marshaler.
func (b *BuilderPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(&builderPreferencesJSON{
		MaxTrustedBid: fmt.Sprintf("%d", b.MaxTrustedBid),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *BuilderPreferences) UnmarshalJSON(input []byte) error {
	var data builderPreferencesJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	if data.MaxTrustedBid == "" {
		return errors.New("max trusted bid missing")
	}
	maxTrustedBid, err := strconv.ParseUint(data.MaxTrustedBid, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid max trusted bid")
	}
	b.MaxTrustedBid = phase0.Gwei(maxTrustedBid)

	return nil
}
