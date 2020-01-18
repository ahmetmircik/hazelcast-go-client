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

package nearcache

import (
	"time"

	"github.com/ahmetmircik/hazelcast-go-client/internal/eviction"
)

var TimeNotSet = time.Now()

const (
	NotReserved   int64 = -1
	Reserved      int64 = -2
	UpdateStarted int64 = -3
	ReadPermitted int64 = -4
)

type Record interface {
	eviction.Evictable
	eviction.Expirable
	SetKey(key interface{})
	Key() interface{}
	SetValue(value interface{})
	SetCreationTime(time time.Time)
	SetAccessTime(time time.Time)
	IsIdleAt(maxIdle time.Duration, now time.Time) bool
	IncrementAccessHit()
	RecordState() int64
	CasRecordState(expect int64, update int64) bool
	PartitionID() int32
	SetPartitionID(partitionID int32)
	InvalidationSequence() int64
	SetInvalidationSequence(sequence int64)
	SetUUID(UUID string)
	HasSameUUID(UUID string) bool
	LessThan(comparator RecordComparator, record Record) bool
}
