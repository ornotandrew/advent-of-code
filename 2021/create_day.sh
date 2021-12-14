#!/bin/bash
day_number=$1
dir=day-$day_number

cargo new $dir
touch $dir/input.txt
rm $dir/src/main.rs
cp template.rs $dir/src/part_1.rs
cp template.rs $dir/src/part_2.rs

echo '

[[bin]]
name = "part_1"
path = "src/part_1.rs"

[[bin]]
name = "part_2"
path = "src/part_2.rs"' >> $dir/Cargo.toml

echo "Created $dir"
