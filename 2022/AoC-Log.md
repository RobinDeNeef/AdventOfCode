# AoC Log

This year I'm keeping a small AoC log with some notes on every day. This way I can keep track of things I learned, algorithms I used or didn't use and other stuff that might be handy to remember.

## Day 1.
Puzzle itself isn't too hard. Unfortunatly my Go skills are quite rusty and coming back from Python I forgot how much work it is to work with a typed language. Most of the functions I wrote were used to get the data in the right data types in the right format.

## Day 2
Please sanitize your eyes after checking this solution. I just brute forced it with an if-else structure. I'm sure there are more clever ways to solve this. 

## Day 3
Some string manipulation. Wasn't the hardest puzzle. I discovered I can use negative returns to return a failed function. `return -1`. This also works for any type. 

## Day 4
Started of the day wrong. I thought I could just fill up the ranges as strings and check if range1 contained range 2 and vice versa. eg `4567` contains `56` and it worked on the test results. But as soon as the ranges went to the double digits it broke. Range 11-13 suddenly contained range 1-2 (`111213` contained `12`). So I had to throw it out and went with logical operators.  