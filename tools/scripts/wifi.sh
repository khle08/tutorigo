#!/bin/bash
#!/usr/bin/env bash

if [ "$1" = "list" ]
then
    # It takes some time to get the scanning result
    nmcli d wifi rescan
    # s: second / m: minute / h: hour / d: day
    sleep 5s
    # After scanning, the list will be more complete
    nmcli d wifi list

elif [ "$1" = "on" ]
then
    # Turn on the wifi of this computer
    nmcli r wifi on

elif [ "$1" = "off" ]
then
    # Turn off the wifi of this computer
    nmcli r wifi off

elif [ "$1" = "new" ]
then
    # Establish a new connection by providing wifi name and password
    nmcli d wifi connect "$2" password "$3"

elif [ "$1" = "connect" ]
then
    # Connect to the wifi that is connected before
    nmcli c up "$2"

elif [ "$1" = "disconnect" ]
then
    # Disconnect to the wifi and the established record will be saved
    nmcli c down "$2"

elif [ "$1" = "forget" ]
then
    # Delete the established record of the wifi name
    nmcli c delete "$2"

elif [ "$1" = "history" ]
then
    # Show all historical wifi connections
    nmcli c

fi