module example/hello

go 1.19

require (
	example/mascot v0.0.0-00010101000000-000000000000
	golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
)

replace example.com/greetings => ../greetings

replace example/mascot => ../mascot
