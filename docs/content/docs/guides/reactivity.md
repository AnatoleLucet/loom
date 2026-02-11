---
title: "REACTIVITY"
---

An introduction to reactivity and signals.

---

> If you're coming from a signal-based JavaScript framework like SolidJS or Svelte, you will rapidly notice some divergence in the reactive model. But you'll get used to it very soon, I promise.

### Signals

Signals are the core primitive to reactivity.

Like a variable, a signal holds a value. This value can be read, updated or kept as is. But unlike a variable, signals can be subscribed to. When its value changes, the subscribers get notified.

Let's look at an example:

```go {style=tokyonight-moon}
// define a signal called "count" with a default value of 0
count, setCount := Signal(0)

// `count` can read the value
count() // 0

// `setCount` can update it
setCount(1)
count() // 1


// example: increment by 1
setCount(count() + 1)
```

From here, there's no reason to use a signal over a regular variable. To see where signals before usefull, we must dive into subscribers.

### Memo

A Memo is one of multiple type of subscriber. It listens to one or more signals, and gets notified when there's a change.

It is responsible for holding a value, and recomputing that value when one of its dependencies (the signals it depends one) changes.

```go {style=tokyonight-moon}
// define the same signal as before
count, setCount := Signal(0)

// but this time declare a memo
double := Memo(func() int {
    // becaues we're readding `count` here,
    // this function will be called each time `count` changes.
    // this is the subscriber, and `count` is the dependency

    // read `count` and double its value
    return count() * 2
})

count() // 0
double() // 0

// update count (notifying the memo)
setCount(2)

count() // 2
double() // 4, the memo's value got updated!
```
