package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"mytorrent/internal/codec"
	"mytorrent/internal/downloader"
	"mytorrent/internal/torrent"
)

func main() {
	fmt.Fprintln(os.Stderr, "MyTorrent CLI - A simple torrent tool")

	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use one of: decode, download, handshake, info, peers")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "decode":
		if len(os.Args) < 3 {
			log.Fatal("Usage: decode <bencoded_string>")
		}
		bencodedValue := os.Args[2]
		decoded, _, err := codec.Decode([]byte(bencodedValue))
		if err != nil {
			log.Fatal(err)
		}
		output, _ := json.MarshalIndent(decoded, "", "  ")
		fmt.Println(string(output))

	case "download":
		if len(os.Args) < 4 {
			log.Fatal("Usage: download <destination_path> <torrent_file>")
		}
		dest := os.Args[2]
		torrentFile := os.Args[3]
		if err := downloader.DownloadTorrent(torrentFile, dest); err != nil {
			log.Fatal(err)
		}

	case "handshake":
		if len(os.Args) < 4 {
			log.Fatal("Usage: handshake <torrent_file> <peer_address>")
		}
		torrentPath := os.Args[2]
		peerAddr := os.Args[3]
		client, err := torrent.New(torrentPath)
		if err != nil {
			log.Fatal(err)
		}
		peerParts := strings.Split(peerAddr, ":")
		if len(peerParts) != 2 {
			log.Fatal("Invalid peer address format. Expected ip:port")
		}
		port, _ := strconv.Atoi(peerParts[1])
		conn := client.ConnectToPeer(peerParts[0], uint16(port))
		if err := conn.Handshake(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Connected to peer. ID: %x\n", conn.PeerID)

	case "info":
		if len(os.Args) < 3 {
			log.Fatal("Usage: info <torrent_file>")
		}
		client, err := torrent.New(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Torrent Metadata:")
		fmt.Printf("Tracker: %s\n", client.Tracker())
		fmt.Printf("Length: %d bytes\n", client.Length())
		fmt.Printf("Info Hash: %x\n", client.InfoHash())

	case "peers":
		if len(os.Args) < 3 {
			log.Fatal("Usage: peers <torrent_file>")
		}
		client, err := torrent.New(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		peers := client.Peers()
		for _, p := range peers {
			fmt.Printf("%s:%d\n", p.IP, p.Port)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
