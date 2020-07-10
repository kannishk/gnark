package r1cs

import (
	backend_bls377 "github.com/consensys/gnark/backend/bls377"
	backend_bls381 "github.com/consensys/gnark/backend/bls381"
	backend_bn256 "github.com/consensys/gnark/backend/bn256"
	backend_bw761 "github.com/consensys/gnark/backend/bw761"
	"github.com/consensys/gnark/encoding/gob"
	"github.com/consensys/gurvy"
)

type R1CS interface {
	Solve(assignment map[string]interface{}, _a, _b, _c, _wireValues interface{}) error
	Inspect(solution map[string]interface{}, showsInputs bool) (map[string]interface{}, error)
	GetNbConstraints() int
}

func Read(path string) (R1CS, error) {
	curveID, err := gob.PeekCurveID(path)
	if err != nil {
		return nil, err
	}
	var r1cs R1CS
	switch curveID {
	case gurvy.BN256:
		r1cs = &backend_bn256.R1CS{}
	case gurvy.BLS377:
		r1cs = &backend_bls377.R1CS{}
	case gurvy.BLS381:
		r1cs = &backend_bls381.R1CS{}
	case gurvy.BW761:
		r1cs = &backend_bw761.R1CS{}
	default:
		panic("not implemented")
	}

	if err := gob.Read(path, r1cs, curveID); err != nil {
		return nil, err
	}
	return r1cs, err
}
