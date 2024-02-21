package image_compression

import (
	"gonum.org/v1/gonum/mat"
)

func approximate(inputMatrix *mat.Dense, rank int) mat.Dense {
	u, vt, sigma := svd(inputMatrix)

	truncatedU := u.Slice(0, inputMatrix.RawMatrix().Rows, 0, rank)
	truncatedV := vt.Slice(0, rank, 0, inputMatrix.RawMatrix().Cols)
	truncatedSigma := sigma.Slice(0, rank, 0, rank)

	USigma := mat.NewDense(inputMatrix.RawMatrix().Rows, rank, nil)
	USigma.Mul(truncatedU, truncatedSigma)

	var resultMatrix mat.Dense
	resultMatrix.Mul(USigma, truncatedV)

	return resultMatrix
}

func svd(inputMatrix *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense) {
	var svd mat.SVD
	svd.Factorize(inputMatrix, mat.SVDFull) // main decomposition

	u, v, sigma, vt := &mat.Dense{}, &mat.Dense{}, &mat.Dense{}, &mat.Dense{}

	svd.UTo(u)
	svd.VTo(v)
	vt.CloneFrom(v.T())

	singularValues := svd.Values(nil)
	// firstly create diag matrix. Next fill new sigma matrix with zeros
	sigma.CloneFrom(mat.NewDiagDense(len(singularValues), singularValues))

	return u, vt, sigma
}
