////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package tester

// Assertion Methods help in comparing Values and checking Errors.

// This Package provides a convenient Way to control Mismatches.
// When the Comparison fails, it shows not only the two Values which failed,
// but also a 'Diff' (Difference) between them.

import (
	"fmt"
	"reflect"

	"github.com/kr/pretty"
)

// Error Messages.
const (
	ErrErrorWasExpected    = "An Error was expected, but None was received"
	ErrfNoErrorWasExpected = "No Error was expected, but One was received: %v"
	ErrfNotEqual           = `Values should be equal, but they are not. 
A=%v 
B=%v 
Diff=%v`
	ErrfNotDifferent = `Values should be different, but they are not. 
A=%v 
B=%v 
Diff=%v`
)

func errorIsSet(
	err error,
) bool {
	return err != nil
}

func errorIsEmpty(
	err error,
) bool {
	return err == nil
}

func interfacesAreEqual(
	a interface{},
	b interface{},
) bool {
	return reflect.DeepEqual(a, b)
}

func interfacesAreDifferent(
	a interface{},
	b interface{},
) bool {
	return !reflect.DeepEqual(a, b)
}

// Ensures that the Error is not nil. If the Error is nil, it stops the Test.
func (test *Test) MustBeAnError(
	err error,
) {
	if errorIsEmpty(err) {
		test.t.Error(ErrErrorWasExpected)
		test.t.FailNow()
	}
}

// Ensures that the Error is nil. If the Error is not nil, it stops the Test.
func (test *Test) MustBeNoError(
	err error,
) {
	if errorIsSet(err) {
		test.t.Errorf(ErrfNoErrorWasExpected, err)
		test.t.FailNow()
	}
}

// Ensures that two Variables have equal Values. If not, it stops the Test.
func (test *Test) MustBeEqual(
	a interface{},
	b interface{},
) {
	if interfacesAreDifferent(a, b) {
		msg := fmt.Sprintf(
			ErrfNotEqual,
			pretty.Sprint(a),
			pretty.Sprint(b),
			pretty.Diff(a, b),
		)
		test.t.Errorf(msg)
		test.t.FailNow()
	}
}

// Ensures that two Variables have different Values. If not, it stops the Test.
func (test *Test) MustBeDifferent(
	a interface{},
	b interface{},
) {
	if interfacesAreEqual(a, b) {
		msg := fmt.Sprintf(
			ErrfNotDifferent,
			pretty.Sprint(a),
			pretty.Sprint(b),
			pretty.Diff(a, b),
		)
		test.t.Errorf(msg)
		test.t.FailNow()
	}
}
