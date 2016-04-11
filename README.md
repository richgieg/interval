# interval


*See the 'example/example.go' for more details regarding the usage.*

--
    import "github.com/richgieg/interval"

Package interval provides the Interval type which allows repeated execution at a
fixed time interval.

## Usage

#### type Interval

```go
type Interval struct {
}
```

An Interval repeatedly calls a user-defined function at a fixed time interval.

#### func  NewInterval

```go
func NewInterval(d time.Duration, f func()) *Interval
```
NewInterval returns a new Interval which, once started, repeatedly calls the
function argument at the time interval specified by the duration arument.

#### func (*Interval) Start

```go
func (i *Interval) Start()
```
Start activates the Interval.

#### func (*Interval) Stop

```go
func (i *Interval) Stop()
```
Stop deactivates the Interval.
