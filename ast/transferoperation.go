/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"encoding/json"

	"cadence-parser/errors"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=TransferOperation

type TransferOperation uint

const (
	TransferOperationUnknown TransferOperation = iota
	TransferOperationCopy
	TransferOperationMove
	TransferOperationMoveForced
)

func TransferOperationCount() int {
	return len(_TransferOperation_index) - 1
}

func (k TransferOperation) Operator() string {
	switch k {
	case TransferOperationCopy:
		return "="
	case TransferOperationMove:
		return "<-"
	case TransferOperationMoveForced:
		return "<-!"
	}

	panic(errors.NewUnreachableError())
}

func (k TransferOperation) IsMove() bool {
	switch k {
	case TransferOperationMove, TransferOperationMoveForced:
		return true
	}

	return false
}

func (k TransferOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}
