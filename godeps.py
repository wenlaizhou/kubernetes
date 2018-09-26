# coding=utf8
import sys

import itertools, pprint

reload(sys)
sys.setdefaultencoding('utf-8')

import os
import re
import utils

for d in utils.getGoDeps():
    if "k8s.io/kubernetes/" not in d:
        print d
