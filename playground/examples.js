examples = {
  helloWorld: `let msg be "Hello, world!" plz
print(msg) plz`,
  fibonacci: `let fibonacci be function(x) please
    if(x==0) please
        return 0 plz
    thanks
    if(x==1) please
        return 1 plz
    thanks
    return fibonacci(x - 1) + fibonacci(x - 2) plz
thanks

print(fibonacci(10))`,
  sum: `let nums be [1, 2, 3, 4, 5]

let sum be function(arr) please
    if(len(arr) == 1) please
        return arr[0] plz
    thanks
        return arr[0] + sum(rest(arr)) plz 
    thanks
    
print(sum(nums)) plz`
};
