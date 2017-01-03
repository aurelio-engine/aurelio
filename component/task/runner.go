package task

import "sync"

// Runner runs runnable tasks in parallel
type Runner struct {
	cap   int
	tasks []Runnable
	mutex sync.Mutex
}

// NewRunner returns new instance of task runner
func NewRunner(cap int) (runner Runner) {
	runner.cap = cap
	return
}

// AddTasks adds task for future execution
func (r *Runner) AddTask(task Runnable) {
	r.mutex.Lock()
	r.tasks = append(r.tasks, task)
	r.mutex.Unlock()
}

func (r *Runner) Run() {
	r.mutex.Lock()
	ch := make(chan Runnable)
	wg := new(sync.WaitGroup)
	wg.Add(r.cap)

	for i := 0; i < r.cap; i++ {
		go r.process(ch, wg)
	}

	for _, runnable := range r.tasks {
		ch <- runnable
	}

	close(ch)
	wg.Wait()
	r.cleanup()
}

func (r *Runner) process(ch chan Runnable, wg *sync.WaitGroup) {
	for runnable := range ch {
		runnable.Run()
	}
	wg.Done()
}

func (r *Runner) cleanup() {
	r.tasks = []Runnable{}
	r.mutex.Unlock()
}
