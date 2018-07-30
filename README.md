# trav
This package is an extension to the native [ast.Walk](https://golang.org/pkg/go/ast/#Walk) function. It traverses the Go AST and calls a hander function for each node. It includes the whole path that leads to the current node as a slice.

# Installation
This package is go-gettable.

    $ go get -u github.com/noxer/trav

# Usage
Start by defining your handler function, it is called once for each node in the AST.

    func handler(p trav.Path) {
        // do stuff
    }

`trav.Path` is a slice of `ast.Node`s, you can `range` over it. Call `p.Current()` to get the current node (the last one in the slice). If you want to continue using the slice once the handler has returned, you can call `p.Copy` to create a copy of the slice. The original slice will be reused in the next iteration.
Now you can call one of the `trav.Traverse*()` functions to traverse the AST.

    trav.TraverseFile("code.go", handler)

