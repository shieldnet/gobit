/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package jwt_maker

type Keys struct {
	Secret []byte `json:"secret"`
	Access string `json:"access"`
}

type KeySet []Keys