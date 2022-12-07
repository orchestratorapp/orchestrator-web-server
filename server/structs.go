package server

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

// The Router handler
type Router struct {
	rules      map[string]map[string]http.HandlerFunc
	JobQueue   chan Job
	Dispatcher *Dispatcher
}

// Create a Router instance
func BuildRouter(maxQueueSize int, maxWorkers int) *Router {
	router := Router{
		rules:    make(map[string]map[string]http.HandlerFunc),
		JobQueue: make(chan Job, maxQueueSize),
	}
	router.Dispatcher = BuildDispatcher(router.JobQueue, maxWorkers)
	return &router
}

// The ErrorResponse struct
type ErrorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}

func BuildErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: err.Error(),
	}
}

type Job struct {
	Name    uuid.UUID
	Writer  http.ResponseWriter
	Request *http.Request
	Handler http.HandlerFunc
}

type Worker struct {
	Id         uuid.UUID
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

func BuildWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         uuid.New(),
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Handler(job.Writer, job.Request)
			case <-w.QuitChan:
				log.Printf("\033[42m INFO \033[0m | Worker with ID %v stopped.",
					w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func BuildDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: workerPool,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := BuildWorker(d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}
