# Torrent-Downloader 🚀

A blazing fast command-line **BitTorrent downloader** written in Go.  
Built as part of a learning challenge and now extended into a personal project, this tool helps you interact with `.torrent` files and magnet links with ease.

---

## 🔧 Features

- 🔍 **Decode** Bencoded strings
- 🧲 **Parse** and handle Magnet links
- 🤝 **Perform Handshakes** with peers
- 📥 **Download** specific pieces or entire files
- 🔌 **Connect to Peers** and explore peer information
- 📂 **Fetch Torrent Metadata**

---

## 🛠️ Commands

```bash
# Decode a bencoded string
torrent-downloader decode "d3:cow3:moo4:spam4:eggse"

# Get info from a .torrent file
torrent-downloader info path/to/file.torrent

# List peers for a torrent
torrent-downloader peers path/to/file.torrent

# Perform handshake with a peer
torrent-downloader handshake path/to/file.torrent 127.0.0.1:6881

# Download a specific piece
torrent-downloader download_piece <output_path> path/to/file.torrent <piece_index>

# Download the entire file
torrent-downloader download <output_path> path/to/file.torrent

# Handle Magnet links
torrent-downloader magnet_parse "magnet:?xt=urn:btih:..."
torrent-downloader magnet_info "magnet:?xt=urn:btih:..."
torrent-downloader magnet_handshake "magnet:?xt=urn:btih:..."


Getting Started
# Clone the repository
git clone https://github.com/pvk-96/Torrent-Downloader.git
cd Torrent-Downloader

# Run (requires Go installed)
go run main.go <command> [args...]
