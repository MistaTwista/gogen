# Gogen

Helps you to generate code using OpenAI API.

```golang
package main

//go:generate gogen --prompt "write golang function that implements debounce algorithm where you get channel with inputs and return channel with outputs. We can parametrize function with duration."

func main() {
}
```

## TODO
- [ ] modularize generator to use with other generators
