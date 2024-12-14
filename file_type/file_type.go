package file_type

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileType struct represents the deprecated way of handling file objects.
type FileType struct {
	mode     string
	bufsize  int
	encoding string
	errors   string
}

// NewFileType creates a new FileType object.
func NewFileType(mode string, bufsize int, encoding, errors string) *FileType {
	// Warning about deprecation
	fmt.Println("Warning: FileType is deprecated. Simply open files after parsing arguments.")

	return &FileType{
		mode:     mode,
		bufsize:  bufsize,
		encoding: encoding,
		errors:   errors,
	}
}

// Open opens the file based on the input string.
func (ft *FileType) Open(filename string) (*os.File, *bufio.Reader, error) {
	// If the filename is "-", treat it as stdin or stdout
	if filename == "-" {
		if strings.Contains(ft.mode, "r") {
			// Opening stdin
			return os.Stdin, bufio.NewReader(os.Stdin), nil
		} else if strings.Contains(ft.mode, "w") || strings.Contains(ft.mode, "x") {
			// Opening stdout
			return os.Stdout, bufio.NewReader(os.Stdout), nil
		} else {
			return nil, nil, fmt.Errorf("invalid mode %s for special filename \"-\"", ft.mode)
		}
	}

	// Try to open the file normally
	file, err := os.OpenFile(filename, getFileMode(ft.mode), os.FileMode(ft.bufsize))
	if err != nil {
		return nil, nil, fmt.Errorf("can't open '%s': %v", filename, err)
	}
	return file, bufio.NewReader(file), nil
}

// getFileMode converts the mode string to appropriate file mode in Go.
func getFileMode(mode string) int {
	var flags int
	if strings.Contains(mode, "r") {
		flags |= os.O_RDONLY
	}
	if strings.Contains(mode, "w") {
		flags |= os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	if strings.Contains(mode, "x") {
		flags |= os.O_EXCL
	}
	if strings.Contains(mode, "a") {
		flags |= os.O_APPEND
	}
	return flags
}

// String returns the string representation of the FileType object.
func (ft *FileType) String() string {
	return fmt.Sprintf("FileType(mode=%s, bufsize=%d, encoding=%s, errors=%s)",
		ft.mode, ft.bufsize, ft.encoding, ft.errors)
}

// // Example function demonstrating usage of FileType
// func main() {
// 	// Create a FileType object (deprecating use, just for illustration)
// 	fileType := NewFileType("r", -1, "", "")

// 	// Open a file (or stdin/stdout)
// 	file, reader, err := fileType.Open("test.txt")
// 	if err != nil {
// 		log.Fatalf("Error: %v\n", err)
// 	}

// 	// If the file opened successfully, use it
// 	if file != nil {
// 		defer file.Close()
// 	}

// 	// For demonstration, just read and print the file content or from stdin
// 	if reader != nil {
// 		content, err := reader.ReadString('\n')
// 		if err != nil {
// 			log.Fatalf("Error reading file: %v\n", err)
// 		}
// 		fmt.Println("Content:", content)
// 	}
// }
