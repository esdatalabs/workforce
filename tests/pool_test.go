package workforce_test

import (
	"context"
	"fmt"
	"testing"
	"workforce"
)

const (
	numberOfJobs    = 2
	numberOfWorkers = 1
)

// Helper function to generate a slice of jobs
func jobs() []workforce.Job {

	//Make a new slice with size capped at number of jobs
	bulk := make([]workforce.Job, numberOfJobs)

	for i := 0; i < numberOfJobs; i++ {

		bulk[i] = workforce.NewJob(
			fmt.Sprintf("This is job %v, to be converted to int", i),
			fmt.Sprintf("%v", i),
			castToInt,
		)
	}

	return bulk
}

func Test_Pool(t *testing.T) {

	//Create the pool
	pool := workforce.NewPool(numberOfJobs)

	//Initialized the workers
	go pool.Run(context.TODO())

	//Bulk load and close the inbox
	pool.LoodInbox(jobs())

	for {

		select {
		//Process messages sent to the outbox
		case resp, ok := <-pool.Results():

			//We only care about Results, ignore everying else
			if !ok {
				continue
			}

			//Valid jobs should not result in errors during processing
			if resp.Err != nil {
				t.Error(resp.Err.Error())
			}

		//Exit processing if any message is received on the done channel
		case <-pool.Done:
			return
		}

	}
}
