# matrix
A linear-algebra based Go library for matrices, vectors, as well as the pageRank algorithm, the Gram-Schmidt process and more.


You can find the full package documentation on GoDoc: https://godoc.org/github.com/Timothy102/matrix.

This Go Linear Algebra Library is based on the popular Mathematics for Machine Learning Specialization on Coursera. It involves classic methods of vector and matrix operations, such as dot and inner product, and some tougher challenges, such as finding the eigenvectors of a given matrix or the Gram-Schmidt process.


I hope this library offers you to dig deeper into the world of linear algebra and to apply it to some cool machine learning concepts.


Let's take a look at some initializing concepts.

Downloading the package should be fairly simple. Run the code below in your directory terminal.
```go get github.com/timothy102/matrix
```
Once you have done that, import the package in your go file.
```
import "github.com/timothy102/matrix"
```
If you prefer not to write matrix all the time or for some other reason you want to change the name, type this.
```
import name "github.com/timothy102/matrix"
```
Let's take a look at a simple matrix addition
```
slice:=[][]float64{
{3.0,4.0,3.0},
{4.0,2.0,1.0},
}

slice2:=[][]float64{
{4.3,0.2,2.3},
{5.4,6.5,7.6},
}
//you can also invoke the matrix struct itself of course.
//mat:=matrix.Matrix{}

mat1,err:=matrix.NewMatrix(slice,3,2)
//do error handling
mat2,err:=matrix.NewMatrix(slice2,3,2)
//do error handling

//Addition is simple
result:=mat1.Add(mat2)

//PrintByRow will print the matrix by row if you wish to see it in full display.
result.PrintByRow()
```
