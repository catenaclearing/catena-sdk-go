package pagination

import (
"context"
"errors"
)

// PageFetcher is a function that fetches a page of items.
// It returns the items, the next cursor, and an error if any.
type PageFetcher[T any] func(ctx context.Context, cursor string) (items []T, next string, err error)

// ErrPaginationStalled is returned when the next page cursor is the same as the current cursor.
var ErrPaginationStalled = errors.New("pagination stalled: next_page equals current cursor")

// Each iterates over all items in a paginated resource.
// It calls the fetch function to get pages of items, and yields each item to the yield function.
// If the yield function returns an error, iteration stops and that error is returned.
func Each[T any](ctx context.Context, start string, fetch PageFetcher[T], yield func(T) error) error {
	cursor := start
	for {
		items, next, err := fetch(ctx, cursor)
		if err != nil {
			return err
		}

		for _, item := range items {
			if err := yield(item); err != nil {
				return err
			}
		}

		if next == "" {
			break
		}

		if next == cursor {
			return ErrPaginationStalled
		}

		cursor = next

		if err := ctx.Err(); err != nil {
			return err
		}
	}
	return nil
}
