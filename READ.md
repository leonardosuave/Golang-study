# executed files
- Every executed files need to have package main, and func main() {}.

# package main
- its a files group in the same directory and are compiled together.

# module
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