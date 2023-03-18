#!/bin/sh

gowsdl -d internal/services/ -p event -o event.go http://www.onvif.org/ver10/events/wsdl/event.wsdl
gowsdl -d internal/services/ -p device -o devicemgmt.go http://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl
gowsdl -d internal/services/ -p media -o media.go http://www.onvif.org/ver10/media/wsdl/media.wsdl
