package fileOperation

import "os"

func ReadFile(fileName string) (string, error) { // read the file contents and return them as a string
	file, err := os.Open(fileName)
	fileContent := ""
	if err != nil {
		return fileContent, err
	} else {
		stat, _ := file.Stat()
		// fmt.Print(stat.Size())
		arr := make([]byte, stat.Size())
		file.Read(arr)
		fileContent = string(arr)
		file.Close()
	}
	return fileContent, err
}

func SaveFile(fileName string, contents string) error { // read the file contents and return them as a string
	file, err := os.Create(fileName)
	if err != nil {
		return err
	} else {
		_, writeErr := file.WriteString(contents)
		if writeErr != nil {
			return writeErr
		}
		file.Close()
		return err
	}
}