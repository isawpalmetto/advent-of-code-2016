#!/usr/bin/env python
import sys
fLoc = sys.argv[1]
print(fLoc)
row1 = [0,0,0,0,0,0,0]
row2 = [0,0,0,1,0,0,0]
row3 = [0,0,2,3,4,0,0]
row4 = [0,5,6,7,8,9,0]
row5 = [0,0,'A','B','C',0,0]
row6 = [0,0,0,'D',0,0,0]
row7 = [0,0,0,0,0,0,0]
row1.extend(row2)
row1.extend(row3)
row1.extend(row4)
row1.extend(row5)
row1.extend(row6)
row1.extend(row7)
keypad = row1
cursor = keypad.index(5)

def move(direction, current):
    if(direction == 'U'):
        tmp = current - 7
    if(direction == 'D'):
        tmp = current + 7
    if(direction == 'L'):
        tmp = current - 1
    if(direction == 'R'):
        tmp = current + 1
    return current if keypad[tmp] == 0 else tmp

with open(fLoc) as f:
    code = []
    for line in f:
        for direction in line:
            if direction != '\n':
                cursor = move(direction, cursor)
        code.append(keypad[cursor])

print(code)


