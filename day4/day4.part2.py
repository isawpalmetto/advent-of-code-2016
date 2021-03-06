#!/usr/bin/env python
import re,sys
from collections import defaultdict

fLoc = sys.argv[1]
sectorIDCount = 0
letters = {'a': 0, }
with open(fLoc) as f:
    for input in f:
        code = re.search(r'([\-a-zA-Z]*)(\d*)(\[[a-zA-Z]*\])', input)
        # remove dashes
        name = code.group(1)
        sectorID = code.group(2)
        # remove square branckets
        checksum = code.group(3)
        originalName = name
        name = name.replace('-', '')
        checksum = checksum.strip('[]')

        d = defaultdict(int)
        for k in name:
            d[k] += 1

        # sort by count highest
        counted = sorted(d.items(), key=lambda (k,v): v)
        length = max(key for (item, key) in counted)
        ordered = [[] for i in range(length+1)]
        for item, key in counted:
            ordered[key].append(item)
        testChecksum = []
        ordered.reverse()
        for l in ordered:
            alphaSort = sorted(l)
            for letter in alphaSort:
                if len(testChecksum) < 5:
                    testChecksum.append(letter)

        testChecksum =  ''.join(testChecksum)
        if testChecksum == checksum:
            originalName = originalName.replace('-', ' ')
            encrypted = []
            for c in originalName:
                if c != ' ':
                    encrypted.append(chr(((ord(c) - 97 + int(sectorID)) % 26)+97))
                else:
                    encrypted.append(c)
            encryptedName = ''.join(encrypted).strip()
            if encryptedName == 'northpole object storage':
                print sectorID


            
