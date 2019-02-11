// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package fdlimit

import "syscall"

// Raise tries to maximize the file descriptor allowance of this process
// to the maximum hard-limit allowed by the OS.
func Raise(max uint64) error {
	// Get the current limit
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return err
	}
	// Try to update the limit to the max allowance
	limit.Cur = limit.Max
	if limit.Cur > max {
		limit.Cur = max
	}
	// https://developer.apple.com/library/archive/documentation/System/Conceptual/ManPages_iPhoneOS/man2/getrlimit.2.html
	//
	// setrlimit() now returns with errno set to EINVAL in places that histori-cally historically
	// cally succeeded.  It no longer accepts "rlim_cur = RLIM_INFINITY" for
	// RLIM_NOFILE.  Use "rlim_cur = min(OPEN_MAX, rlim_max)".
	//
	// OPEN_MAX is 10240
	if limit.Cur > 10240 {
		return errors.New("darwin only allows OPEN_MAX (10240) open files")
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return err
	}
	return nil
}

// Current retrieves the number of file descriptors allowed to be opened by this
// process.
func Current() (int, error) {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	return int(limit.Cur), nil
}

// Maximum retrieves the maximum number of file descriptors this process is
// allowed to request for itself.
func Maximum() (int, error) {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	if limit.Max > 10240 { // OPEN_MAX is 10240
		return 10240, nil
	}
	return int(limit.Max), nil
}
