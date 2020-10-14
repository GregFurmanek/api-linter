// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0164

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules/internal/utils"
	"github.com/jhump/protoreflect/desc"
)

// Undelete methods should not have an HTTP body.
var httpBody = &lint.MethodRule{
	Name:   lint.NewRuleName(164, "http-body"),
	OnlyIf: isUndeleteMethod,
	LintMethod: func(m *desc.MethodDescriptor) []lint.Problem {
		// Establish that the RPC has correct HTTP body.
		for _, httpRule := range utils.GetHTTPRules(m) {
			if httpRule.Body != "*" {
				return []lint.Problem{{
					Message:    `Undelete methods should use "*" as the HTTP body.`,
					Descriptor: m,
				}}
			}
		}

		return nil
	},
}