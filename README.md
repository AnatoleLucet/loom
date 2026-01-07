<h1 align="center">「#」</h1>

<p align="center">A reactive component framework for TUIs, the Web, and more.</p>

```go
func App() Node {
	count, setCount := Signal(0)
	increment := func(*Event) {
		setCount(count() + 1)
	}

	return Div(
		H2(Text("Count: "), BindText(count)),
		Button(Text("Increment"), On("click", increment)),
	)
}
```

## Features

- Signal-based reactivity, similar to SolidJS or Svelte (see [loom/signals](https://github.com/AnatoleLucet/loom/tree/main/core/signals))
- JSX-like components architecture
- Multiples renderers ([TUI](https://github.com/AnatoleLucet/loom/tree/main/renderers/term), [Web](https://github.com/AnatoleLucet/loom/tree/main/renderers/web), and more to come)
-
