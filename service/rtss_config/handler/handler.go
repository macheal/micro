package handler

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"gitee.com/smartsteps/go-micro/v2/client"
	cr "gitee.com/smartsteps/go-micro/v2/config/reader"
	jr "gitee.com/smartsteps/go-micro/v2/config/reader/json"
	"gitee.com/smartsteps/go-micro/v2/config/source"
	"github.com/patrickmn/go-cache"

	"gitee.com/smartsteps/go-micro-plugins/store/mongo"
	"gitee.com/smartsteps/go-micro/v2/errors"
	"gitee.com/smartsteps/go-micro/v2/store"
	pb "github.com/micro/micro/v2/service/rtss_config/proto"
)

var (
	PathSplitter = "."
	WatchTopic   = "go.micro.rtss_config.events"
	watchers     = make(map[string][]*watcher)

	// we now support json only
	reader = jr.NewReader()
	mtx    sync.RWMutex
)

type Config struct {
	Store store.Store
	Cache     *cache.Cache
}

func setNamespaceReal(ctx context.Context, v string) string {
	if len(v) == 0 {
		return "/rtss"
	}
	return v
}

func setKey(ctx context.Context, v string, _path string) string {
	return _path
}

func CheckNamespacePath(id, namespace, path string) error {
	if len(namespace) == 0 {
		return errors.BadRequest(id, "invalid namespace")
	}

	if len(path) == 0 {
		return errors.BadRequest(id, "invalid path")
	}

	if !strings.HasPrefix(path, namespace) {
		return errors.BadRequest(id, "invalid request, the path must contain the namespace as a prefix")
	}
	return nil
}

func (c *Config) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	if err := CheckNamespacePath("go.micro.rtss_config.Read", req.Namespace, req.Path); err != nil {
		return err
	}

	key := setKey(ctx, req.Namespace, req.Path)

	// read cache
	v, ok := c.Cache.Get(key)
	if ok {
		obj := v.(*pb.Change)
		rsp.Change = obj
		return nil
	}

	ch, err := c.Store.Read(key)
	if err != nil {
		if err == store.ErrNotFound || err.Error() == "not found" {
			return errors.NotFound("go.micro.rtss_config.Read", "not found key: %s", key)
		}
		return errors.InternalServerError("go.micro.rtss_config.Read", "read key %s error: %v", key, err)
	}
	// mongo store 实现问题, 需判断返回空的情况
	if ch == nil {
		return errors.NotFound("go.micro.rtss_config.Read", "not found key: %s", key)
	}

	rsp.Change = new(pb.Change)

	// Unmarshal value
	if err = json.Unmarshal(ch[0].Value, rsp.Change); err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Read", "unmarshal key %v value error: %v", key, err)
	}

	// if dont need path, we return all of the data
	if len(req.Path) == 0 {
		return nil
	}

	return nil
}

func (c *Config) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	if req.Change == nil || req.Change.ChangeSet == nil {
		return errors.BadRequest("go.micro.rtss_config.Create", "invalid change")
	}

	if err := CheckNamespacePath("go.micro.rtss_config.Create", req.Change.Namespace, req.Change.Path); err != nil {
		return err
	}

	req.Change.ChangeSet.Timestamp = time.Now().Unix()

	namespace := setNamespaceReal(ctx, req.Change.Namespace)
	key := req.Change.Path

	record := &store.Record{
		Key: key,
	}

	var err error
	record.Value, err = json.Marshal(req.Change)
	if err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Create", "marshal key %v value error: %v", err)
	}

	// set cache
	c.Cache.Set(key, req.Change, cache.DefaultExpiration)

	if err := c.Store.Write(record); err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Create", "write key %s error: %v", err)
	}

	_ = publish(ctx, &pb.WatchResponse{Namespace: namespace, Path: key, ChangeSet: req.Change.ChangeSet})

	return nil
}

func (c *Config) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	if req.Change == nil || req.Change.ChangeSet == nil {
		return errors.BadRequest("go.micro.rtss_config.Update", "invalid change")
	}

	if err := CheckNamespacePath("go.micro.rtss_config.Update", req.Change.Namespace, req.Change.Path); err != nil {
		return err
	}

	// set the changeset timestamp
	req.Change.ChangeSet.Timestamp = time.Now().Unix()

	oldCh := &pb.Change{}

	namespace := setNamespaceReal(ctx, req.Change.Namespace)
	key := setKey(ctx, req.Change.Namespace, req.Change.Path)

	// Get the current change set
	var record *store.Record
	records, err := c.Store.Read(key)
	if err != nil {
		if err.Error() != "not found" {
			return errors.NotFound("go.micro.rtss_config.Update", "read old value error: %v", err)
		}
		// create new record
		record = new(store.Record)
		record.Key = key
	} else {
		// mongo store 实现问题, 需判断返回空的情况
		if records == nil {
			return errors.NotFound("go.micro.rtss_config.Update", "read old value error: %v", err)
		}
		// Unmarshal value
		if err := json.Unmarshal(records[0].Value, oldCh); err != nil {
			return errors.InternalServerError("go.micro.rtss_config.Read", "unmarshal key %s value error: %v", key, err)
		}
		record = records[0]
	}

	record.Value, err = json.Marshal(req.Change)
	if err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Update", "marshal error: %v", err)
	}

	// set cache
	c.Cache.Set(key, req.Change, cache.DefaultExpiration)

	if err := c.Store.Write(record); err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Update", "write key %s error: %v", key, err)
	}

	_ = publish(ctx, &pb.WatchResponse{Namespace: namespace, Path: key, ChangeSet: req.Change.ChangeSet})

	return nil
}

func (c *Config) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	if req.Change == nil {
		return errors.BadRequest("go.micro.rtss_config.Delete", "invalid change")
	}

	if err := CheckNamespacePath("go.micro.rtss_config.Delete", req.Change.Namespace, req.Change.Path); err != nil {
		return err
	}

	if req.Change.ChangeSet == nil {
		req.Change.ChangeSet = &pb.ChangeSet{}
	}

	req.Change.ChangeSet.Timestamp = time.Now().Unix()

	namespace := setNamespaceReal(ctx, req.Change.Namespace)
	key := setKey(ctx, req.Change.Namespace, req.Change.Path)

	// delete cache
	c.Cache.Delete(key)

	if err := c.Store.Delete(key); err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Delete", "delete key %s error: %v", key, err)
	}

	_ = publish(ctx, &pb.WatchResponse{Namespace: namespace, Path: key, ChangeSet: nil})

	return nil
}

func (c *Config) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) (err error) {
	list, err := c.Store.List(
		store.ListPrefix(req.Namespace),
		store.ListSuffix(req.Suffix),
		mongo.SetListSubstr(req.Substr))
	if err != nil {
		return errors.BadRequest("go.micro.rtss_config.List", "query value error: %v", err)
	}

	for _, v := range list {
		//if !strings.HasPrefix(v, req.Namespace) {
		//	continue
		//}

		rec, err := c.Store.Read(v)
		if err != nil {
			return errors.InternalServerError("go.micro.rtss_config.Read", "read key %s error: %v", v, err)
		}
		if rec == nil {
			continue
		}

		ch := &pb.Change{}
		if err := json.Unmarshal(rec[0].Value, ch); err != nil {
			return errors.InternalServerError("go.micro.rtss_config.Read", "unmarshal key %s value error: %v", rec[0].Key, err)
		}

		if ch.ChangeSet != nil {
			ch.ChangeSet.Data = string(ch.ChangeSet.Data)
		}

		rsp.Values = append(rsp.Values, ch)
	}

	return nil
}

func (c *Config) Watch(ctx context.Context, req *pb.WatchRequest, stream pb.Config_WatchStream) error {
	if len(req.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.Watch", "invalid id")
	}

	namespace := setNamespaceReal(ctx, req.Namespace)

	watch, err := Watch(namespace)
	if err != nil {
		return errors.InternalServerError("go.micro.rtss_config.Watch", "watch error: %v", err)
	}
	defer watch.Stop()

	go func() {
		select {
		case <-ctx.Done():
			watch.Stop()
			stream.Close()
		}
	}()

	for {
		ch, err := watch.Next()
		if err != nil {
			return errors.InternalServerError("go.micro.rtss_config.Watch", "listen the Next error: %v", err)
		}

		if err := stream.Send(ch); err != nil {
			return errors.InternalServerError("go.micro.rtss_config.Watch", "send the Change error: %v", err)
		}
	}
}

// Used as a subscriber between config services for events
func Watcher(ctx context.Context, ch *pb.WatchResponse) error {
	mtx.RLock()
	for _, sub := range watchers[ch.Namespace] {
		select {
		case sub.next <- ch:
		case <-time.After(time.Millisecond * 100):
		}
	}
	mtx.RUnlock()
	return nil
}

func merge(ch ...*source.ChangeSet) (*source.ChangeSet, error) {
	return reader.Merge(ch...)
}

func values(ch *source.ChangeSet) (cr.Values, error) {
	return reader.Values(ch)
}

// publish a change
func publish(ctx context.Context, ch *pb.WatchResponse) error {
	req := client.NewMessage(WatchTopic, ch)
	return client.Publish(ctx, req)
}
