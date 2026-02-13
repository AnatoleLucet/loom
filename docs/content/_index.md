---
title: ""
toc: false
---

<div id="home-hero">
    <h1 class="not-prose">「#」</h1>
    <p>The reactive framework for Go.</p>
    <br/>
    <a href="/docs/intro">GET STARTED -></a>
</div>

---

Loom is a lightweight UI framework used to build Terminal UIs and Web SPAs.

It lets you write surprisingly concise UI components in pure Go, that can be brought to life with reactive signals, and do expensive work in goroutines.

If you're coming from SolidJS or ReactJS, you'll feel right at home.

<br/>

#### A simple counter

```go {style=tokyonight-moon}
func Counter() Node {
	count, setCount := Signal(0)

    go func() {
        for {
            time.Sleep(time.Second / 30)
            setCount(count() + 1)
        }
    }()

	return P(Text("Count: "), BindText(count))
}
```

<video src="/medias/counter-term.mp4" autoplay loop muted playsinline></video>
<a href="/docs/examples" class="hx:text-sm">MORE EXAMPLES -></a>
<br />

---

#### Features

[*] **Pure Go** | No extra compiler. Use the tools you already know.

[*] **Multi-plateform** | The only thing loom needs is a [renderer](/docs/concepts.md#renderer). Build a UI for your calculator if you want.

[*] **Signal-based** | Performant fine-grained reactive updates. You [decide](/docs/guides/binding) what part of the tree gets to update.

[*] **Concurrency-safe** | Run goroutines from your components and update signals from them. It's built for that!

---

<div id="home-bottom-hero">
    <h3>GET STARTED</h3>
    <br/>
    <p align="center">
        <a href="/docs">LOOM -></a>
        <br/>
        The core framework.
        <br/>
        -------------------
        <br/><br/><br/>
        <a href="/term/getting-started">LOOM-TERM -></a>
        <br/>
        For powerful Terminal UIs.
        <br/>
        --------------------------
        <br/><br/><br/>
        <a href="/web/getting-started">LOOM-WEB -></a>
        <br/>
        For simple Web SPAs.
        <br/>
        --------------------
    </p>
</div>
