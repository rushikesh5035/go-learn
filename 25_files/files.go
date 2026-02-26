package main

import (
	"fmt"
	"os"
)

// files


func main(){
	// f, err :=os.Open("example.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err) // panic means to stop the program and print the error message	
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File name:", fileInfo.Name()) // example.txt
	
	// fmt.Println("Is directory:", fileInfo.IsDir()) // check if it's a directory (folder or file)
	
	// fmt.Println("File size:", fileInfo.Size()) // get file size in bytes -> `Hello Golang` total 12 letters
	
	// fmt.Println("File Permission Mode:", fileInfo.Mode()) // file permission mode -> -rw-r--r-- (read and write for owner, read for group and others)	
	
	// fmt.Println("File Modification Time:", fileInfo.ModTime()) // last modified time -> 2024-06-01 12:00:00 +0000 UTC
	
	// fmt.Println("File Modification Time:", fileInfo.ModTime()) // last modified time -> 2024-06-01 12:00:00 +0000 UTC

	//-----------------------READ FILE----------------------------
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close() // ensure the file is closed after we're done with it

	// buffer := make([]byte, 12) // create a buffer to hold the file content

	// d, err := f.Read(buffer) // read the file content into the buffer
	
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i<len(buffer); i++{
	// 	fmt.Println("Data", d, string(buffer[i]))
	// }

	//-------------------------------------------------------------
	// READ File using os.ReadFile (simpler way)
	// it will directly read the entire file content and return it as a byte slice

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))
	
	
	//------------------------------------------------------------
	// READ DIRECTORY/ FOLDER
	//-------------------------------------------------------------
	

	// dir, err := os.Open(".")
	// dir, err := os.Open("../") // open the parent directory
	// if err != nil {
	// 	panic(err)
	// }
	
	// defer dir.Close()

	// // fileInfo, err := dir.ReadDir(2) // read the directory entries, 2 means read 2 file 

	// fileInfo, err := dir.ReadDir(-1) // read all the directory entries, -1 means read all files

	// for _, file := range fileInfo {
	// fmt.Println("File name:", file.Name())
	// }


	//------------------------------------------------------------
	// CREATE FILE
	//------------------------------------------------------------

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	
	// // f.WriteString("Hi Go") // create file and write "Hi Go" to it
	// // f.WriteString("Nice Language") // write "Nice Language" to the file

	// bytes := []byte ("Hello golang, welcome to file handling in Go!") // create a byte slice with the content to write

	// f.Write(bytes) // write the byte slice to the file, it will overwrite the existing content in the file


	//------------------------------------------------------------
	// READ and WRITE to ANOTHER FILE (streaming fashion)
	//------------------------------------------------------------

	// sourceFile, err := os.Open("example.txt") // open the source file to read from
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt") // create the destination file to write to
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	
	// reader := bufio.NewReader(sourceFile) // create a buffered reader for the source file

	// writer := bufio.NewWriter(destFile) // create a buffered writer for the destination file

	// for {
	// 	b, err := reader.ReadByte() // read a byte by byte from the source file
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err) // panic if the error is not end of file
	// 		}
	// 		break // break the loop if we reach the end of the file
	// 	}	

	// 	e := writer.WriteByte(b) // write the byte to the destination file

	// 	if err != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush() // flush the buffered writer to ensure all data is written to the destination file

	// fmt.Println("Written to new filesuccessfully!")

	//------------------------------------------------------------
	// DELETE FILE
	//------------------------------------------------------------
	err := os.Remove("example2.txt") // delete the file, it will return an error if the file does not exist
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted Successfully!")

}

// 1. we open the source file and close it after we're done with it using defer
// 2. we create the destination file and close it after we're done with it using defer
// 3. we create a buffered reader for the source file and a buffered writer for the destination file using bufio package
// bufio -> it provides buffered I/O operations, which can improve performance when reading and writing files by reducing the number of system calls
// 4. we read the source file byte by byte using reader.ReadByte() and write it to the destination file using writer.WriteByte() inside the loop until we reach the end of the file (EOF)
// 5. we flush the buffered writer to ensure all data is written to the destination file using writer.Flush() at the end of the loop