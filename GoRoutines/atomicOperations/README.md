### Atomic Operations

The other way to prevent race conditions is to use atomic transactions.

Atomicity is a guarantee of isolation from interrupts,
 signals, concurrent processes and threads. 
 It is relevant for thread safety and reentrancy.
  Additionally, atomic operations commonly have a 
  succeed-or-fail definition—they either successfully 
  change the state of the system, or have no relevant effect.

https://godoc.org/sync/atomic