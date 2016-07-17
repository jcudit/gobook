package my_safeslice

// import "fmt"

type safeSlice chan commandData

type SafeSlice interface {
	Append(interface{})     // Append the given item to the slice
	At(int) interface{}     // Return the item at the given index position
	Close() []interface{}   // Close the channel and return the slice
	Delete(int)             // Delete the item at the given index position
	Len() int               // Return the number of items in the slice
	Update(int, UpdateFunc) // Update the item at the given index position
}

type commandData struct {
	action  commandAction
	index   int
	value   interface{}
	result  chan<- interface{}
	data    chan<- []interface{}
	updater UpdateFunc
}

type commandAction int

const (
	remove commandAction = iota
	end
	insert
	at
	length
	update
)

type UpdateFunc func(int, interface{}) interface{}

func New() SafeSlice {
	ss := make(safeSlice)
	go ss.run()
	return ss
}

func (ss safeSlice) run() {
	store := make([]interface{}, 10)
	for command := range ss {
		switch command.action {
		case insert:
			store = append(store, command.value)
		case remove:
			tmp := append(store[:command.index], store[command.index+1:]...)
			store = tmp
		case at:
			if len(store) >= command.index+1 {
				command.result <- store[command.index]
			} else {
				command.result <- nil
			}
		case length:
			command.result <- len(store)
		case update:
			if len(store) >= command.index+1 {
				value := store[command.index]
				store[command.index] = command.updater(command.index, value)
			}
		case end:
			close(ss)
			command.data <- store
		}
	}
}

func (ss safeSlice) Append(value interface{}) {
	ss <- commandData{action: insert, value: value}
}

func (ss safeSlice) Delete(index int) {
	ss <- commandData{action: remove, index: index}
}

func (ss safeSlice) At(index int) (value interface{}) {
	reply := make(chan interface{})
	ss <- commandData{action: at, index: index, result: reply}
	return (<-reply).(interface{})
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeSlice) Update(index int, updater UpdateFunc) {
	ss <- commandData{action: update, index: index, updater: updater}
}

func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: end, data: reply}
	return <-reply
}
