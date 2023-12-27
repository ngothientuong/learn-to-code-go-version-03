# SET UP GO PROJECT
## Set up new project
- mkdir -p <project_folder>
- cd <project_folder>
- go mod init <project_folder>
- Create a go.work file at the TOP level of workspace, i.e. `/ngo/project` folder
- Add path of **all** <project_folders> that contains go.mod files, e.g.
    ```
    go 1.21.5

    use (
        ./corbra-cli/02-cobra
        ./learn-to-code-go-version-03	
    )

    ```
## Add local go package
- `go get ./path/to/local/package/package_folder`
- From here, you can reference the above package's function in the `.go` file, go will import that package for you