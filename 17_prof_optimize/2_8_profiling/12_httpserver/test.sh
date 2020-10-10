#!/bin/bash
wrk -c100 -d10s -t50 http://127.0.0.1:8081/
