package workerpool_processtask

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	MessageBody string
	Subject     string
}

// Way to process the email task
func (e *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", e.Email)
	// Simulate a time-consuming process
	time.Sleep(time.Second * 5)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	Concurrency int // Number of workers that can run at a time
	TasksChan   chan Task
	Wg          sync.WaitGroup
}

// Function to execute the worker pool
func (wp *WorkerPool) Worker() {
	// Lops the channel to process the task
	for task := range wp.TasksChan {
		task.Process()
		wp.Wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.TasksChan = make(chan Task, len(wp.TasksChan))

	// Start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.Worker()
	}

	// Send tasks to the tasks channel
	wp.Wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.TasksChan <- task
	}
	close(wp.TasksChan)

	wp.Wg.Wait()
}

func GetDataTest() []Task {
	return []Task{
		&EmailTask{
			Email:       "a@gmail.com",
			MessageBody: "Message body A",
			Subject:     "test",
		},
		&EmailTask{
			Email:       "f@gmail.com",
			MessageBody: "Message body F",
			Subject:     "test",
		},
		&EmailTask{
			Email:       "b@gmail.com",
			MessageBody: "Message body B",
			Subject:     "test",
		},
		&EmailTask{
			Email:       "c@gmail.com",
			MessageBody: "Message body C",
			Subject:     "test",
		},
		&EmailTask{
			Email:       "d@gmail.com",
			MessageBody: "Message body D",
			Subject:     "test",
		},
		&EmailTask{
			Email:       "e@gmail.com",
			MessageBody: "Message body E",
			Subject:     "test",
		},
	}
}

func Main() {
	wp := WorkerPool{
		Tasks:       GetDataTest(),
		Concurrency: 2,
	}
	wp.Run()
	fmt.Println("All tasks have been run successfully")

}
