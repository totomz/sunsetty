#!/bin/bash
cd /home/pi/sunsetty
echo "$(date) starting" >> /home/pi/sunsetty/dio.log

# Get the screen from the camera
fswebcam -p YUYV -r 1024x768 -S 60 /tmp/image.jpg
echo "$(date) pic taken" >> /home/pi/sunsetty/dio.log

# Send the email
/home/pi/sunsetty/mail /tmp/image.jpg >> dio.log 2>&1
echo "$(date) email sent" >> /home/pi/sunsetty/dio.log
