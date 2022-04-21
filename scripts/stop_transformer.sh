#!/bin/bash

ps ax | grep 'data-warehouse transformer run' | awk '{print $1}' | xargs kill