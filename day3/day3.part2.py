#!/usr/bin/env python

import sys
fLoc = sys.argv[1]

def sidesGreater(a, b, c):
    return int(a) + int(b) > int(c)

def isTriangle(side, n): 
    if (sidesGreater(side[0+n], side[3+n], side[6+n]) and
        sidesGreater(side[3+n], side[6+n], side[0+n]) and
        sidesGreater(side[6+n], side[0+n], side[3+n])):
        return True 

count = 0

with open(fLoc) as f:
    triangles = []
    lineCounter = 0
    for line in f:
        lineCounter += 1
        triangles.extend(line.split())
        # check sides if we have 3 lines
        if lineCounter == 3:
            for i in range(3):
                if isTriangle(triangles, i):
                    count += 1
            lineCounter = 0
            triangles = []

print(count)


