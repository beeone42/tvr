#!/usr/bin/env python

import sys
import os

os.chdir("scripts")
os.system("chmod +x generating_conf.py fabfile.py")
os.system("./generating_conf.py" + ' ' + sys.argv[1])
os.system("fab  deploy:host=10.62.1.2")
