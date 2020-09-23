package matrix

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Vector struct {
	row []float64
}

//NewVector returns a vector type
func NewVector(slice []float64) Vector {
	return Vector{row: slice}
}

// PrintVector prints the vector components
func PrintVector(vec Vector) {
	fmt.Println(vec.row)
}

//VecToArray returns a slice of vector elements.
func VecToArray(vec Vector) []float64 {
	slice := make([]float64, len(vec.row))
	for i := range vec.row {
		slice = append(slice, vec.row[i])
	}
	return slice
}

//RandomVector returns a random valued vector.
func RandomVector(size int) Vector {
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Float64()/0.3)
	}
	return NewVector(slice)
}

// InnerProduct returns the inner product of two vectors via the matrix.
// Check the dotProduct function for a simplified dot product of two vectors.
func InnerProduct(matrix Matrix, vector1, vector2 Vector) float64 {
	var product Vector
	for _, r := range matrix.slice {
		for i := 0; i < len(r); i++ {
			product.row[i] = r[i] * vector1.row[i]
		}
	}
	return vector2.DotProduct(product)
}

func (v Vector) NumberOfElements() int {
	var sum int
	for i := 0; i < len(v.row); i++ {
		sum += i
	}
	return sum
}

// Slice returns vector.slice.
// You can perform indexing with this method.
func (v Vector) Slice() []float64 {
	return v.row
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

//AddMany takes a slice of vectors and outputs their sum.
func (v Vector) AddMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] + vec.row[i]
		}
	}
	return resultVector
}

//SubstractMany takes a slice of vectors and outputs their divergence.
func (v Vector) SubstractMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] - vec.row[i]
		}
	}
	return resultVector
}

//Map maps the vector by with the function
func (vec Vector) Map(f func(float64) float64) Vector {
	for i := range vec.row {
		vec.row[i] = f(vec.row[i])
	}
	return vec
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

//GramSchmidt function returns the concatenated matrix of orthogonal based vectors out of the vectors dataset.
//Based on the Mathematics for Machine learning Specialization.
func GramSchmidt(vectors []Vector) ([]Vector, error) {
	var normalizedvectors []Vector
	var xs []Vector

	for i, vector := range vectors {
		normalized := vectors[i-1].MultiplyByScalar(1 / vector.GetLength())
		if i != 1 {
			s := vectors[i].DotProduct(normalized)
			x := normalized.MultiplyByScalar(s)
			xs = append(xs, x)
			normalizedvectors[i] = vectors[i].Substract(vectors[i].SubstractMany(xs))
		}
	}
	return normalizedvectors, nil
}

//PageRank algorithm returns the probabilites of randomly ending up on either of the pages.
//Page connections should be put inside the link matrix.
func PageRank(linkMatrix Matrix, pages int) {
	slice := make([]float64, pages)
	for i := range slice {
		slice[i] = 1 / float64(pages)
	}
	initVector := Vector{row: slice}
	for i := 0; i < len(initVector.row); i++ {
		initVector = initVector.ApplyMatrix(linkMatrix)
	}
}

//CalculateEigenvectors2x2 returns the eigenvectors of the given 2x2 matrix
func CalculateEigenvectors2x2(m Matrix) ([]Vector, error) {
	var vectors []Vector
	q := m.slice
	det := m.Find2x2Determinant()
	lambda1, lambda2 := Quadratic(1, q[0][0]+q[1][1], det)
	vec1 := []float64{lambda1, 0}
	vec2 := []float64{0, lambda2}
	vector1 := Vector{row: vec1}
	vector2 := Vector{row: vec2}

	vectors = append(vectors, vector1)
	vectors = append(vectors, vector2)
	return vectors, nil
}

//Quadratic takes a,b,c parameters of a quadratic equation as inputs and returns both solutions via the discriminant.
func Quadratic(a, b, c float64) (float64, float64) {
	disc := math.Pow(b, 2) - 4*a*c
	if disc < 0 {
		log.Fatalf("The discriminant is negative. Therefore solutions will be complex numbers. ")
	}
	x1 := (-b + math.Sqrt(disc)) / 2 * a
	x2 := (-b - math.Sqrt(disc)) / 2 * a
	return x1, x2
}
