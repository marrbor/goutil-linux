package goutil

import (
	"golang.org/x/sys/unix"
)

func uname2str(u []byte) string {
	str := ""
	for _, v := range u {
		m := int(v)
		if m <= 0 {
			break
		}
		str += string(m)
	}
	return str
}

func GetUnixVer() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Version[:]), nil
}

func getUnixSysName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Sysname[:]), nil
}

func getUnixNodeName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Nodename[:]), nil
}

func getUnixMachineName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Machine[:]), nil
}

func getUnixRelease() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Release[:]), nil
}
