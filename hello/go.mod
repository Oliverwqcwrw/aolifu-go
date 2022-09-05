module oliver.com/hello

go 1.18

replace oliver.com/greetings => ../greetings

require (
	github.com/google/go-cmp v0.5.8
	oliver.com/greetings v0.0.0-00010101000000-000000000000
)

require golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
