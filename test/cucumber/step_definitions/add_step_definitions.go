package step_definitions

import (
	"context"
	"errors"

	"github.com/denniskreussel/scm/internal/helpers"
)

type addCtxKey struct{}

func IAddUpTheNumbersAnd(ctx context.Context, a, b int) (context.Context, error) {
	ctx = context.WithValue(ctx, addCtxKey{}, helpers.Add(a, b))
	return ctx, nil
}

func IShouldGet(ctx context.Context, arg1 int) error {
	result, ok := ctx.Value(addCtxKey{}).(int)
	if !ok {
		return errors.New("could not retrieve the result")
	}
	if result != arg1 {
		return errors.New("result does not match")
	}
	return nil
}
