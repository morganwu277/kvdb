#!/usr/bin/env bash


for i in {1..1000}; do
  time ./kvdb-client 1000 >> run.log 2>&1
done
