package dir

import (
	"strings"
	"path/filepath"
	"log"
	"os"
)

func GetTargetFiles(targetList []string)(files []string){
	for _, targetName := range targetList{
		if strings.Contains(targetName, "*"){
			if match, err:= filepath.Glob(targetName); err!=nil{
				log.Println("error path")
				log.Fatalln(err)
			}else{
				files = append(files, match...)
				continue
			}
		}

		file, err:= os.Open(targetName)
		if err!=nil{
			log.Fatalln(err)
		}
		state, err := file.Stat()
		if err!=nil{
			log.Fatalln(err)
		}
		if !state.IsDir(){
			files = append(files, targetName)
			continue
		}
		if err := filepath.Walk(targetName,  func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		}); err!= nil{
			log.Fatalln(err)
		}
	}
	return
}
