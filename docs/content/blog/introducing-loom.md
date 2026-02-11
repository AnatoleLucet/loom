---
date: "2026-02-11T10:24:52+01:00"
title: "Introducing Loom"
description: "The component framework for Go."
---

---

{{< rawhtml >}}

<h1 align="center">「#」</h1>

{{< /rawhtml >}}

```go {style=tokyonight-moon}
func App() Node {
	count, setCount := Signal(0)
	increment := func(*Event) {
		setCount(count() + 1)
	}

	return Div(
		H2(Text("Count: "), BindText(count)),

		Button(
            Text("Increment"),
            On("click", increment),
        ),
	)
}
```
