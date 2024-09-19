package caching_test

import (
	"context"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
)

func TestCacheEmbedded_Delete(t *testing.T) {
	db, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// page invalidation
	// 1. Create a reg in cache with nomenclature: {ENTITY_HASH_ID}_ref
	// 2. Iterate through page, append page hash for each entity
	// 3. When entity is mutated, delete both normal entry, each item in ref space (e.g. criteria hashes) and ref key itself
	_ = db.Append("entity_id_0_ref", []byte("criteria_hash_0"))
	_ = db.Append("entity_id_0_ref", []byte("criteria_hash_1"))
	_ = db.Append("entity_id_1_ref", []byte("criteria_hash_0"))
	out, _ := db.Get("criteria_hash")
	t.Log(string(out))
}
