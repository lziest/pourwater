#PourWater
This is a general solver for water pouring problem with 2 jugs of
capacity of X and Y liters respectively. We assume there is X+Y liters
of water in the third jug which has a infinite capacity.

It will solve for getting N liter of water by pouring water between jugs,
it will also optimize the solution. It even attempt to solve the complement
problem, which is getting X+Y-N liter of water, if that requires smaller
number of steps. 

## Build
```
go build .
```

## Install
```
go install .
```
## Run example
```
$./pourwater -x 7 -y 5 -n 4
Minimal number of pouring 6
Step 1 :
Pouring water to glass of 7-liter
Situation: 7, 0, 5
Step 2 :
Pouring water from glass of 7-liter to glass of 5-liter
Situation: 2, 5, 5
Step 3 :
Pouring water from glass of 5-liter to reservior
Situation: 2, 0, 10
Step 4 :
Pouring water from glass of 7-liter to glass of 5-liter
Situation: 0, 2, 10
Step 5 :
Pouring water to glass of 7-liter
Situation: 7, 2, 3
Step 6 :
Pouring water from glass of 7-liter to glass of 5-liter
Situation: 4, 5, 3
```
## More examples
```
$./pourwater -x 147 -y 245 -n 49
Minimal number of pouring 4
Step 1 :
Pouring water to glass of 147-liter
Situation: 147, 0, 245
Step 2 :
Pouring water from glass of 147-liter to glass of 245-liter
Situation: 0, 147, 245
Step 3 :
Pouring water to glass of 147-liter
Situation: 147, 147, 98
Step 4 :
Pouring water from glass of 147-liter to glass of 245-liter
Situation: 49, 245, 98
```

```
$./pourwater -x 20 -y 10 -n 5
It is impossible
```
