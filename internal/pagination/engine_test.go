package pagination

import (
	"context"
	"errors"
	"testing"
)

func TestEachIteratesAllPages(t *testing.T) {
	t.Helper()

	var got []int
	err := Each(context.Background(), "", func(_ context.Context, cursor string) ([]int, string, error) {
		switch cursor {
		case "":
			return []int{1, 2}, "p2", nil
		case "p2":
			return []int{3}, "", nil
		default:
			return nil, "", errors.New("unexpected cursor")
		}
	}, func(v int) error {
		got = append(got, v)
		return nil
	})
	if err != nil {
		t.Fatalf("Each returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("unexpected items: %#v", got)
	}
}

func TestEachStopsOnYieldError(t *testing.T) {
	t.Helper()

	stopErr := errors.New("stop")
	calls := 0
	err := Each(context.Background(), "", func(_ context.Context, _ string) ([]int, string, error) {
		calls++
		return []int{1, 2}, "next", nil
	}, func(v int) error {
		if v == 2 {
			return stopErr
		}
		return nil
	})
	if !errors.Is(err, stopErr) {
		t.Fatalf("expected stopErr, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected 1 fetch call, got %d", calls)
	}
}

func TestEachReturnsStalledError(t *testing.T) {
	t.Helper()

	err := Each(context.Background(), "same", func(_ context.Context, _ string) ([]int, string, error) {
		return []int{1}, "same", nil
	}, func(_ int) error {
		return nil
	})
	if !errors.Is(err, ErrPaginationStalled) {
		t.Fatalf("expected ErrPaginationStalled, got %v", err)
	}
}
