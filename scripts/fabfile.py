#!/usr/bin/env python
from __future__ import with_statement
from fabric.api import *
from fabric.contrib.console import confirm
from fabric.contrib.files import exists
import os
import sys
import json

# gauche = 10.62.1.2
# droite = 10.62.1.3

env.hosts = ['10.62.1.2', '10.62.1.3']
env.user = "root"
env.password = "openelec"

def check():
    print env.hosts[0]
    print "erasing old videos"

def listing():
    run('ls')
    run('rm -rf videos')
    run('mkdir videos')

def transfer():
    tmp = "/root/go/src/github.com/beeone42/tvr/tmp/all.m3u"
    f = open(tmp)
    content =  f.readlines()
    content.pop(0)
    os.chdir("../video")
    for vid in content:
        os.system('ls')
        vid_sub = vid.split('/')
        vid_sub.pop(0)
        vid_sub.pop(0)
        vid_sub.pop(0)
        test = ''.join(vid_sub)
        test = test.strip()
        put(test, '/storage/videos')

    os.chdir("../tmp")
    if os.path.isfile('all.m3u'):
        put('all.m3u', '/storage/videos')
    else:
        print 'no all.m3u detected'
'''
    with quiet():
        os.chdir("../video")
        os.system('pwd')
        os.system('ls')
        if os.path.isfile(glob.glob('*.avi')):
            put('*avi', '/storage/videos')
        else:
            print 'no avi detected'
        os.system('pwd')
        if os.path.isfile('*.AVI'):
'''            

def restart():
    with settings(warn_only=True):
        result = run('reboot')
#        if result.failed and not confirm("Everything's ready, press 'y' please ?"):
 #           abort("Aborting at user request.")
           #disconnect_all()
            #sys.exit()

def deploy():
    check()
    listing()
    transfer()
    restart()
