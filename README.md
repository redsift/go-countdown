# go-countdown
countdown implements a context that uses a countdown on time before timing out

The standard context package let's you create a context based on timeout but if you need the timeout to be based on a countdown (either counter based or time based) then this package is for you. You can manually control the countdown by calling Sub(t) on the countdown context. t is subtracted from the timeout and if it hits 0 then the context is cancelled.

Usage:-
```

ctdwn := NewCountdown(ctx, 2*time.Second)
...
ctdwn.Sub(1*time.Second)
...
ctdwn.Sub(1*time.Second) // context get's cancelled after this call

```