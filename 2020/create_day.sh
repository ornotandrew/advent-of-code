#!/bin/bash
day_number=$1
dir=day-$day_number

mkdir $dir
touch $dir/input.txt
cp template.py $dir/part_1.py
cp template.py $dir/part_2.py

echo "Created $dir"
