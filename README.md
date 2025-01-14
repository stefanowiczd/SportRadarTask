# Football World Cup Score Board

## Assumptions
- fixed number of teams (32) takes place in the tournament
- to fake some actions there is applied pseudo logic of 5 actions for the played games
  - goal scored by home team
  - goal scored by away team 
  - no goal (gives a possibility to end game with score 0-0 or other draw result, like 1-1, 2-2, etc.)

## How to
- run program
```makefile
make run
```
- run unit tests
```makefile
make test
```
- run test to verify race condition hazard
```makefile
make test-race-cond
```