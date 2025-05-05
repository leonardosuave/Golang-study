# Resume
This project will run commanders by commander line and will take the ips and servers.

## First steep
- Create go.mod -> go mod init "app name"
- Install package -> go get github.com/urfave/cli (it will create go.sum as package.json)
- Create a main.go in the root.
- Import the function with the logic.
- Create folder app and after the app.go with logic runned in each commander line created.
    - Import the package intalled
    - Name and Usage are required to cli run.
    - Commands is a slice/array and each index is a line commander
    - The commander need to have Name, Usage, Flags(with Name and Value defaults) and Action(external function with the logic executed in the commander)

## Run commanders
- To run commanders without build, execute:
    - go run main.go "commander name" --host "sites url" -> go run main.go ip --host devbook.com.br

- To build app
    - go build (will compile the app in .exe file (faster to execute the things))
    - path to file .exe "commander name" --host "sites url" -> ./commander-line ip --host google.com  