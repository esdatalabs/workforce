package workforce_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"workforce"
)

var castToInt workforce.ExecuteFunction = func(ctx context.Context, thingToCast interface{}) (interface{}, error) {

	maybeAnInt, ok := thingToCast.(string)

	if !ok {
		return nil, errors.New("The provided variable cannot be cast to an int")
	}

	return strconv.Atoi(maybeAnInt)
}

func Test_JobProcessesSuccessfully(t *testing.T) {

	ctx := context.TODO()

	job1 := workforce.NewJob("String to integer casting", "3", castToInt)

	result := job1.Execute(ctx)

	if result.Value != 3 {
		t.Errorf("Expected %s to be cast to int successfully: %s", "3", result.Err.Error())
	}
}
