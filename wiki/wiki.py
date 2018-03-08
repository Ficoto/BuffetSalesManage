# !/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import subprocess


current_dir = os.path.join(os.path.abspath(os.path.dirname(os.path.dirname(__file__))))

root_dir = os.path.abspath(os.path.join(current_dir, os.pardir))

target_dir = os.path.abspath(os.path.join(root_dir, os.pardir))

print "current_dir: ", current_dir
print "target_dir: ", target_dir

wiki_dir = current_dir + "/markdown"
print "wiki_dir: ", wiki_dir

if os.path.exists(wiki_dir):
    print "wiki dir exist."
    p = subprocess.Popen('git pull', cwd=wiki_dir, shell=True)
else:
    print "wiki dir not exist."
    sh = 'git clone git@gitlab.xinghuolive.com:Backend-Go/kangaroo.wiki.git markdown'
    p = subprocess.Popen(sh, cwd=current_dir, shell=True)

md = "go run {0} --toml={1}" .format(current_dir + "/main/wiki_main.go", current_dir + '/wiki.toml')
print "md :", md
p = subprocess.Popen(md, cwd=target_dir, shell=True)
p.wait()

print "wiki dir", wiki_dir
sh = 'git add --all && git commit -m "Update Wiki" && git push origin master'
p = subprocess.Popen(sh, cwd=wiki_dir, shell=True)
p.wait()
