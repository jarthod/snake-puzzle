This is my first go project and I wrote it to learn go. Its purpose is to solve a 4x4 "snake cube" wooden puzzle, like this one:

![](https://i.ytimg.com/vi/iTzVPgFjE9c/maxresdefault.jpg)

```
> go run snake.go

 Start:
            z: 0        z: 1        z: 2        z: 3
 y: 3    -  -  -  -      -  -  -  -      -  -  -  -      -  -  -  -
 y: 2    -  -  -  -      -  -  -  -      -  -  -  -      -  -  -  -
 y: 1    -  -  4  -      -  -  -  -      -  -  -  -      -  -  -  -
 y: 0    1  2  3  -      -  -  -  -      -  -  -  -      -  -  -  -

 Solved: true
            z: 0        z: 1        z: 2        z: 3
 y: 3   22 23 26 27     37 24 25 28     38 49 50 51     39 48 63 64
 y: 2   21 20 15 14     36 19 16 29     35 18 17 52     40 47 62 53
 y: 1    6  5  4 13      7 32 31 30     34 33 56 55     41 46 61 54
 y: 0    1  2  3 12      8  9 10 11     43 44 57 58     42 45 60 59
 ```