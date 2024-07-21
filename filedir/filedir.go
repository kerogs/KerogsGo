package filedir

import (
	"fmt"
	"os"
)

// FileMake creates a file with the given name and writes the given content to it.
//
// Parameters:
// - name: the name of the file to be created (string)
// - content: the content to be written to the file (string)
//
// Return type: none
func FileMake(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("file creation error :", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}