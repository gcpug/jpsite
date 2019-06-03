package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

func TestOrganizationStore_Create(t *testing.T) {
	ctx := context.Background()

	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(uuid.New().String()))
	if err != nil {
		t.Fatal(err)
	}

	ostore, err := NewOrganizationStore(ctx, ds)
	if err != nil {
		t.Fatal(err)
	}

	e := &Organization{}
	so, err := ostore.Create(ctx, ostore.NameKey(ctx, "hoge"), e)
	if err != nil {
		t.Fatal(err)
	}
	if so.Key == nil {
		t.Fatalf("Organization.Key is nil")
	}
}

func TestOrganizationStore_Get(t *testing.T) {
	ctx := context.Background()

	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(uuid.New().String()))
	if err != nil {
		t.Fatal(err)
	}

	ostore, err := NewOrganizationStore(ctx, ds)
	if err != nil {
		t.Fatal(err)
	}

	e := &Organization{}
	key := ostore.NameKey(ctx, "hoge")
	so, err := ostore.Create(ctx, key, e)
	if err != nil {
		t.Fatal(err)
	}
	if so.Key == nil {
		t.Fatalf("Organization.Key is nil")
	}

	o, err := ostore.Get(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	if e, g := key, o.Key; e != g {
		t.Fatalf("expected Key = %v, got %v", e, g)
	}
}

func TestOrganizationStore_ListAll(t *testing.T) {
	ctx := context.Background()

	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(uuid.New().String()))
	if err != nil {
		t.Fatal(err)
	}

	ostore, err := NewOrganizationStore(ctx, ds)
	if err != nil {
		t.Fatal(err)
	}

	e := &Organization{}
	_, err = ostore.Create(ctx, ostore.NameKey(ctx, "fuga"), e)
	if err != nil {
		t.Fatal(err)
	}

	l, err := ostore.ListAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if e, g := 1, len(l); e != g {
		t.Fatalf("expected list.len = %v, got %v", e, g)
	}
	fmt.Printf("key is %+v\n", l[0].Key)
}
