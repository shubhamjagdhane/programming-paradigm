# My Log Framework - mylog
```
This project is build using the built-in 'log' package of Golang. 
I used that library and implemented my six functions which are Debug, Info, Warn, Error, Fatalf, Panicf.

The implementation has been made with respect to understanding the every detail of built-in package 'log'. 
The motive of this project is to understand use case of Singleton pattern.
```

### Getting Started
1. Clone the repository
2. Import the module in your code
3. Use the above mentioned functions
4. You can used go mod as well to get the mylog framework

### Newly added functions:

#### func Debug(format, args...)
```
  This is function is implemented the Debugging logs.
  You can call this function as mylog.Debug("anything", number_of_arguments)
``` 

#### func Info(format, args...)
```
  This is function is implemented the informative logs.
  You can call this function as mylog.Info("anything", number_of_arguments)
``` 

#### func Warn(format, args...)
```
  This is function is implemented the warning logs.
  You can call this function as mylog.Warn("anything", number_of_arguments)
```  

#### func Error(format, args...)
```
  This is function is implemented the error logs.
  You can call this function as mylog.Error("anything", number_of_arguments)
```

### Modified functions:

#### func Fatalf(format, args...)
```
This is function is implemented the fatal logs. This function is similar as the built-in Fatalf function modified with the some flags and easily reading message.
You can call this function as mylog.Fatalf("anything", number_of_arguments)
```
#### func Panicf(format, args...)
```
This is function is implemented the panic logs. This function is similar as the built-in Panicf function modified with the some flags and easily reading message.
You can call this function as mylog.Fatalf("anything", number_of_arguments)  
```

### Prerequisites:

```
Golang
```

### Installation:
1) Golang: To install Golang [Click Here](https://golang.org/dl/)

## Built With Stack

* [Golang](https://golang.org/) - The language is used for implementation
* [log](https://golang.org/pkg/log/) - The logging package

## Versioning
Version 1.0


## Developer Help
If you stuck somewhere mail me to

```
me.shubhamjagdhane@gmail.com
```
