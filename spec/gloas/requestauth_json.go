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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethpandaops/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// MaxDataSize is the maximum number of bytes in the request auth data field.
const MaxDataSize = 4096

// requestAuthJSON is the spec representation of the struct.
type requestAuthJSON struct {
	Data string `json:"data"`
	Slot string `json:"slot"`
}

// MarshalJSON implements json.Marshaler.
func (r *RequestAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(&requestAuthJSON{
		Data: fmt.Sprintf("%#x", r.Data),
		Slot: fmt.Sprintf("%d", r.Slot),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *RequestAuth) UnmarshalJSON(input []byte) error {
	var data requestAuthJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	if data.Data == "" {
		return errors.New("data missing")
	}
	dataBytes, err := hex.DecodeString(strings.TrimPrefix(data.Data, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid data")
	}
	if len(dataBytes) > MaxDataSize {
		return errors.New("data too long")
	}
	r.Data = dataBytes

	if data.Slot == "" {
		return errors.New("slot missing")
	}
	slot, err := strconv.ParseUint(data.Slot, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid slot")
	}
	r.Slot = phase0.Slot(slot)

	return nil
}
