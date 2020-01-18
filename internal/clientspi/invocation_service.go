// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clientspi

import (
	"time"

	"github.com/ahmetmircik/hazelcast-go-client/core"
	"github.com/ahmetmircik/hazelcast-go-client/internal/proto"
)

type InvocationResult interface {
	Result() (*proto.ClientMessage, error)
	ResultWithTimeout(duration time.Duration) (*proto.ClientMessage, error)
}

type InvocationService interface {
	InvokeOnTarget(message *proto.ClientMessage, address core.Address) InvocationResult
}
