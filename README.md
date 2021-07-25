# _umpire

A naive and incomplete mediator example in Go.

## Motivation

I wanted to see what the feasibility of a mediator implementation with an API similar
to [mediatr](https://github.com/jbogard/MediatR) but written in pure Golang was. Written in a couple of hours with a couple of 🥃 on a Saturday night.

## Example usage

```go
// create a new request and handler, this request returns a result
type (
	createBookCommand struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	createBookCommandHandler struct {
		model bookModel
	}
)

// implement requestReturnHandler for createBookCommandHandler
func (h *createBookCommandHandler) Handle(request interface{}, ctx context.Context) (interface{}, error) {
	command := request.(createBookCommand)

	b := book{
		command.Title,
		command.Author,
	}

	return h.model.Create(&b)
}

// -----------------------

func main() {
	// create a new mediator
	umpire := newMediator()
	model := newMockModel()

	// register the request with its appropriate handler
	umpire.registerReturnHandler(createBookCommand{}, &createBookCommandHandler{
		model,
	})

	// tell the mediator to dispatch the given command and return the result (and error)
	book, err := umpire.dispatchReturn(createBookCommand{
		Title:  "Jake's Book",
		Author: "Jake Smith",
	})

	if err != nil {
		panic(err)
	}

	// print the created book
	fmt.Printf("Created book %+v", book)
}
```