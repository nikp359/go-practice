#!/bin/bash

netstat -n | grep -i 8099 | grep -i time_wait | wc -l