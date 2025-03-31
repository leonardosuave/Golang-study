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

- Func type:
    - When instaced a variable like a function, variable has func type with all types params and funct return type
    - Funct can return 2 or more values, its neccessary only intaced the variable quantity equal return quantity. To ignore a return, change name variable by _     

- start values
    - int and float: start as 0.
    - string: start as empty
    - boolean: start as false.
    - error: start as nil (nil === null from javascript)

# Operators
- The operators are equal javascript, but not use 3 equals (===) and havent ternário operator(need to check by if/else)

# Structs
- The structs are like class, where can create a specific object with a type specific.
    - To this, create a type "name" struct {"object format"} and put type in each value.
    - Can declair values like: 
        - object.value = value;
        - object := nameStruct{"value", "value2"} - values into {} are in the order declair in the struct
        - object:= nameStruct{value: "value1"} - When there are optional values, call the key that has values only

# Heritage
- In GO, not have heritage like TypeScript with extends and type things. But we can reuse a struct into another to not declair equals values.
    - To this, its neccessary only pass the struct name into another struct like a key without redeclair the type, so this is like a typescript's extends.

# Arrays and Slices
- Go has 2 types to do a list of values, arrays and slices.
    - Arrays: the arrays are fixeds, so we need to declair how many index and the unic type accept. After i can not push a new value, only change a exist index by another same type value.
    - Slices: the slices are flixibler than arrays, so i can create a slice with x indexs and after push a new value without error (need to be same type).
    - OBS: Arrays and Slices are differents, so i can not do operation between then.    

# Pointer
- The pointer is a variable that save some things memory address (like JS reference when create a new variable with value from another variable and edit the value and automaticly edit original value too)
- To create a pointer variable, need to be "*type"
- Default value is "nil"
- To attribute a value into a pointer variable, the variable attributed need to declair with & beforer variable name.
    - Ex: var variable2 *int = &variable1 (where variable1 is a int with a value)
- When print pointer variable, will show the memory address and not variable value attributed
- To show pointer variable value, pass as *variable2.    

# Map
- The maps are structure with key: value and are allow only one type value to key and value. If i declair that are string, so only string are allows to create key and values.
- To create a map. Ex"
    - user2 := map[string]string {"key": "value"}
    - string into [] is the type of key and string outside is type value.
    - OBS: Different than structs, the keys and values need to be declair like "value" when string.

# Control Structure
- The control structures are if/elses used to control a flow.
- Its possible create a variable directtly in the if line and this variable will be avaliable only to the if/else block that was created.

# Switch case
- The switch work like javascript, but not neccessary put "break" after the case to out switch.
    - The switch can be done 3 different form:
        - switch "variable"
            case "value or conditional":
                return "value"

        - switch {
            case "conditional with variable":
                return "value"
        }

        var "receiveValue" type
        switch {
            case "conditional":
                "receiveValue" = "value"
        }
        return "receiveValue"        

# For (looping)
- The for structure is unic, can be written only 1 format, not equal JS.
- The FOR can be by control variable with increments or by range to looping array/slices
- Ex:
    - Increment:
    for i := 0; i < 10; i++ {

    }

    - Range
	sliceValues := []string{"some 1", "some 2", "some 3", "some 4", "some 5"}
    for i, value := range sliceValues {

    }
    - OBS: To range, need to declar the 2 params (index, value). If want only value, declair index as _  .

# Functions
- There are a lot of type of functions.
    - Named return: Functions that declair return variables in the type of function and is not neccessary return the variables
        - func calculate(n1, n2 int) (sum int, subtraction int) {
            return 
        }
        - Will return the sum and subtraction without do nothing.

    - Variadic: Functions that declair only one params type, but can send a lot of values to this params. Like spreed operator from JS.
        - func showValues(txt string, n ...int) {

        }
        - When txt its only a string value and n can be a lot of int values.

    - Anonymous functions: As JS, are functions without names and executed by the instance. The params sended are declair in the end of function.

    - Recursive functions: are functions call yourself - Its not often used.

    - Defer: Are function executed after all another logic. Ex: After all logic and beforer a return.

    - Panic and Recover:
        - Panic: Its different Erro, Panic kill the programm and nothing is execution after the panic, but beforer its call all defer function.
        - Recover: Recover its a function that recover everything from a panic function to not kill the application.
    
    - Closure
        - Are function that reference variable outside body(scope)

    - Pointer function
        - Like as variables, function can received a memory address by reference and change the value. This will reflect in the variable value referenced.

    - Init function
        - init function is a fisrt function executed in the file. Its executed first than main() and can have 1 init function by file, different than main() that can have only one by package.
        - Can be used to init a variables or another configuration first the main().

    - Method
        - The methods sare equals functions, but aways associate with something.
        - To create a method, create a func but before between () declair a variable and struct type referenced and after function name
            - Ex: func (u user) save() {}.
        - When call the struct its same that create a intanced from struct and after this variables has access to all structs methods created.
        - To modify a value existed in the method, use pointer to change value from memory address.    

