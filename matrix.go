//Package matrix implements matrix and vector operations along with some difficult methods such as the Gram-Schmidt process, Einstein summation convention, eigenvector calculation and more.
package matrix

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

)

//Matrix type does matrix math
type Matrix struct {
	slice [][]float64
}

//NewMatrix returns a matrix and an error
func NewMatrix(slice [][]float64) Matrix {
	rows := sliceRows(slice)
	columns := sliceColumns(slice)

	if columns < 2 || rows < 2 {
		log.Fatalf("This is not a matrix. Please, enter a proper number of elements.")
	}

	return Matrix{slice: slice}
}

//helper functions
func sliceColumns(slice [][]float64) int {
	return len(slice[len(slice)-1])
}

//helper functions
func sliceRows(slice [][]float64) int {
	return len(slice)
}

//NumberOfColumns returns the number of columns.
func (m Matrix) NumberOfColumns() int {
	return sliceColumns(m.slice)
}

//NumberOfRows returns the number of columns.
func (m Matrix) NumberOfRows() int {
	return sliceRows(m.slice)
}

//Dimensions returns the number of rows and columns of m.
func (m Matrix) Dimensions() (int, int) {
	return m.NumberOfRows(), m.NumberOfColumns()
}

//NumberOfElements returns the number of elements.
func (m Matrix) NumberOfElements() int {
	return m.NumberOfColumns() * m.NumberOfRows()
}

//RoundtoDecimals round all elements of a matrix to decimals accuracy.
func (m Matrix) RoundtoDecimals(decimals int) Matrix {
	for _, r := range m.slice {
		for _, e := range r {
			e = roundTo(e, decimals)
		}
	}
	return m
}
func roundTo(number float64, decimals int) float64 {
	s := math.Pow(10, float64(decimals))
	return math.Round(number*s) / s
}

//Randomize randomizes m to random values
func (m Matrix) Randomize() Matrix {
	row := m.NumberOfRows()
	column := m.NumberOfColumns()
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			m.slice[i][j] = rand.Float64() * 000.3
		}
	}
	return m
}

//RandomValuedMatrix returns a row*column random valued matrix.
func RandomValuedMatrix(row, column int) Matrix {
	rand.Seed(time.Now().UnixNano())
	slc := make([][]float64, row)
	for i := 0; i < row; i++ {
		sl := make([]float64, column)
		for j := 0; j < column; j++ {
			sl[j] = randFloats(row * column)[i*j]
			slc[i] = sl
		}
	}
	return Matrix{slice: slc}
}
func randFloats(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	fls := make([]float64, n)
	for i := range fls {
		fls[i] = rand.Float64()
	}
	return fls
}

//Slice returns matrix.slice
// You can perform indexing with this method.
func (m Matrix) Slice() [][]float64 {
	return m.slice
}

//PrintByRow prints the matrix by row.
func (m Matrix) PrintByRow() {
	for r := range m.slice {
		fmt.Println(m.slice[r])
	}
}

//At method finds the value at rowIndex,columnIndex
func (m *Matrix) At(rowIndex, columnIndex int) float64 {
	return m.slice[rowIndex][columnIndex]
}

//Identity function returns an n*n identity matrix
func Identity(n int) Matrix {
	matrix := Matrix{}
	k := 0
	for i := 0; i < n; i++ {
		slice := make([]float64, n)
		slice[k] = 1
		k++
		matrix.slice = append(matrix.slice, slice)
	}
	return matrix
}

//Zeros returns a matrix of zeros.
func Zeros(row, column int) Matrix {
	b := make([][]float64, row)
	v := make([]float64, column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			v[j] = 0
			b[i] = v
		}
	}

	return Matrix{slice: b}
}

//Ones returns a matrix of ones.
func Ones(row, column int) Matrix {
	b := make([][]float64, row)
	v := make([]float64, column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			v[j] = 1
			b[i] = v
		}
	}

	return Matrix{slice: b}
}

//AllSameNumber returns a row * column matrix of number.
func AllSameNumber(row, column int, number float64) Matrix {
	b := make([][]float64, row)
	v := make([]float64, column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			v[j] = number
			b[i] = v
		}
	}

	return Matrix{slice: b}
}

//FromValues returns a matrix from a set of values
func FromValues(row, column int, values []float64) Matrix {
	slice := make([][]float64, row)
	for i := range values {
		if (i+1)%column == 0 {
			slc := make([]float64, column)
			slc = append(slc, values[i+1], values[i+2], values[i+2])
			slice = append(slice, slc)
		}
	}
	return Matrix{slice: slice}
}

//Sigmoid returns sigmoid of x.
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

//SigmoidPrime returns the sigmoid derivative of x.
func SigmoidPrime(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}

//Matmul does the matrix multiplication. A's rows must match B's columns
func Matmul(a, b matrix.Matrix) matrix.Matrix {
	result := matrix.RandomValuedMatrix(a.NumberOfRows(), b.NumberOfColumns())
	for i := 0; i < a.NumberOfRows(); i++ {
		for j := 0; j < b.NumberOfColumns(); j++ {
			summa := 0.0
			for k := 0; k < a.NumberOfColumns(); k++ {
				summa += a.Slice()[i][k] * b.Slice()[k][j]
			}
			result.Slice()[i][j] = summa
		}
	}
	return result
}

//Add performs elementary matrix addition
func (m Matrix) Add(mat Matrix) Matrix {
	var product Matrix
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			product.slice[i][j] = m.slice[i][j] + mat.slice[i][j]
		}
	}
	return product
}

//Subtract performs elementary matrix subtraction
func (m Matrix) Subtract(mat Matrix) Matrix {
	var product Matrix
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			product.slice[i][j] = m.slice[i][j] - mat.slice[i][j]
		}
	}
	return product
}

//Multiply performs elementary matrix multiplication
func (m Matrix) Multiply(mat Matrix) Matrix {
	var product Matrix
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			product.slice[i][j] = m.slice[i][j] * mat.slice[i][j]
		}
	}
	return product
}

//Divide performs elementary matrix division
func (m Matrix) Divide(mat Matrix) Matrix {
	var product Matrix
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			product.slice[i][j] = m.slice[i][j] / mat.slice[i][j]
		}
	}
	return product
}

//ScalarMultiplication multiplies every element with a scalar
func (m Matrix) ScalarMultiplication(scalar float64) Matrix {
	for _, r := range m.slice {
		for i := range r {
			r[i] = r[i] * scalar
		}
	}
	return m
}

//ScalarAdition adds a scalar to every elements
func (m Matrix) ScalarAdition(scalar float64) Matrix {
	for _, r := range m.slice {
		for i := range r {
			r[i] = r[i] + scalar
		}
	}
	return m
}

//Transpose returns the tranpose of a matrix
func (m Matrix) Transpose() Matrix {
	for i, r := range m.slice {
		for j := range r {
			m.slice[i][j] = m.slice[j][i]
		}
	}
	return m
}

//FindDeterminant returns the matrix determinant
func (m Matrix) FindDeterminant() float64 {
	dims := m.NumberOfRows()
	var determinant, p float64
	for k := 0; k < dims; k++ {
		if k%2 == 0 {
			p = -1.0
		} else {
			p = 1.0
		}
		if dims == 1 {
			log.Fatalf("This is a single valued matrix.")
		} else if dims == 2 {
			determinant += m.slice[0][k] * m.Shorten(0, k).Find2x2Determinant() * p
		} else {
			determinant += m.slice[0][k] * m.Shorten(0, k).FindDeterminant() * p
		}
	}
	return determinant
}

//Find2x2Determinant returns the determinant of a 2x2 matrix
func (m Matrix) Find2x2Determinant() float64 {
	return m.slice[0][0]*m.slice[1][1] - m.slice[1][0]*m.slice[1][0]
}

//Shorten returns the so-called minor matrix, it shrinks all numbers that lie either with one coordinate on rowIndex or columnIndex
func (m Matrix) Shorten(rowIndex, columnIndex int) Matrix {
	for j, r := range m.slice {
		for i := range r {
			m.slice[rowIndex][j] = 0.0
			m.slice[i][columnIndex] = 0.0
			m.slice[i][j] = m.slice[i-1][j-1]
		}
	}
	return m
}

//Adjoint returns the adjoint matrix
func (m Matrix) Adjoint() (Matrix, error) {
	for i, r := range m.slice {
		for j := range r {
			m.slice[i][j] = math.Pow(-1, float64(i+j)) * m.Shorten(i, j).FindDeterminant()
		}
	}
	return m, nil
}

//Inverse returns the inverse of a matrix
func (m Matrix) Inverse() Matrix {
	var inverse Matrix
	det := m.FindDeterminant()
	adjoint, err := m.Adjoint()
	if err != nil {
		log.Fatalf("unable to create adjoint matrix :%v", err)
	}
	inverse = adjoint.ScalarMultiplication(1 / det)
	return inverse
}

//Inverse2x2 returns the inverse of a 2x2 matrix
func (m Matrix) Inverse2x2() Matrix {
	if m.NumberOfRows() != 2 {
		log.Fatalf("This is not a 2x2 matrix.")
	}
	var result Matrix
	result.slice[0][0] = m.slice[1][1]
	result.slice[1][1] = m.slice[0][0]
	result.slice[0][1] = -m.slice[0][1]
	result.slice[1][0] = -m.slice[1][0]
	return result
}

//EinsteinConvention returns the multiplication matrix of two matrices, given that rows of A matches columns of B.
//According to this convention, when an index variable appears twice in a single term and is not otherwise defined, it implies summation of that term over all the values of the index.
func (m Matrix) EinsteinConvention(m2 Matrix) Matrix {
	var result Matrix
	sum := 0
	for range m2.slice {
		sum++
	}
	if len(m.slice) != sum {
		log.Fatal("Rows of A must match columns of B")
	}
	for n := 0; n < sum; n++ {
		for h := 0; h < len(m.slice); h++ {
			for i := 0; i < sum; i++ {
				for j := 0; j < len(m2.slice); j++ {
					result.slice[n][h] += m.slice[2][i] * m2.slice[j][3]
				}
			}
		}
	}
	return result
}

//DotProduct returns the dot product of two matrices
func (m Matrix) DotProduct(m2 Matrix) float64 {
	var sum float64
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			sum += m.slice[i][j] * m2.slice[i][j]
		}
	}
	return sum
}

//FromArray returns a matrix from array
func FromArray(arr []float64) Matrix {
	m := Zeros(len(arr), 1)
	for i := 0; i < len(arr); i++ {
		m.slice[i][0] = arr[0]
	}
	return m
}

//ToArray returns the matrix in array form.
func (m Matrix) ToArray() []float64 {
	var arr []float64
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			arr = append(arr, m.slice[i][j])
		}
	}
	return arr
}

//MapFunc applies f to every element
func (m Matrix) MapFunc(f func(x float64) float64) Matrix {
	for i := 0; i < m.NumberOfRows(); i++ {
		for j := 0; j < m.NumberOfColumns(); j++ {
			m.slice[i][j] = f(m.slice[i][j])
		}
	}
	return m
}

//TransformationInAChangedBasis function takes a given matrix as an input and outputs it in a changed basis
func (m Matrix) TransformationInAChangedBasis(basis Matrix) Matrix {
	inv := basis.Inverse()
	transform := inv.Multiply(m)
	result := transform.Multiply(basis)
	return result
}
