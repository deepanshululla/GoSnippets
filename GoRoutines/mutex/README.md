## Mutex(mutual exclusion objects)
In computer programming, a mutex (mutual exclusion object) 
is a program object that is created so that multiple program 
thread can take turns sharing the same resource,
 such as access to a file.
 
 The mutex locks just prevent one thread from accessing another.
 
#### Using Mutex:
 
 A mutex provides mutual exclusion, either producer or consumer
  can have the key (mutex) and proceed with their work. 
  As long as the buffer is filled by producer,
   the consumer needs to wait, and vice versa.
 
 At any point of time, only one thread can work
  with the entire buffer. The concept can be generalized
   using semaphore.
 
#### Using Semaphore:
 
 A semaphore is a generalized mutex. In lieu of single buffer,
  we can split the 4 KB buffer into four 1 KB buffers
   (identical resources).
  
 A semaphore can be associated with these four buffers. 
 The consumer and producer can work on different buffers
  at the same time.
 
#### Difference between Mutex and Semaphore
Strictly speaking, a mutex is locking mechanism
 used to synchronize access to a resource. 
Only one task (can be a thread or process based 
on OS abstraction) can acquire the mutex.
 It means there is ownership associated with mutex,
  and only the owner can release the lock (mutex).

Semaphore is signaling mechanism 
(“I am done, you can carry on” kind of signal). 
For example, if you are listening songs (assume it as one task)
 on your mobile and at the same time your friend calls you,
  an interrupt is triggered upon which an
   interrupt service routine (ISR) 
signals the call processing task to wakeup.

Binary Semaphore is not the same as Mutex.


```text
Also Note that just becuase you are using mutex doesn't mean you are avoiding
all possible race conditions. 
You can still have race conditions if you don't use mutex properly


```