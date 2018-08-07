# coding=utf8
import sys

import itertools, pprint

reload(sys)
sys.setdefaultencoding('utf-8')

import os
import re

imports = {}

reg1 = re.compile('''import \((.*?)\)''', flags=re.MULTILINE + re.DOTALL)
reg2 = re.compile('''import "(.*?)"''')

for root, dirs, files in os.walk('.'):
    for f in files:
        if not f.endswith('.go'):
            continue

        fileName = "{}{}{}".format(root, os.path.sep, f)
        with open(fileName) as fs:
            fileContent = fs.read()
            imports1 = reg1.findall(fileContent)
            import2 = reg2.findall(fileContent)
            for i in itertools.chain(imports1, import2):
                i = i.strip()
                if not i:
                    continue
                if '"' in i:
                    words = i.split('"')
                    for w in words:
                        w = w.strip()
                        if w:
                            imports[w] = True
                else:
                    imports[i] = True


def exe(cmd):
    _, stdOut = os.popen2(cmd)
    res = stdOut.readlines()
    pid = os.getpid()
    return res, pid


for k in imports:
    if re.match("(\w+\.\w+)", k):
        res, _ = exe("go get -v {}".format(k))
        print(res)
