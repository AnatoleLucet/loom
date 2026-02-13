---
title: Self()
weight: 1
---

```go {style=tokyonight-moon}
func Self() Component
```

```go {style=tokyonight-moon}
type Component interface {
	Disposed() (done <-chan struct{})
    IsDisposed() (disposed bool)

	OnCleanup(fn func())

    Context() context.Context
}
```
