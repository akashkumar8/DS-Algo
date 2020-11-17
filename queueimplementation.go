package main
import "container/list"
import "fmt"

func main() {
    // new linked list
    queue := list.New()

    // Simply append to enqueue.
    queue.PushBack(40)
    queue.PushBack(50)
    queue.PushBack(60)

    // Dequeue
    front:=queue.Front()
    fmt.Println(front.Value)
    // This will remove the allocated memory and avoid memory leaks
    queue.Remove(front)
}