## Jottings Of My Random Findings

- In Go, you run packages (folders) not file
- Go is alot like C++ in terms of main function and all

***Functions syntax***
```
func name(param1 param1Type, ...paramN paramNType) returnType { body }
```

- Command line flags, day-in:day-out over enviroment variables for configuration
- As a rule of thumb, you should avoid using the `Panic()` and `Fatal()` variations outside ofyour `main()` function — it’s good practice to return errors instead, and only panic or exit
directly from `main()`.

- Database connections should alwys be instantiated in the main function (longest living function) because the connection object is a pool of many connections underneath and are intended to be long-lived

- Interfaces are similar to polymorphism in oop, where derived classes can override the methods of their parent classes. but in this case, the derived class (type of that interface) must implement its own version of the interface type, even if it's the same as others