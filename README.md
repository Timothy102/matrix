# matrix
A linear-algebra based Go library for matrices, vectors, as well as the Gram-Schmidt process,Einstein summation convention and more.


[![GoDoc](https://godoc.org/github.com/Timothy102/matrix?status.svg)](https://godoc.org/github.com/Timothy102/matrix)

This Go Linear Algebra Library is based on the popular Mathematics for Machine Learning Specialization on Coursera. It involves classic methods of vector and matrix operations, such as dot and inner product, and some tougher challenges, such as finding the eigenvectors of a given matrix or the Gram-Schmidt process.



# Installation
Downloading the package should be fairly simple. Run the code below in your directory terminal.
```go
go get github.com/timothy102/matrix
```
Once you have done that, import the package in your go file.
```go
import "github.com/timothy102/matrix"
```
If you prefer not to write matrix all the time or for some other reason you want to change the name, type this.
```go
import name "github.com/timothy102/matrix"
```

# Usage
Let's take a look at simple matrix addition. 
 ```go

  m1 :=matrix.RandomMatrix(3,2)
  m2 :=matrix.RandomMatrix(3,2) 
  result :=m1.Add(m2)

  result.PrintByRow()
  
 ```
 And the output:
 ```go
  [[7.2,4.2,5.3]
  [9.4,8.5,8.6]]
  
 ```
 
 Same applies for vectors.
 
 ```go
 vector:=NewVector([]float64{2.0,4.0,3.2})
 vector.MultipliedByScalar(2.0)
 ```
 Output:
 ```go
 [4.0,8.0,6.4]
 ```
 
 
I hope this library offers you to dig deeper into the world of linear algebra and to apply it to some cool machine learning concepts.
Looking forward for feedback!
 
 
# Contact

LinkedIn : (https://www.linkedin.com/in/tim-cvetko-32842a1a6/)

Medium : (https://cvetkotim03.medium.com/)

Gmail : (tim@metawaveai.com)
 
