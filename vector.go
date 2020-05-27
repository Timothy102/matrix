package matrix

import (
	"fmt"
	"math"
)

//Vector struct
type Vector struct {
	row      []float64
	elements int
}

//NewVector returns a vector type
func NewVector(slice []float64) Vector {
	return Vector{row: slice}
}

//InnerProduct returns the inner product of two vectors via the matrix.
//Check the dotProduct function for a simplified dot product of two vectors.
func InnerProduct(matrix Matrix, vector1, vector2 Vector) float64 {
	var product Vector
	for _, r := range matrix.slice {
		for i := 0; i < len(r); i++ {
			product.row[i] = r[i] * vector1.row[i]
		}
	}
	return vector2.DotProduct(product)
}

//Slice returns vector.slice
//you can perform indexing with this method
func (m Vector) Slice() []float64 {
	return m.row
}

//DotProduct returns the dot product of two vectors.
func (v Vector) DotProduct(v2 Vector) float64 {
	var sum float64
	for i := 0; i < len(v.row); i++ {
		sum += v.row[i] * v2.row[i]
	}
	return sum
}

//ApplyMatrix returns the vector through a matrix transformation.
func (v Vector) ApplyMatrix(matrix Matrix) Vector {
	var product Vector
	for _, r := range matrix.slice {
		for i := 0; i < len(r); i++ {
			product.row[i] = r[i] * v.row[i]
		}
	}
	return product
}

//Add returns an elementary operation on two vectors.
func (v Vector) Add(v2 Vector) Vector {
	var resultVector Vector
	for i := 0; i < len(v.row); i++ {
		resultVector.row[i] = v.row[i] + v2.row[i]
	}
	return resultVector
}

//Substract returns an elementary operation on two vectors.
func (v Vector) Substract(v2 Vector) Vector {
	var resultVector Vector
	for i := 0; i < len(v.row); i++ {
		resultVector.row[i] = v.row[i] - v2.row[i]
	}
	return resultVector
}

//AddMany takes a slice of vectors and outputs the sum of them.
func (v Vector) AddMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] + vec.row[i]
		}
	}
	return resultVector
}

//SubstractMany takes a slice of vectors and outputs the divergence of them.
func (v Vector) SubstractMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] - vec.row[i]
		}
	}
	return resultVector
}

//GetLength returns the length of a vector.
func (v Vector) GetLength() float64 {
	var result float64
	for _, g := range v.row {
		result += math.Pow(g, 2)
	}
	return math.Sqrt(result)
}

//AngleBetween returns the angle between two vectors.
func (v Vector) AngleBetween(v2 Vector) float64 {
	return math.Cos(v.DotProduct(v2) / (v.GetLength() * v2.GetLength()))
}

//ScalarProjection returns the scalar projection of v onto v2.
func (v Vector) ScalarProjection(v2 Vector) float64 {
	x := v2.GetLength() * v.AngleBetween(v2)
	return x
}

//MultiplyByScalar operates by adding a scalar elementary.
func (v Vector) MultiplyByScalar(scalar float64) Vector {
	for _, g := range v.row {
		g = g * scalar
	}
	return v
}

//VectorProjection returns the vector projection of v onto v2.
func (v Vector) VectorProjection(v2 Vector) Vector {
	x := v.DotProduct(v2)
	y := v2.DotProduct(v2)
	return v.MultiplyByScalar(x / y)
}

//ChangingBasis returns the transformed vector if we were to change its basis vectors.
//Similar to a simple matrix transformation but inputs are two basis vectors.
//Basis vectors have to be perpendicular to each other in order to apply this transformation.
//The function takes care of this for you by sending back an error
func (v Vector) ChangingBasis(b1, b2 Vector) (Vector, error) {
	var newVector Vector
	var db1, db2 float64
	if b1.AngleBetween(b2) != 0 {
		return Vector{}, fmt.Errorf("basis vectors have to be perpendicular")
	}
	db1 = v.DotProduct(b1) / math.Pow(b1.GetLength(), 2)
	db2 = v.DotProduct(b2) / math.Pow(b2.GetLength(), 2)

	newVector.row[0] = db1
	newVector.row[1] = db2

	return newVector, nil
}
