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

package nearcachespi

import (
	"github.com/ahmetmircik/hazelcast-go-client/config/property"
	"github.com/ahmetmircik/hazelcast-go-client/core"
	"github.com/ahmetmircik/hazelcast-go-client/internal/clientspi"
	"github.com/ahmetmircik/hazelcast-go-client/internal/nearcache"
	"github.com/ahmetmircik/hazelcast-go-client/internal/nearcache/internal"
	"github.com/ahmetmircik/hazelcast-go-client/internal/nearcache/internal/invalidation"
	"github.com/ahmetmircik/hazelcast-go-client/serialization/spi"
)

func NewDefaultNearCacheManager(service spi.SerializationService,
	properties *property.HazelcastProperties) *internal.DefaultNearCacheManager {
	return internal.NewDefaultNearCacheManager(service, properties)
}

func NewRepairingTask(properties *property.HazelcastProperties, service spi.SerializationService,
	partitionService clientspi.PartitionService, invocationService clientspi.InvocationService,
	cluster core.Cluster, localUUID string) nearcache.RepairingTask {
	return invalidation.NewRepairingTask(properties, service, partitionService, invocationService, cluster, localUUID)
}
