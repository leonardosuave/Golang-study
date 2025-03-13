# Executed files
- Every executed files need to have package main, and func main() {}.

# Package main
- its a files group in the same directory and are compiled together.

# Module
- The modules are like package.json, where there are all dependencies neccessaries and the module do the files management to interact between 2 or more files, like a module from Angular or NestJS.
- To create a module file, execute "go mod init module" and will create a go.mod file
- The command "go build" will compile all code from the application and will create a binarie file "module.exe" and after to run the binarie file execute "./module".
- When create another file that will exported to another file, the function name its important, because the function its visible to another files when import the files and the function start in UPPERCASE.
    - Function started with LOWERCASE = visible only in your own package
    - Function started with UPPERCASE = visible to files that imported the package (Need to do a comment before the function to show that is a exported function).

    - This work like Public, Private and Protected.

- Install package from a reposItory
    - command "go get link". Obs: without https
- Removed package from a repository not used    
    - command "go mod tidy"

- To call functions from a imported package, call the last name after the last bar from the "import". Ex: "github.com/badoux/checkmail", call only checkmail.    

# Variables
To create variables, can be done by 2 types.
- Write var and the name and type. Ex: var example string = "value"
    - To create two or more variables in same time. 
        Ex: var (
            variable1 string = "value"
            variable2 number = 8
        ) 
- Without var. Ex: example2 := "value". This case the type its take by the value and its not neccessary declair the type like by var.
    - To create two or more variables in same time. 
        Ex: variable1, variable2 := "value1", "value2"

- OBS: It can change var by const and this work like javascript where not allowed change const values.

# Types
- Basic types: 
    - int: int, int8, int16, int32 and int64 (int default is by computer system default. Ex computer 64bits, default int64).
    - uint: uint, uint8, uint16, uint32, uint64 (uint default is by computer system default. Ex computer 64bits, default uint64)
    - alias: can put a type by alias. int32 = rune and uint8 = byte.
    - float: float32 or float64 (float default by computer system. Ex computer 64bits, default float64).
    - string: string.
    - boolean: boolean (true or false).
    - error: use go package errors.New() to create error.

- start values
    - int and float: start as 0.
    - string: start as empty
    - boolean: start as false.
    - error: start as nil (nil === null from javascript)