#!/bin/bash

multitail $(find ./build -maxdepth 2 -name '*.log')