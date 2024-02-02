package worker

import (
    "testing"
)

func TestGetWorkerTypeIsMapper(t *testing.T) {
    worker := MapWorker{}
    workerType := worker.GetWorkerType()

    if workerType != Mapper {
        t.Errorf("Expected type Mapper but got something else instead")
    }
}
