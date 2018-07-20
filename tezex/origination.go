/* 
 * TezosAPI
 *
 * BETA Tezos API, this may change frequently
 *
 * OpenAPI spec version: 0.0.2
 * Contact: office@bitfly.at
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tezex

import (
	"time"
)

type Origination struct {

	Hash string `json:"hash,omitempty"`

	Branch string `json:"branch,omitempty"`

	Source string `json:"source,omitempty"`

	PublicKey string `json:"public_key,omitempty"`

	Fee int32 `json:"fee,omitempty"`

	Counter int32 `json:"counter,omitempty"`

	Operations []OriginationOperation `json:"operations,omitempty"`

	Level int32 `json:"level,omitempty"`

	BlockHash string `json:"block_hash,omitempty"`

	Time time.Time `json:"time,omitempty"`
}
