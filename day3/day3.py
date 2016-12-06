#!/usr/bin/env python

import sys
fLoc = sys.argv[1]

def sidesGreater(a, b, c):
    return int(a) + int(b) > int(c)

count = 0

with open(fLoc) as f:
    for line in f:
        side = line.split()
        if(sidesGreater(side[0], side[1], side[2]) and
            sidesGreater(side[1], side[2], side[0]) and
            sidesGreater(side[2], side[0], side[1])):
            count +=1
print(count)


