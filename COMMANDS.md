#### Run a module
> cd /path/to/mod/folder   
> go run .   
> go run modname


#### Generate swagger documents
In addition to the setup from the reference video, run the following commands to support swag command in terminal:
> go install github.com/swaggo/swag/cmd/swag@latest   
> echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc   
> source ~/.zshrc   
> swag --version   

Command to generate docs:
> swag init --dir cmd/api --parseDependency —-parseInternal —-parseDepth 1   
> Then, go to http://localhost:8080/swagger/index.html


#### Add Hot Reloading support
> go install github.com/air-verse/air@latest
> which air OR air --version
> air init
> Then run, air
