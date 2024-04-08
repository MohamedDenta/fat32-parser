package main

import (
	"fat32_parser/bootsector"
	"fat32_parser/directory"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("fat32.img") // Open the FAT32 image file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bootSector, err := bootsector.ReadBootSector(file) // Read the boot sector
	if err != nil {
		log.Fatal(err)
	}

	// Print boot sector information
	fmt.Printf("Boot Sector: %+v\n", bootSector)

	// Read directory entries
	dirEntries, err := directory.ReadDirectoryEntries(file)
	if err != nil {
		log.Fatal(err)
	}

	// Print directory entries
	for _, entry := range dirEntries {
		fmt.Printf("Directory Entry: %+v\n", entry)
		// Do something with directory entries
	}
}
