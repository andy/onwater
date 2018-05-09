package onwater

import (
	"context"
	"testing"
)

func TestCheckOnWater(t *testing.T) {
	if ok, err := New("").OnWater(context.Background(), 55.714733, 37.544347); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Error("should be on water")
	}
}

func TestCheckOnLand(t *testing.T) {
	if ok, err := New("").OnLand(context.Background(), 55.753675, 37.621339); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Error("should be on land")
	}
}
