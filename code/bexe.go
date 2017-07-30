package main

import(
"fmt"
"os"
"archive/zip"
"io"
"os/exec"
"strings"
"io/ioutil"
"flag"
)


//funcion para añadir los archivos a zip y verificar que termina con la extersion .exe
func add_file(filename string,filezip *os.File,archive *zip.Writer){
	info,_:=os.Stat(filename)
	header,_:=zip.FileInfoHeader(info)
	if strings.HasSuffix(info.Name(),".exe") == true{

		}else{
			header.Name = "name"+".txt"
		}
	header.Method= zip.Deflate
	writer,_:= archive.CreateHeader(header)
	file,_:=os.Open(filename)
	io.Copy(writer,file)


}

func main(){
	var names string
	exe_uno := flag.String("exe1","",`example="C:\\Users\\Anonimo\\Desktop\\exe1.exe"`)
	exe_dos := flag.String("exe2","",`example="C:\\Users\\Anonimo\\Desktop\\exe2.exe"`)
	output :=flag.String("output","",`example="C:\\Users\\Anonimo\\Desktop\\output.exe"`)
	flag.Parse()
	//abrir binario
	b,err:= ioutil.ReadFile("include\\bin.exe")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	ejecutable := string(b)


	if *exe_uno != "" && *exe_dos != "" && *output !=""{
		if info,err:=os.Stat(*exe_uno);err == nil && true== strings.HasSuffix(info.Name(),".exe"){
			names+=info.Name()+"\n"
			if info,err:=os.Stat(*exe_dos);err == nil && true==strings.HasSuffix(info.Name(),".exe"){
				//Crear Zip 
				filezip,_:=ioutil.TempFile("","")
				defer filezip.Close()
				//defer os.Remove(filezip.Name())
				filename:=filezip.Name()
				
				archive:=zip.NewWriter(filezip)
				defer archive.Close()

				//Añadir Archivo a Zip
				names+=info.Name()+"\n"
				filenames,_:= ioutil.TempFile("","")
				defer filenames.Close()
		       
		        filenames.WriteString(names)
				add_file(*exe_uno,filezip,archive)
		        add_file(*exe_dos,filezip,archive)
		        
		        add_file(filenames.Name(),filezip,archive)

				ejecutable_temporal,_:=ioutil.TempFile("","")
				defer ejecutable_temporal.Close()
				ejecutable_temporal.WriteString(ejecutable)
                //union de exe + zip
			    cmd:= exec.Command("cmd","/c","copy /b "+ejecutable_temporal.Name()+"+"+filename+" "+(*output))
			    cmd.Start()
			    fmt.Println("Output:"+*output)
		       
		        


	   
		        
		}else{
			fmt.Println("File not found:"+*exe_dos)
		}
	}else{
		fmt.Println("File not found:"+*exe_uno)
	}

}else{
banner := `
 /$$$$$$$  /$$$$$$$$ /$$   /$$ /$$$$$$$$
| $$__  $$| $$_____/| $$  / $$| $$_____/
| $$  \ $$| $$      |  $$/ $$/| $$      
| $$$$$$$ | $$$$$    \  $$$$/ | $$$$$   
| $$__  $$| $$__/     >$$  $$ | $$__/   
| $$  \ $$| $$       /$$/\  $$| $$      
| $$$$$$$/| $$$$$$$$| $$  \ $$| $$$$$$$$
|_______/ |________/|__/  |__/|________/ 
By Alanprogrammer            Version:1.0`
fmt.Println(banner)
fmt.Println(`usage->bexe.exe -exe1="[exe]" -exe2="[exe]" -output="[output]"`)
fmt.Println(`example->bexe.exe -exe1="win.exe" -exe2="door.exe" -output="house.exe"`)
}

}