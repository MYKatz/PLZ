<br/>
<div style="text-align:center">
    <img src="assets/logo.png">
    <h2>PLZ</h2>
    <h2>The world's politest programming language</h2>
</div>

## Summary

PLZ is an [esoteric programming language](https://en.wikipedia.org/wiki/Esoteric_programming_language) (read: joke language) interpreted in Go. Despite its humorous origins, it is turing-complete and relatively feature-full.

Here's an example of a PLZ program. This will return the 10th fibonacci number (counting from 0):

```
let fibonacci be function(x) please
    if(x==0) please
        return 0 plz
    thanks
    if(x==1) please
        return 1 plz
    thanks
    return fibonacci(x - 1) + fibonacci(x - 2) plz
thanks
```

The above code will return 55.

## Syntax

### Variable Assignment

Variables are assigned with the 'be' operator.

```
let name be "Matt" plz
let age be 18 plz
let isCool be False plz
```

### Data types

There are five currently supported data types in PLZ: Strings, booleans, integers, arrays, and hashtables (aka dictionaries, associative arrays).

#### Arrays

These are declared in a familiar way:

```
let fruits be ["apple", "banana", "pineapple", "kiwi"] plz
```

Values in an array can be accessed by their index:

```
fruits[1] //"banana"
```

Like in languages such as Python, arrays can contain data of different types.

```
let things be ["this is a string", "but we can also have ints", 1, 2, 3]
```

Arrays can also contain other arrays!

```
let arrs be [ [0,1], [2,3] ] plz arrs[0]
arrs[1] // [2,3]
arrs[0][1] // 1
```

### Hash Tables

Hash tables/dictionaries are declared and accessed as follows:

```
let ages be {"Mark": 35, "Sergey": 46, "Jeff": 55} plz
ages["Jeff"] //55
```

Accessing a key that has not been assigned will not throw an error, instead simply returning Null:

```
ages["Jimmy"] //None
```

Dictionary keys can be Strings, booleans, or integers. Dictionary values can be of any type:

```
let dict be {
    "num": 45,
    22: "Grover Cleveland",
    "firstfive": ["Washington", "Adams", "Jefferson", "Madison", "Monroe"]
} plz
```

## Standard Library

PLZ also comes with a basic standard library. The standard library is being expanded at the moment. The currently implemented functions are described here.

### print

The print function prints values to the console. It takes any amount of arguments greater or equal to one, of any data type.

```
print("hello, world!") plz //hello, world!
print("My", "name", "is", "Jonas") plz //"My \n name \n is \n Jonas \n"
print("Here is my todo list:", ["Clean room", "check emails", "go to the gym"])
```

### len

Returns the length of a string OR an array.

```
len("hello") //5
len(["one", "two", "three"]) //3
```

### assign

Reassigns a value in an array or hash table.

```
let resources be ["brick", "lumber", "grain", "ore", "sheep"] plz
assign(resources, 4, "wool") plz
resources[4] //"wool"
```

```
let student be {"name": "Matt", "major": "Econ"} plz
assign(student, "major", "CS") plz
student["major"] //"CS"
```

### append

Takes an existing array and a variable; returns a new array with the variable appended.

```
let pidigits be [3, 1, 4, 1] plz
let pidigits be append(pidigits, 5) plz
pidigits //[3, 1, 4, 1, 5]
```

### peek

Returns the last value of an array

```
let fruits be ["apple", "orange", "banana"] plz
peek(fruits) //"banana"
```

### first

Returns the first value of an array

```
let fruits be ["apple", "orange", "banana"] plz
first(fruits) //"apple"
```

### rest

Returns an array with the first element removed

```
let fruits be ["apple", "orange", "banana"] plz
rest(fruits) //["orange", "banana"]
```
