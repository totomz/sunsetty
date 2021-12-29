#!/bin/bash
cd /home/pi/sunsetty

for i in {0..10}; do    
    image="/tmp/image_$(date +%d%m%Y_%H%M%S).jpg"
    echo "$(date) today I'm taking $image / $i" >> /home/pi/sunsetty/dio.log
    
    # Get the screen from the camera
    fswebcam -p YUYV -r 1024x768 -S 80 --rotate 180 $image
    echo "$(date) pic taken" >> /home/pi/sunsetty/dio.log
    
    ./dropbox_uploader.sh upload $image /
    echo "$(date) pic uploaded" >> /home/pi/sunsetty/dio.log
    
    # Wait for the next minute
    # thanks to https://unix.stackexchange.com/a/439589/81902
    sleep $(( 60 - 10#$(date +%S) ));
    
done;

# Send the email
/home/pi/sunsetty/mail $image >> dio.log 2>&1
echo "$(date) email sent" >> /home/pi/sunsetty/dio.log
