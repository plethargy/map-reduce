package worker

import (
    "testing"
)

func TestGetWorkerTypeIsReducer(t *testing.T) {
    worker := ReduceWorker{}
    workerType := worker.GetWorkerType()

    if workerType != Reducer {
        t.Errorf("Expected type Mapper but got something else instead")
    }
}
