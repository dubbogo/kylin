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

package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEquals(t *testing.T) {
	u1 := &User{Username: "user1", Password: "password1"}
	u2 := &User{Username: "user2", Password: "password2"}
	u3 := &User{Username: "user1", Password: "password1"}
	assert.False(t, u1.Equals(u2))
	assert.True(t, u1.Equals(u3))
}

func TestNodeEquals(t *testing.T) {
	n1 := &Node{Name: "node1", Host: "host1", Port: 1234, Database: "db1", Username: "user1", Password: "pass1", Weight: "1", Labels: map[string]string{"label1": "label2"}, Parameters: map[string]string{"param1": "value1"}}
	n2 := &Node{Name: "node2", Host: "host2", Port: 5678, Database: "db2", Username: "user2", Password: "pass2", Weight: "2", Labels: map[string]string{"label3": "label4"}, Parameters: map[string]string{"param2": "value2"}}
	n3 := &Node{Name: "node1", Host: "host1", Port: 1234, Database: "db1", Username: "user1", Password: "pass1", Weight: "1", Labels: map[string]string{"label1": "label2"}, Parameters: map[string]string{"param1": "value1"}}
	assert.False(t, n1.Equals(n2))
	assert.True(t, n1.Equals(n3))
}

func TestRulesEquals(t *testing.T) {
	rules1 := Rules{
		&Rule{Columns: []*ColumnRule{{Name: "col1"}}},
		&Rule{Columns: []*ColumnRule{{Name: "col2"}}},
	}
	rules2 := Rules{
		&Rule{Columns: []*ColumnRule{{Name: "col3"}}},
		&Rule{Columns: []*ColumnRule{{Name: "col4"}}},
	}
	rules3 := Rules{
		&Rule{Columns: []*ColumnRule{{Name: "col1"}}},
		&Rule{Columns: []*ColumnRule{{Name: "col2"}}},
	}
	assert.False(t, rules1.Equals(rules2))
	assert.True(t, rules1.Equals(rules3))
}

func TestTableEquals(t *testing.T) {
	t1 := &Table{
		DbRules:        Rules{&Rule{Columns: []*ColumnRule{{Name: "col1"}}}},
		TblRules:       Rules{&Rule{Columns: []*ColumnRule{{Name: "col2"}}}},
		Topology:       &Topology{"node1", "node2"},
		ShadowTopology: &Topology{"shadow1", "shadow2"},
		Attributes:     map[string]string{"attr1": "value1"},
	}
	t2 := &Table{
		DbRules:        Rules{&Rule{Columns: []*ColumnRule{{Name: "col3"}}}},
		TblRules:       Rules{&Rule{Columns: []*ColumnRule{{Name: "col4"}}}},
		Topology:       &Topology{"node3", "node4"},
		ShadowTopology: &Topology{"shadow3", "shadow4"},
		Attributes:     map[string]string{"attr2": "value2"},
	}
	t3 := &Table{
		DbRules:        Rules{&Rule{Columns: []*ColumnRule{{Name: "col1"}}}},
		TblRules:       Rules{&Rule{Columns: []*ColumnRule{{Name: "col2"}}}},
		Topology:       &Topology{"node1", "node2"},
		ShadowTopology: &Topology{"shadow1", "shadow2"},
		Attributes:     map[string]string{"attr1": "value1"},
	}
	assert.False(t, t1.Equals(t2))
	assert.True(t, t1.Equals(t3))
}
