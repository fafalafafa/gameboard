#!/bin/bash

redis-cli flushall
redis-cli rpush characters "doctor"
redis-cli rpush characters "mafia"
redis-cli rpush characters "mafia"
redis-cli rpush characters "investigator-1"
redis-cli rpush characters "investigator-2"
redis-cli rpush characters "civilian"
redis-cli rpush characters "civilian"
redis-cli rpush characters "civilian"
redis-cli rpush characters "civilian"
redis-cli rpush characters "civilian"
