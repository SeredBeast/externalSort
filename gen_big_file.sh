#!/bin/bash

# 10M*1000 = 10GB
dd if=/dev/urandom bs=10M count=1000 2>/dev/null | base64 >source.txt

