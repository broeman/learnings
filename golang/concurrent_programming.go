/* Concurrent Programming in Go
* notes continued from effective programing in Go
 */

/* Gorutines:
* not a thread, coroutine or process, it is more simple
* a function executing concurrently with other goroutines in the same address space
* cost little more than alloc of stack
* grows by alloc heap storage if required
* multiplexed onto multiple OS threads, if one waits, others continue to run
* hiding the complexity of thread creation and management
 */

go list.Sort() // run list.Sort concurrently; don't wait for it

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // note, must call the function
}

/* Channels:
* signaling
* alloc with make
*/
ci := make(chan int)	// unbuffered channel of integers
cj := make(chan int, 0)
cs := make(chan *os.File, 100)	// buffered channel of pointers to Files

// Example 1:
c := make(chan int)
go func() {
	list.Sort()
	c <- 1		// send a signal; value doesn't matter	
}()
doSomethingForAWhile()
<-c 			// wait for sort to finish; discard sent value

// Example 2:
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
	sem <- 1	// wait for active queue to drain
	process(r)	// may take a long time
	<-sem		// done; enable request to run
}

func Serve(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}

// Example 3:
func handle(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit
}

// channels of channels
type Request struct {
	args		[]int
	f			func([]int) int
	resultChan 	chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
clientRequests <- request	// send request
fmt.Printf("answer: %d\n", request.resultChan)	// wait for response

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

// parallelization
type Vector []float64
const NCPU = 4	// number of CPU cores

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)
	for i:=0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// draining the channel
	for i:=0; i < NCPU; i++ {
		<-c // wait for one task to comple
	}
	// all done at this point
}

// Leaky buffer
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		// grab a buffer if available, allocate if not
		select {
		case b = <-freeList:
			// got one, nothing more to do
		default:
			// none free, so allocate a new one
			b = new(Buffer)
		}
		load(b)				// read next message from the net
		serverChan <- b		// send to server
	}
}

func server() {
	for {
		b := <-serverChan	// wait for work
		process(b)
		// reuse buffer if there's room
		select {
		case freeList <- b:
			// buffer is on free list, nothing to do
		default:
			// free list full, carry on
		}
	}
}

/* Errors:
* use multivalue returns for errors
* kind of a "comma, ok" idiom
*/
type error interface {
	Error() string
}

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

/* Panic:
* unrecoverable errors, where the program cannot continue
*/
var user = os.Getenv("USER")

func init() {
	if user = "" {
		panic("no value for $USER")
	}
}

/* Recover:
* stop the unwinding of panic, and returns the argument passed to panic
*/

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

// continued in webservereffective.go