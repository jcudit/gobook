package my_safeslice_test

import (
	"fmt"
	"my_safeslice"
	"sync"
	"testing"
)

func TestSafeSlice(t *testing.T) {
	store := my_safeslice.New()
	fmt.Printf("Initially has %d items\n", store.Len())

	var waiter sync.WaitGroup

	waiter.Add(1)
	go func() { // Concurrent Inserter
		for i := 0; i < 100; i++ {
			store.Append(fmt.Sprintf("0x%04X", i))
			if i > 0 && i%15 == 0 {
				fmt.Printf("Inserted %d items\n", store.Len())
			}
		}
		fmt.Printf("Inserted %d items\n", store.Len())
		waiter.Done()
	}()

	waiter.Wait()

	waiter.Add(1)
	go func() { // Concurrent Deleter
		for i := 0; i < 15; i++ {
			before := store.Len()
			store.Delete(0)
			fmt.Printf("Deleted (%d) before=%d after=%d\n",
				i, before, store.Len())
		}
		waiter.Done()
	}()
	fmt.Printf("Now has %d items\n", store.Len())

	waiter.Wait()

	waiter.Add(1)
	go func() { // Concurrent Finder
		for i := 0; i < 15; i++ {
			value := store.At(i)
			fmt.Printf("Value at index %s is %d\n", i, value)
		}
		waiter.Done()
	}()

	waiter.Wait()

	updater := func(index int, value interface{}) interface{} {
		return value.(string) + "!"
	}
	for _, i := range []int{5, 10, 15, 20, 25, 30, 35} {
		if store.Len()+1 >= i {
			fmt.Printf("Original %s == %d\t", i, store.At(i))
			store.Update(i, updater)
			fmt.Printf("Updated %s == %5d\n", i, store.At(i))
		}
	}

	fmt.Printf("Finished with %d items\n", store.Len())
	// not needed here but useful if you want to free up the goroutine
	data := store.Close()
	fmt.Println("Closed")
	fmt.Printf("len == %d\n", len(data))
	for k, v := range data {
		fmt.Printf("%s = %v\n", k, v)
	}
}
