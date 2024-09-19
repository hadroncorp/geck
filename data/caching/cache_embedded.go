package caching

import (
	"bufio"
	"bytes"
	"context"
	"errors"

	"github.com/allegro/bigcache/v3"
)

const (
	embeddedAppendValueSeparator byte = '\n'
)

type CacheEmbedded struct {
	DB *bigcache.BigCache
}

var _ Cache = (*CacheEmbedded)(nil)

func NewCacheEmbedded(db *bigcache.BigCache) CacheEmbedded {
	return CacheEmbedded{
		DB: db,
	}
}

func (m CacheEmbedded) Set(_ context.Context, key string, value []byte) error {
	return m.DB.Set(key, value)
}

func (m CacheEmbedded) SetMany(_ context.Context, keyValues map[string][]byte) (err error) {
	successKeys := make([]string, 0, len(keyValues))
	defer func() {
		if err == nil {
			return
		}

		// atomic operation
		for _, key := range successKeys {
			_ = m.DB.Delete(key)
		}
	}()
	for k, v := range keyValues {
		if err = m.DB.Set(k, v); err != nil {
			return
		}
		successKeys = append(successKeys, k)
	}
	return nil
}

func (m CacheEmbedded) Append(_ context.Context, key string, value []byte) error {
	return m.DB.Append(key, value)
}

func (m CacheEmbedded) Add(_ context.Context, key string, value []byte) error {
	formattedVal := make([]byte, 0, 1+len(value))
	formattedVal = append(formattedVal, embeddedAppendValueSeparator)
	formattedVal = append(formattedVal, value...)
	return m.DB.Append(key, formattedVal)
}

func (m CacheEmbedded) List(_ context.Context, key string) ([][]byte, error) {
	listBytes, err := m.DB.Get(key)
	if err != nil {
		return nil, err
	} else if len(listBytes) <= 1 {
		return nil, nil
	}

	// remove first item separator
	listBytes = listBytes[1:]
	out := make([][]byte, 0)
	scanner := bufio.NewScanner(bytes.NewReader(listBytes))
	for scanner.Scan() {
		item := scanner.Bytes()
		out = append(out, item)
	}
	return out, nil
}

func (m CacheEmbedded) Get(_ context.Context, key string) ([]byte, error) {
	return m.DB.Get(key)
}

func (m CacheEmbedded) Delete(_ context.Context, key string) error {
	return m.DB.Delete(key)
}

func (m CacheEmbedded) DeleteMany(_ context.Context, keys []string) error {
	errs := make([]error, 0, len(keys))
	for _, key := range keys {
		if err := m.DB.Delete(key); err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return errors.Join(errs...)
}
