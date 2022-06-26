/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package merge

import (
	"container/heap"
	"database/sql"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	consts "github.com/arana-db/arana/pkg/constants/mysql"
	"github.com/arana-db/arana/pkg/merge/impl/order"
	"github.com/arana-db/arana/pkg/mysql"
	"github.com/arana-db/arana/pkg/mysql/rows"
	"github.com/arana-db/arana/pkg/proto"
)

func TestPriorityQueue(t *testing.T) {
	fields := []proto.Field{
		mysql.NewField("id", consts.FieldTypeLong),
		mysql.NewField("score", consts.FieldTypeLong),
	}
	items := []order.OrderByItem{
		{"id", false},
		{"score", true},
	}

	r1 := &order.RowItem{rows.NewTextVirtualRow(fields, []proto.Value{
		int64(1),
		int64(80),
	}), 1}
	r2 := &order.RowItem{rows.NewTextVirtualRow(fields, []proto.Value{
		int64(2),
		int64(75),
	}), 1}
	r3 := &order.RowItem{rows.NewTextVirtualRow(fields, []proto.Value{
		int64(1),
		int64(90),
	}), 1}
	r4 := &order.RowItem{rows.NewTextVirtualRow(fields, []proto.Value{
		int64(3),
		int64(85),
	}), 1}
	pq := NewPriorityQueue([]*order.RowItem{
		r1, r2, r3, r4,
	}, items)

	assertScorePojoEquals(t, fakeScorePojo{
		id:    int64(1),
		score: int64(90),
	}, heap.Pop(pq).(*order.RowItem).Row)

	assertScorePojoEquals(t, fakeScorePojo{
		id:    int64(1),
		score: int64(80),
	}, heap.Pop(pq).(*order.RowItem).Row)

	assertScorePojoEquals(t, fakeScorePojo{
		id:    int64(2),
		score: int64(75),
	}, heap.Pop(pq).(*order.RowItem).Row)

	assertScorePojoEquals(t, fakeScorePojo{
		id:    int64(3),
		score: int64(85),
	}, heap.Pop(pq).(*order.RowItem).Row)

}

func assertScorePojoEquals(t *testing.T, expected fakeScorePojo, actual proto.Row) {
	var pojo fakeScorePojo
	err := scanScorePojo(actual, &pojo)
	assert.NoError(t, err)
	assert.Equal(t, expected, pojo)
}

type fakeScorePojo struct {
	id    int64
	score int64
}

func scanScorePojo(row proto.Row, dest *fakeScorePojo) error {
	s := make([]proto.Value, 2)
	if err := row.Scan(s); err != nil {
		return err
	}

	var (
		id    sql.NullInt64
		score sql.NullInt64
	)
	_, _ = id.Scan(s[0]), score.Scan(s[1])

	dest.id = id.Int64
	dest.score = score.Int64

	return nil
}
