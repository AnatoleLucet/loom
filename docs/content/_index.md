---
title: ""
toc: false
---

{{< rawhtml >}}

<div id="home-hero">
    <h1 class="not-prose">「#」</h1>
    <p>The reactive framework for Go.</p>
    <br/>
    <a href="/docs/getting-started">GET STARTED -></a>
</div>

{{< /rawhtml >}}

---

Loom is a lightweight UI framework used to build Terminal UIs and Web SPAs.

It lets you write surprisingly concise UI components in pure Go, that can be brought to life with reactive signals, and do expensive work in goroutines.

If you're coming from SolidJS or ReactJS, you'll feel right at home.

{{< rawhtml >}}
<br/><br/>
{{< /rawhtml >}}

#### A simple counter

```go {style=tokyonight-moon}
func Counter() Node {
	count, setCount := Signal(0)

    go func() {
        for {
            time.Sleep(time.Second)
            setCount(count() + 1)
        }
    }()

	return P(Text("Count: "), BindText(count))
}
```

{{< rawhtml >}}

<a href="/docs/examples" class="hx:text-sm">MORE EXAMPLES -></a>
<br/>
<br/>

{{< /rawhtml >}}

---

#### Features

[*] **Pure Go** | No extra compiler. Use the tools you already know.

[*] **Multi-plateform** | The only thing loom needs is a [renderer](/docs/concepts.md#renderer). Build a UI for your calculator if you want.

[*] **Signal-based** | Performant fine-grained reactive updates. You [decide](/docs/components/bind) what part of the tree gets to update.

[*] **Concurrency-safe** | Run goroutines from your components and update signals from them. It's built for that!

---

{{< rawhtml >}}

<div id="home-bottom-hero">
    <h3>GET STARTED</h3>
    <br/>
    <a href="/term/getting-started">LOOM-TERM -></a>
    <a href="/web/getting-started">LOOM-WEB -></a>
</div>

{{< /rawhtml >}}
