#!/usr/bin/env python

import os
import sys
import json

def create_m3u():
    address = "/root/go/src/github.com/beeone42/tvr/tmp"
    print('generating new m3u file')
    name = 'all.m3u' 
    try:
        file = open(address + '/' + name,'w')
        file.close()
    except:
        print('Something went wrong! Can\'t tell what?')
        sys.exit(0)


def completing():
    address = "/root/go/src/github.com/beeone42/tvr/playlist"
    tmp = "/root/go/src/github.com/beeone42/tvr/tmp/all.m3u"
    f = open(address + '/' +  sys.argv[1] + '.json', 'r')
    #print f
    content = f.read()
    #print "voici" + content
    json_pl = json.loads(content)
    #print json_pl
    json_items = json_pl['Items']
    #print json_items
    c_dir = tmp
    p = open(tmp, 'r+')
    p.write('#EXTM3U\n')
    for vid in json_items:
        p.write('/storage/videos/' + vid + '\n')


create_m3u()
completing()
