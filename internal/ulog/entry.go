package ulog

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type (
	contextKey = struct{}
)

type WithFieldFn func(context.Context, *logrus.Entry) *logrus.Entry

func WithField(field string, value interface{}) WithFieldFn {
	return func(_ context.Context, entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(field, value)
	}
}

func WithContextField(field string, key interface{}) WithFieldFn {
	return func(ctx context.Context, entry *logrus.Entry) *logrus.Entry {
		value := ctx.Value(key)
		return entry.WithField(field, value)
	}
}

func WithComponent(component interface{}) WithFieldFn {
	return WithField("component", fmt.Sprintf("%T", component))
}

func WithComponentName(name string) WithFieldFn {
	return WithField("component", name)
}

func GetContext(ctx context.Context, fns ...WithFieldFn) *logrus.Entry {
	entry := ctx.Value(contextKey{})
	if entry == nil {
		return nil
	}

	if entry, ok := entry.(*logrus.Entry); ok {
		for _, fn := range fns {
			entry = fn(ctx, entry)
		}
		return entry
	}

	return nil
}

func PutContext(ctx context.Context, entry *logrus.Entry, fns ...WithFieldFn) context.Context {
	for _, fn := range fns {
		entry = fn(ctx, entry)
	}
	return context.WithValue(ctx, contextKey{}, entry)
}
