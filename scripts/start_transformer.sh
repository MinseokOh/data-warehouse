#!/bin/bash

./build/data-warehouse transformer run --config ./configs/transformer_config.json > ./build/transformer/transformer1.log &
./build/data-warehouse transformer run --config ./configs/transformer_config.json > ./build/transformer/transformer2.log &