// Copyright 2018 Google LLC. All Rights Reserved.
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

package core

import (
	"context"
	"testing"

	ct "github.com/rarimo/certificate-transparency-go"
)

func TestVerifyConsistencyEmptyHead(t *testing.T) {
	controller := new(Controller)
	if controller.verifyConsistency(context.Background(), 0, []byte("abc"), &ct.SignedTreeHead{TreeSize: 100}) != nil {
		t.Errorf("verifyConsistency should always succeed given empty root")
	}
}
