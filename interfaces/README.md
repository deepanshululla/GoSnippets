### Interfaces
Interfaces make the code more flexible, scalable and it’s a way to achieve polymorphism in Golang. Instead of requiring a particular type, interfaces allow to specify that only some behaviour is needed.
 
 Behaviour is defined by set of methods:

```text
type I interface {
    f1(name string)
    f2(name string) (error, float32)
    f3() int64
}
```

No particular implementation is enforced. 

It’s enough that type defines methods with desired names and signatures (input and output parameters) to say it implements (satisfies) an interface:
```text
type T int64
func (T) f1(name string) {
    fmt.Println(name)
}
func (T) f2(name string) (error, float32) {
    return nil, 10.2
}
func (T) f3() int64 {
    return 10
}
```

There are two different types
```text
concrete type   | interface type
---------------------------------
map struct int  |  bot
englishBot      |


The main difference is we can create a variable directly from
the concrete types but we can't directly create a variable from
interface type.

Concrete types are types that come form language like slices,maps
and also derived from the native types


```


Other Important notes

1) Interfaces do not provide support for generic types. In fact go
doesn't support generic types.

2) Interfaces are implicit. We don't need to manually specify 
the relation between the interface type and custom types that implement
them.

3) Interfaces are a contract to help us manage types.  if our custom types'
implementation of a function is broken then interfaces won't help us.
