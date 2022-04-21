#!/bin/bash

ps ax | grep 'data-warehouse producer run' | awk '{print $1}' | xargs kill