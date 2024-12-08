# List Util

`list-util` is a Go implementation of the `ls` command that supports several flags including the default behavior, `-1`, `-l`, `-h`, `-a`, and combinations of these flags like e.g `-lha`. This tool lists the contents of a directory in a user-friendly format, providing various ways to display file and directory information.

## Features

- **Default `ls` behavior**: Lists the files and directories in the current directory in a simple, readable format.
- **`-1` flag**: Displays each file or directory on a new line (one per line).
- **`-l` flag**: Provides a detailed listing of files and directories, including permissions, owner, group, size, and modification time.
- **`-h` flag**: Displays file sizes in a human-readable format (e.g., K, M, G).
- **`-a` flag**: Includes hidden files (those starting with a dot).
- **Flag combinations**: Supports combining multiple flags, such as `-lha` to show detailed, human-readable sizes including hidden files.

## Installation

To build and use `list-util`, you need Go installed on your system. Follow these steps to install and run the tool:

1. Clone the repository:

   ```bash
   git clone https://github.com/Anuolu-2020/list-util.git
   cd list-util
   ```
2. Build Project
  
   ```bash
   go build -o list-util
   ```   
3. Run the tool
   ```bash
   ./list-util [flags] 
   ```

## Usage
 The basic usage is:
 ```bash
 ./list-util [flags] 
 ```

### Flags

- `-1`: Display one file per line.
- `-l`: Display detailed file information.
- `-h`: Show file sizes in human-readable format.
- `-a`: Include hidden files (files starting with a dot).
- `-help`: Display usage of flags

### Example Commands

- `./list-util`: List files and directories in the current directory.
- `./list-util -1`: List files and directories with one item per line.
- `./list-util -l`: Show detailed information about files and directories.
- `./list-util -h`: Display file sizes in human-readable format.
- `./list-util -a`: Include hidden files in the listing.
- `./list-util -lha`: Show detailed information with human-readable sizes and include hidden files.

## To Be Implemented

 While the current version of `list-util` supports several key flags, there are still some commands and features that are planned for future releases:

- **`-R` (Recursive)**: Recursively list subdirectories in the specified directory.
- **`-S` (Sort by size)**: Sort files and directories by their size.
- **`-t` (Sort by modification time)**: Sort files and directories by their last modification time.
- **`--color`**: Add color-coding to the output based on file type and permissions.
- **`-i` (Inode number)**: Display the inode number of each file.
- **`-d` (Directories only)**: List directories themselves, not their contents.
  
