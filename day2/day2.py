#!/usr/bin/env python
import sys
fLoc = sys.argv[1]
print(fLoc)

keypad = range(1,10)
cursor = 4

def move(direction, current):
    if(direction == 'U'):
       tmp = current - 3
       return current if tmp < 0 else tmp
    if(direction == 'D'):
        tmp = current + 3
        return current if tmp > 8 else tmp
    if(direction == 'L'):
        return current if current in [0,3,6] else current - 1
    if(direction == 'R'):
        return current if current in [2,5,8] else current + 1


with open(fLoc) as f:
    code = []
    for line in f:
        for direction in line:
            if direction != '\n':
                cursor = move(direction, cursor)
        code.append(keypad[cursor])

print(code)


