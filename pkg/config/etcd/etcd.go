/*
 * //
 * // Licensed to Apache Software Foundation (ASF) under one or more contributor
 * // license agreements. See the NOTICE file distributed with
 * // this work for additional information regarding copyright
 * // ownership. Apache Software Foundation (ASF) licenses this file to you under
 * // the Apache License, Version 2.0 (the "License"); you may
 * // not use this file except in compliance with the License.
 * // You may obtain a copy of the License at
 * //
 * // http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing,
 * // software distributed under the License is distributed on an
 * // "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * // KIND, either express or implied.  See the License for the
 * // specific language governing permissions and limitations
 * // under the License.
 * //
 */

package etcd

import (
	"context"
	"github.com/arana-db/arana/pkg/config"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math"
	"sync"
	"time"
)

import (
	etcdv3 "github.com/dubbogo/gost/database/kv/etcd/v3"

	"github.com/pkg/errors"
)

func init() {
	config.Register(&storeOperate{})
}

type storeOperate struct {
	client    *etcdv3.Client
	lock      *sync.RWMutex
	receivers map[string]*watcher
	cancelfs  []context.CancelFunc
}

func (c *storeOperate) Init(options map[string]interface{}) error {
	endpoints, _ := options["endpoints"].([]string)
	tmpClient, err := etcdv3.NewConfigClientWithErr(
		etcdv3.WithName(etcdv3.RegistryETCDV3Client),
		etcdv3.WithTimeout(10*time.Second),
		etcdv3.WithEndpoints(endpoints...),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize etcd client")
	}

	c.client = tmpClient
	c.lock = &sync.RWMutex{}
	c.receivers = make(map[string]*watcher)

	return nil
}

func (c *storeOperate) Save(key string, val []byte) error {
	return c.client.Put(key, string(val))
}

func (c *storeOperate) Get(key string) ([]byte, error) {
	v, err := c.client.Get(key)
	if err != nil {
		return nil, err
	}

	return []byte(v), nil
}

func (c *storeOperate) Delete(key string) error {
	return c.client.Delete(key)
}

type watcher struct {
	revision  int64
	lock      *sync.RWMutex
	receivers []chan []byte
	ch        clientv3.WatchChan
}

func (w *watcher) run(ctx context.Context) {
	for {
		select {
		case resp := <-w.ch:
			for i := range resp.Events {
				event := resp.Events[i]
				if event.Type == mvccpb.DELETE || event.Kv.ModRevision <= w.revision {
					continue
				}
				w.revision = event.Kv.ModRevision
				for p := range w.receivers {
					w.receivers[p] <- event.Kv.Value
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *storeOperate) Watch(key string) (<-chan []byte, error) {
	defer c.lock.Unlock()
	c.lock.Lock()
	if _, ok := c.receivers[key]; !ok {
		watchCh, err := c.client.Watch(key)
		if err != nil {
			return nil, err
		}
		w := &watcher{
			revision:  math.MinInt64,
			lock:      &sync.RWMutex{},
			receivers: make([]chan []byte, 0, 2),
			ch:        watchCh,
		}
		ctx, cancel := context.WithCancel(context.Background())
		go w.run(ctx)
		c.cancelfs = append(c.cancelfs, cancel)
		c.receivers[key] = w
	}

	w := c.receivers[key]

	defer w.lock.Unlock()
	w.lock.Lock()

	rec := make(chan []byte)
	c.receivers[key].receivers = append(c.receivers[key].receivers, rec)
	return rec, nil
}

func (c *storeOperate) Name() string {
	return "etcd"
}

func (c *storeOperate) Close() error {
	for _, f := range c.cancelfs {
		f()
	}

	return nil
}
