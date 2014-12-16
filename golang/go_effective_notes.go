/*
	Long explanation of the implementation
	comment style
*/
package main

// using comment on a name, can be used in godoc
type T struct {
	name  string // name of the object
	value int    // its value
}

// grouping variables can indicate relationships, like a mutex here:
var (
	countLock   sync.Mutex
	inputCount  uint32
	outputCount uint32
	errorCount  uint32
)

// only use brackets right afterwards, the lexer cannot read other styles
if i < f() {
	g()
}

// simple if statement
if x > 0 {
	return y
}

// if accepts init statements, useful for local variables
if err := file.Chmod(0664); err != nil {
	log.Print(err)
	return err
}

// else is omitted, because errors typically ends in returns
f, err := os.Open(name)
if err != nil {
	return err
}
codeUsing(f)

// c-like for loop
for init; condition; post {}

// c-like while loop
for condition {}

// c-like for(;;) loop
for {}

// index variable declared in loop
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}

// looping over array, slice, string or map
for key, value := range oldMap {
	newMap[key] = value
}

// if you only need keys
for key := range m {
	if key.expired() {
		delete(m, key)
	}
}

// if you only need the value, use underscore (blank identifier)
sum := 0
for _, value := range array {
	sum += value
}

// reverse a using parallel assignments
for i, j := 0, len(a)-1; i < j; i, j = i + 1, j - 1 {
	a[i], a[j] = a[j], a[i]
}

// switch
func unhex(c byte) byte {
	// if no switch expression is set, it switches on true cases
	switch {
		case '0' <= c && c <= '9':
			return c - '0'
		case 'a' <= c && c <= 'f':
			return c - 'a' + 10
		case 'A' <= c && c <= 'F':
			return c - 'A' + 10
	}
	return 0
}

// switch has no auto fall through, but multiple cases can be represented commawise
func shouldEscape(c byte) bool {
	switch c {
		case ' ', '?', '&', '=', '#', '+', '%':
			return true		
	}
	return false
}

// break and continue can be used to break labels
LabelName:
	if this {
		break LabelName
	}

// a switch can be used to discover the type
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
	default:
		fmt.Printf("Unexpected type %T", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
}

// multiple return types
func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x * 10 + int(b[i]) - '0' // int is casted on b
	}
	return x, i
}

for i := 0; i < len(b); {
	x, i = nextInt(b, i)
	fmt.Println(x)
}

// named return values "result parameters"
func ReadFull(r Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
}

// defer: schedules a call to run after the function is finished executing

// Data
type SyncedBuffer struct {
	lock sync.Mutex
	buffer bytes.Buffer
}

p := new(SyncedBuffer) // type *SyncedBuffer
var v SyncedBuffer     // type SyncedBuffer

// Constructors
func NewFile(fd int, name, string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

// using a composite literal instead to simplify:
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0}
	return &f
}

// even more simpler
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	//return &File{fd, name, nil, 0}
	return &File{fd: fd, name: name} // saves even the null values
}

// allocation with make
make([]int, 10, 100) // array with 100 ints, sliced with 10 elements

var p *[]int = new([]int)	// allocates slice structure, *p == nil
var v []int = make([]int, 100)	// the slice v refers to a new array of 100 ints

var p *[]int = new([]int)	// Don't do
*p = make([]int, 100, 100)	// this

v := make([]int, 100) // ideomatic

// arrays is copied in go, so no need to iterate over them like in c,
// this example is too expensive:
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)

// use slices instead, which is wrappers for arrays
// this example:
n, err := f.Read(buf[0:32])

// is much more efficient and readable than this:
var n int
var err error
for i := 0; i < 32; i++ {
	nbytes, e := f.Read(buf[i:i+1])
	if nbytes == 0 || e != nil {
		err = e
		break
	}
	n += nbytes
}

// appending data to a slice
// len(slice) => length
// cap(slice) => capacity
func Append(slice, data[]byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		// allocate double for future growth
		newSlice := make([]byte, (l+len(data))*2)
		// copy function is predeclared
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

// two-dimensional arrays
type Transform [3][3]float64;	// 3 x 3 array
type LinesOfText [][]byte		// a slice of byte slices

// inner slices can have different length
text := LinesOfText{
	[]byte("Now is the time"),
	[]byte("for all good gophers"),
	[]byte("to bring some fun to the party.")
}

// 2D-slice: one line at the time (if they grow or shrink):
picture := make([][]uint8, YSize)		// Ys
for i := range picture {
	picture[i] = make([]uint8, XSize)	// Each Xs in Ys
}

// 2D-slice: one alloc with slices made into lines (if they are static):
picture := make([][]uint8, YSize)		// Ys
pixels := make([]uint8, XSize * YSize)
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}

// maps: dictionary
var timeZone = map[string]int{
	"UTC": 0*60*60,
	"EST": -5*60*60
}

offset := timeZone["EST"]

// fetching a key that doesn't exist returns a zero, and can be used as boolean
attended := map[string]bool{
	"Ann": true,
	"Joe": true
}

if attended[person] {
	fmt.Println(person, "was at the meeting")
}

// discriminate with using multiple assigments
var seconds int
var ok bool
seconds, ok = timeZone[tz] // the "comma ok" ideom

func offset(tz string) int {	
	if seconds, ok := timeZone[tz]; ok {	// if ok is true
		return seconds
	}
	// else
	log.Println("unknown time zone:", tz)
	return 0
}

// deleting an item
delete(timeZone, "EST")

// printing from the "fmt" package
// each of these print the same:
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n") // first argument is to implement io.Writer
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))

// %d: decimal, %x: hex

// %v => value:
fmt.Printf("%v\n", timeZone) // or fmt.Println(timeZone)
//output: map[EST:-18000 UTC:0]

// %T => type:
fmt.Printf("%T\n", timeZone)
//output: map[string] int

type MyString string
func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", string(m)) // needs to be converted!!
}

// arbitrary number of parameters
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}

// ...int => list of integers
func Min(a ...int) int {
	min := int(^uint(0) >> 1) 	// largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// now we have everything to build append:
func append(slice []T, elements ...T) []T {

}

// constants
const (
	_			= iota
	KB ByteSize = 1 << (10 * iota)
)

// variables
var (
	home = os.Getenv("HOME")
	user = os.Getenv("USER")	
)

// init for each source file.
// Common use is to verify or repair corretness of the program state
// before real execution begins
func init() {
	if user == "" {
		log.Fatal("$USER not set")	
	}
	if home == "" {
		home = "/home" + user
	}	
}

//// Methods:
type ByteSlice []byte

// by defining type in front of the function, it becomes a method of that type
func (slice ByteSlice) Append(data []byte) []byte {
	// ...
	return newSlice // needs to return an updated slice
}

// using a pointer makes the transition direct
func (p *ByteSlice) Append(date []byte) {
	slice := *p
	// ...
	*p = slice
}

// redone to comply with io.Writer:
func (p *ByteSlice) Write(date []byte) (n int, err error) {
	slice := *p
	// ...
	*p = slice // using a pointer makes the transition direct
	return len(data), nil
}

//// Interfaces:
type Sequence []int

// implementing the sort.Interface:
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing
func (s Sequence) String() string {
	sort.Sort(s)	// sorted by those three rules defined above
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	retrun str + "]"
}

// Conversions
func (s Sequence) String() string {
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

// refactor all this into this idiom:
type Sequence []int

func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Spring([]int(s))
}

// interface conversions
// this example mixes types:
type Stringer interface {
	String() string
}

var value interface{} // Value provided by caller
switch str := value.(type) {
case string:
	return str
case Stringer:
	return str.String()
}

// Generality, methods
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

import "net/http"
// ...
ctr := new(Counter)
http.Handle("/counter", ctr)

// Channeling
func Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// struct or interface embedding (since no subclasses is allowed)
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter is an interface that combines those above
// It is an union of both, so it can do both
type ReadWriter interface {
	Reader
	Writer
}

// Readwriter stores pointers to Reader and Writer
type ReadWriter struct {
	reader *Reader
	writer *Writer
}

// we still need to implement those interfaces needed:
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}
// ...

// convenient embedding:
// Job type has now access to methods of Logger
type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
job.Log("starting now")

func (job *Job) Logf(format string, args ...interface{}) {
	job.Logger.Logf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

// continued in concurrent_programming.go ...