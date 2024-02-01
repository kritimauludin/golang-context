# Created Go-Lang Context

- This code was written by	        : Kriti Mauludin
- Find my portfolio on      	      : https://kritimauludin.github.io
- I am actively writing articles on	: https://arahin.ocumps.com
- Special Thanks to Programer Zaman Now(PZN) For Created This Course

## Project setup and Installation
```
Website Go-Lang : https://golang.org/
1. Download Compiler Go-Lang
2. Install Compiler Go-Lang
3. Check in Command Prompt / Terminal with syntax -> go version

Last Save this folder in your desired go-lang directory.
```
### Compiles and hot-reloads for development
```
1. go build {filename}.go -> and then ->./{filename}
2. Or if you want it easier and faster use -> go run {filename}.go

But you can only use method number 2 for development
```
### Add dependencies
```
1. go get {name-module} will automatically download the latest version
```
### How to update dependencies to the latest version?
```
1. we can change the contents of the go.mod directly, then tag it to the latest tag version
2. after that to download the latest version, use the command: go get
```
### Created Module
```
1. to create a new module, use the command: go mod init {name-module}
2. add, commit and push modules to github
3. To release the version of the module, just create a tag in git, usually at the beginning of the word using keyword(v)
   for example {v1.0.0} or customize the version. run with command : git tag v1.0.0
4. lastly run the command: git push {remote-name} {v1.0.0} or adjust to your version

```
### How to upgrade module version?
when creating an application there are times when we need to upgrade the module version, to upgrade the module. 
```
we just need to create a new tag in git. With the following steps:
1. update your code module
2. add, commit, and push code
3. add tag and push again to github : git push {remote-name} {new-versio}
4. successfully release new version module

```
### What if there is a major upgrade/major changes ?
- Major upgrades usually occur because there is a change in the contents of the program code, thus making it not backward compatible
- This should be avoided as much as possible
- But if it's unavoidable, the best strategy is to change the module name

```
Best Practices:
- usually go programmers will change the module name by adding a keyword at the end of the module name using the version name
- hopefully if anyone does get module without (v2) they will get version 1
- so if you want to get v2 you need to add keywords behind it
```

```
So the strategy is:
1. Change the module name by adding a keyword at the end, for example (v2)
2. Create a new major changes tag as before, for example from v1.1.0 to v2.0.0
3. push tags to github as before
```
### Next ?
```
1. Go-Lang Basic 
2. Go-Lang Standard Library 
2. Go-Lang Modules 
3. Go-Lang Unit Test 
4. Go-Lang Goroutine
5. Go-Lang Database <-next position
6. Go-Lang Web 
```