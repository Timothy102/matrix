package matrix

import (
	"log"
	"math"
)

//Vector struct
type Vector struct {
	row      []float64
	elements int
}

func newVector(slice []float64) *Vector {
	return &Vector{row: slice}
}

func innerProduct(matrix Matrix, vector1, vector2 Vector) float64 {
	var product Vector
	for _, r := range matrix.slice {
		for i := 0; i < len(r); i++ {
			product.row[i] = r[i] * vector1.row[i]
		}
	}
	return vector2.dotProduct(product)
}
func (v Vector) dotProduct(v2 Vector) float64 {
	var sum float64
	for i := 0; i < len(v.row); i++ {
		sum += v.row[i] * v2.row[i]
	}
	return sum
}
func (v Vector) applyMatrix(matrix Matrix) Vector {
	var product Vector
	for _, r := range matrix.slice {
		for i := 0; i < len(r); i++ {
			product.row[i] = r[i] * v.row[i]
		}
	}
	return product
}
func (v Vector) add(v2 Vector) Vector {
	var resultVector Vector
	for i := 0; i < len(v.row); i++ {
		resultVector.row[i] = v.row[i] + v2.row[i]
	}
	return resultVector
}
func (v Vector) substract(v2 Vector) Vector {
	var resultVector Vector
	for i := 0; i < len(v.row); i++ {
		resultVector.row[i] = v.row[i] - v2.row[i]
	}
	return resultVector
}
func (v Vector) addMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] + vec.row[i]
		}
	}
	return resultVector
}
func (v Vector) substractMany(vectors []Vector) Vector {
	var resultVector Vector
	for _, vec := range vectors {
		for i := 0; i < len(v.row); i++ {
			resultVector.row[i] = v.row[i] - vec.row[i]
		}
	}
	return resultVector
}
func (v Vector) getLength() float64 {
	var result float64
	for _, g := range v.row {
		result += math.Pow(g, 2)
	}
	return math.Sqrt(result)
}
func (v Vector) scalarProduct(scalar float64) Vector {
	for _, g := range v.row {
		g += scalar
	}
	return v
}
func (v Vector) angleBetween(v2 Vector) float64 {
	return math.Cos(v.dotProduct(v2) / (v.getLength() * v2.getLength()))
}
func (v Vector) scalarProjection(v2 Vector) float64 {
	x := v2.getLength() * v.angleBetween(v2)
	return x
}
func (v Vector) vectorProjection(v2 Vector) Vector {
	x := v.dotProduct(v2)
	y := v2.dotProduct(v2)
	return v.multiplyByScalar(x / y)
}
func (v Vector) multiplyByScalar(scalar float64) Vector {
	for _, g := range v.row {
		g = g * scalar
	}
	return v
}
func (v Vector) changingBasis(b1, b2 Vector) Vector {
	var newVector Vector
	var db1, db2 float64
	if b1.angleBetween(b2) != 0 {
		log.Fatal("basis vectors have to be perpendicular")
	} else {
		db1 = v.dotProduct(b1) / math.Pow(b1.getLength(), 2)
		db2 = v.dotProduct(b2) / math.Pow(b2.getLength(), 2)
	}
	newVector.row[0] = db1
	newVector.row[1] = db2

	return newVector
}

//2x2 matrix, todo later bigger dimensions
func calculateEigenvectors2x2(m Matrix) ([]Vector, error) {
	var vectors []Vector
	q := m.slice
	det := m.findDeterminant()
	l1, l2 := quadratic(1, q[0][0]+q[1][1], det)
	vec1 := []float64{l1, 0}
	vec2 := []float64{0, l2}
	vector1 := Vector{row: vec1}
	vector2 := Vector{row: vec2}

	vectors = append(vectors, vector1)
	vectors = append(vectors, vector2)
	return vectors, nil

}
func quadratic(a, b, c float64) (float64, float64) {
	disc := math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(disc)) / 2 * a
	x2 := (-b - math.Sqrt(disc)) / 2 * a
	return x1, x2

}
