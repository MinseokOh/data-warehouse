#!/bin/bash

./build/data-warehouse producer run --config ./configs/producer_1_config.json > ./build/producer/producer1.log &
./build/data-warehouse producer run --config ./configs/producer_2_config.json > ./build/producer/producer2.log &
./build/data-warehouse producer run --config ./configs/producer_3_config.json > ./build/producer/producer3.log &
./build/data-warehouse producer run --config ./configs/producer_4_config.json > ./build/producer/producer4.log &
./build/data-warehouse producer run --config ./configs/producer_5_config.json > ./build/producer/producer5.log &
