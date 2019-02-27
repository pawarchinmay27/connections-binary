package main

import (
	"fmt"
	"log"
	ps"github.com/mitchellh/go-ps"
	"github.com/pytimer/win-netstat"
)

func tcp4() {
	conns, err := winnetstat.Connections("tcp4")
	if err != nil {
		log.Fatal(err)
	}
	for _, conn := range conns {
		fmt.Printf("%s:%d\t%s:%d\t%d\t%s\n", conn.LocalAddr, conn.LocalPort,conn.RemoteAddr,conn.RemotePort,conn.OwningPid, conn.State)
		p, err := ps.FindProcess(conn.OwningPid)
		if err != nil {
 		fmt.Println("Error : ", err)	
 	}

 	fmt.Println("Process ID : ", p.Pid())
 	fmt.Println("Parent Process ID : ", p.PPid())
 	fmt.Println("Process ID binary name : ", p.Executable())
	}
}

func tcp4WithPid(pid int) {
	conns, err := winnetstat.ConnectionsWithPid("tcp4", pid)
	if err != nil {
		log.Fatal(err)
	}
	for _, conn := range conns {
		fmt.Println(conn)
		fmt.Printf("%s:%d\t%d\t%s\n", conn.LocalAddr, conn.LocalPort, conn.OwningPid, conn.State)
	}
}

func main() {
	tcp4()
	//tcp4WithPid(3848)
}