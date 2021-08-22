package main

import "test"

func main()  {
	test.Test()
}

// Can't work
// go: cannot find main module, but found .git/config in /docker/go/learn-golang
//         to create a module there, run:
//         cd ../../../.. && go mod init