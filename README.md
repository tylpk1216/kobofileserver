# kobofileserver
Run it on Kobo device, then use browser to transfer file to device.

# Installation
01. Install NickelMenu, I use it to launch application.
02. Put KoboFileServer folder in /mnt/onboard/.adds
03. Create "/mnt/onboard/kobofileserver" folder.
04. Modify NickelMenu config file.
```
menu_item :main    :Force Wi-Fi On (toggle)  :nickel_setting     :toggle:force_wifi
menu_item :main    :IP Address               :cmd_output         :500:/sbin/ifconfig | /usr/bin/awk '/inet addr/{print substr($2,6)}'
menu_item :main    :Import Books             :nickel_misc        :rescan_books_full
menu_item :main    :KoboFileServer (toggle)  :cmd_output         :500:quiet  :/usr/bin/pkill -f "^/mnt/onboard/.adds/KoboFileServer/kobofileserver"
  chain_success:skip:3
  chain_failure                              :cmd_spawn          :/mnt/onboard/.adds/KoboFileServer/kobofileserver
  chain_failure                              :dbg_toast          :Error starting KoboFileServer
  chain_always:skip:-1
  chain_success                              :dbg_toast          :Stopped KoboFileServer
```
05. Adjust your sleep settings of device for transfering large file. The processing speed is 0.52MB/second on Kobo Elipsa when I use HyRead Gaze Pocket to uploading file. If you use cell phone, I think the speed is better.

# How to use it
01. Click "Force Wi-Fi On (toggle)" of NickelMenu.
02. Turn-on Wi-Fi of device.
03. Click "KoboFileServer (toggle)" of NickelMenu.
04. Click "IP Address" of NickelMenu to get your device IP.
05. Open browser on any device, then type "http://IP/".
06. Select a file.
07. Click "Click to upload file" to upload file.
08. After uploading file is done, click "Import Books" of NickelMenu.

# Test Video
[Use Android E-Ink device(HyRead Gaze Pocket) to upload files to Kobo Elipsa.](https://youtu.be/mZ4C3v0sqL0 "kobofileserver")

# License
MIT
