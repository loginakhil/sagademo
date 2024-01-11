package main

import (
	"log"
	"sagademo/pkg/saga"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, saga.TransferMoneyTaskQueue, worker.Options{})

	// registers a workflow function with the worker
	w.RegisterWorkflow(saga.TransferMoney)

	w.RegisterActivity(saga.Deposit)
	w.RegisterActivity(saga.DepositCompensation)
	w.RegisterActivity(saga.StepWithError)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start deposit worker", err)
	}
}
