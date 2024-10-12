package workforce

import "context"

// The ExecuteFunction is the single unit of work to be performed
type ExecuteFunction func(ctx context.Context, args ...interface{}) (interface{}, error)

// A job has the fields required to submit a unit of work to the worker pool
type Job struct {
	Description string
	Arguments   interface{}
	executable  ExecuteFunction
}
