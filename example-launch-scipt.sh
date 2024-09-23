#!/bin/bash

choice=$(zenity --list --radiolist \
    --title="New Timer" \
    --height 670 \
    --text="" \
    --column="" --column="Duration" \
    FALSE "1 minute" \
    FALSE "1 minute 30 seconds" \
    FALSE "2 minutes" \
    TRUE "2 minutes 30 seconds" \
    FALSE "3 minutes" \
    FALSE "5 minutes" \
    FALSE "7 minutes" \
    FALSE "10 minutes" \
    FALSE "15 minutes" \
    FALSE "30 minutes")

if [ -z "$choice" ]; then
    exit 1
fi

convert_to_seconds() {
    local time_str="$1"
    local minutes=$(echo "$time_str" | grep -o '[0-9]\+' | head -n1)
    local seconds=$(echo "$time_str" | grep -o '[0-9]\+' | tail -n1)

    if [[ "$time_str" == *"minute"* && "$time_str" != *"seconds"* ]]; then
        seconds=0
    fi

    echo $((minutes * 60 + seconds))
}


seconds=$(convert_to_seconds "$choice")

./mpris-timer "$seconds"
if [ $? -eq 0 ]; then
  notify-send -a Clock -i clock -u critical -e "Time is up!"
fi
