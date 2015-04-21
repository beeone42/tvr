#!/usr/bin/env python
from __future__ import with_statement
from fabric.api import *
from fabric.contrib.console import confirm
from fabric.contrib.files import exists
import os
import sys

# gauche = 10.62.1.2
# droite = 10.62.1.3

env.hosts = ['10.62.1.2', '10.62.1.3']
#env.user = "root"
#env.password = "openelec"

def check():
    print env.hosts[0]
    print "erasing old videos"
'''
def listing():
    run('ls')
    put('videos', '/storage')

def transfer():
    with quiet():
        if os.path.isfile('*.avi'):
            put('*.avi', '/storage/videos')
if os.path.isfile('*.mkv'):
    put('*.mkv', '/storage/videos')
if os.path.isfile('*.mp4'):
    put('*.mp4', '/storage/videos')
if os.path.isfile('all.m3u'):
    put('all.m3u', '/storage/videos')

def restart():
    with settings(warn_only=True):
        result = run('reboot')
        if result.failed and not confirm("Everything's ready, press 'y' please ?"):
            abort("Aborting at user request.")
            #disconnect_all()
            #sys.exit()
'''
def deploy():
    check()
#    listing()
 #   transfer()
  #  restart()
