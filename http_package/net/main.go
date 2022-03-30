package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// membuat koneksi ke server golang.org melalui tcp network
	conn, err := net.Dial("tcp", "golang.org:80")
	// periksa jika ada error saat koneksi ke server
	// jika ada error tidak akan melanjutkan ke proses berikutnya
	// dan keluar dari program `log.Fatal(...)`
	if err != nil {
		log.Fatal("The following error occured", err)
	}
	// koneksi tersambung, kirim GET request ke server
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// kirim request "enter" untuk mengakhiri koneksi
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("Send request error ", err)
	}

	fmt.Println("Response: ", status)
}
