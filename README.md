## Search

Search is a little tool that I made to search file on my computer. Search is designed for linux.

# Installation & running

You will need go1.22

Clone the repository and the run
````
go get ./...
````
It will install all the dependencies of the project.
Then run this to created executable:
````
go build search.go -o search
````
This command will create a search file that you can execute like this:
````
./search
````
