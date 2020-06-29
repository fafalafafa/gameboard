#!/bin/bash

redis-cli flushall
redis-cli rpush characters "doctor"
redis-cli rpush characters "mafia"
redis-cli rpush characters "mafia"
redis-cli rpush characters "civilian"
redis-cli rpush characters "investigator"

