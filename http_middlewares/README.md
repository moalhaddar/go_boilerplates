# HTTP middlewares

This is one of those confusing moments when i discovered adapter types in golang.

The middleware idea is that we want to construct a master http handler that chains handlers together

This works through function composition, bottom up.

If we want a middlware that goes through a -> b -> final

Then we have to compose it like (a(b(final))), this is higher order functions pattern.

Such that b(final) is evaluated first, then injected into a

a middleware is an http handler that takes a "next' http handler as it's input

# The confusing part (for me)
An http handler can be constructed from a type.
This type is an adapter type type.

The type name is HandlerFunc. It takes a function and returns an http handler

```go
http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    // some logic
})
```

Now this is a handler, a middleware is a functin that takes a handler and returns a handler, so:

```go
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        // do stuff
        next.ServeHTTP(w, r)
        // do stuff
    })
}
```

Now this works, and we can even go one level deeper and compose the middleware under another functio to parametrize it:

```go
func LoggerMiddleware(options Options) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
            // do stuff
            next.ServeHTTP(w, r)
            // do stuff
        })
    }
}
```

