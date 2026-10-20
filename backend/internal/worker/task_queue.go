package worker

import (
    "fmt"
    "time"
)

type TaskQueue struct {
    jobs chan func()
}

func NewTaskQueue(size int) *TaskQueue {
    return &TaskQueue{jobs: make(chan func(), size)}
}

func (q *TaskQueue) Enqueue(task func()) {
    q.jobs <- task
}

func (q *TaskQueue) Worker() {
    for job := range q.jobs {
        job()
        time.Sleep(200 * time.Millisecond)
    }
}

func (q *TaskQueue) Submit(name string, fn func()) {
    q.Enqueue(func() {
        fmt.Println("processing:", name)
        fn()
    })
}
