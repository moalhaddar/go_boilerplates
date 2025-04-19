# Context

Basically it's an immutable tree like datastructure.

There are different context types, but we usually deal with a single interface, Context. The Types are: 
- emptyCtx
- cancelCtx
- timeoutCtx
- valueCtx

cancelCtx and timeoutCtx maintain a list of children to propagagte the cancellations down the tree.


Context is used to pass values and propagate timeouts across the services.


