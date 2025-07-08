**Repository:** Where codes are hosted and developers works in collaboration.

**Modules and Packages:** Go code is grouped into packages and packages are grouped into modules.

**Modules:** A module specifies the dependencies needed to run our code, including the Go version and the set of other modules it required in the go.mod file.

**Creating module**

`go mod init <unique-identifier>` : To create a new go module rooted at the current directory.

`go mod init example.com/learn`

* go.mod file consists of the module declaration, the minimum compatible version for go, and the dependencies for the imported third-party packages.

* Note: Collection of go code becomes a module when there is a valid go.mod file in its root directory.

`go mod tidy`: It will add (import) any missing module that are required in go.mod file.

`replace director`: To use a local module that is not pushed to version control yet. We use a replace director.

`go build`: Compile the packages named by the import paths, along with their dependencies into an executable (in the current directory).

`go install`: compile and move the executable to $GOPATH/bin. We can run this executable from any path of the terminal.

`go env GOPATH`: tells the gopath.

`go get`: download and install remote Go modules/packages, add them to your go.mod file and update the existing module versions.

**Developing and publishing modules**

* Inititilize git repo.
* Check repo status.
* Tag commit after git add and git commit (`git tag v0.1.0`): It tags the commit with the version number.
* Push the tag to the github (`git push origin v0.1.0`)

**godoc**

* `go doc PACKAGE_NAME`: It parses the source file and extracts the comments to produce HTML or plain text documentation.
* `go doc PACKAGE_NAME.IDENTIFIER_NAME`
* Note: Place a regular comment just before package declaration without any blank lines in go source file.

**Package name collision**

When two or more packages share the same name and are imported in a go source code, To use thema and avoid its package name collision, we can use `alias`

```
import (
    crand "crypto/rand"
    "math/rand"
)
```