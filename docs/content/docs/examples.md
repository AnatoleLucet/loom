---
title: "EXAMPLES"
weight: 10
---

#### A list of fruits

coucou c'est super tout ca dis donc

{{< tabs items="TERM,WEB" >}}

{{< tab >}}

```go {style=tokyonight-moon}
func FruitList() Node {
    fruits, setFruits := Signal([]string{"banana", "apple", "orange"})

    return Box(
        For(
            fruits,
            func(fruit Accessor[string], index Accessor[int]) Node {
                return P(BindText(fruit))
            },
        ),
    )
}
```

{{< /tab >}}

{{< tab >}}

```go {style=tokyonight-moon}
func FruitList() Node {
    fruits, setFruits := Signal([]string{"banana", "apple", "orange"})

    return Ul(
        For(
            fruits,
            func(fruit Accessor[string], index Accessor[int]) Node {
                return Li(BindText(fruit))
            },
        ),
    )
}
```

{{< /tab >}}

{{< /tabs >}}

{{< rawhtml >}}
<br/>
{{< /rawhtml >}}

#### Taking user input

{{< tabs items="TERM,WEB" >}}

{{< tab >}}

```go {style=tokyonight-moon}
func UserInput() Node {
    value, setValue := Signal("")

    update := func(e *EventInput) {
        setInput(e.InputValue())
    }

    return Box(
        P(Text("Value: "), BindText(value)),

        Input(On("input", udpate)),
    )
}
```

{{< /tab >}}

{{< tab >}}

```go {style=tokyonight-moon}
func UserInput() Node {
    value, setValue := Signal("")

    update := func(e *EventInput) {
        setInput(e.InputValue())
    }

    return Div(
        P(Text("Value: "), BindText(value)),

        Input(
            Attr("type", "text"),
            BindAttr("value", input),
            On("input", udpate),
        ),
    )
}
```

{{< /tab >}}

{{< /tabs >}}
