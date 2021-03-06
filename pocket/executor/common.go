package executor

// Select offer unified method for single & abtest
func (e *Executor) Select(stmt string) error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestSelect(stmt)
	case "single":
		return e.SingleTestSelect(stmt)
	}
	panic("unhandled select switch")
}

// Insert offer unified method for single & abtest
func (e *Executor) Insert(stmt string) error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestInsert(stmt)
	case "single":
		return e.SingleTestInsert(stmt)
	}
	panic("unhandled select switch")
}

// Update offer unified method for single & abtest
func (e *Executor) Update(stmt string) error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestUpdate(stmt)
	case "single":
		return e.SingleTestUpdate(stmt)
	}
	panic("unhandled select switch")
}

// Delete offer unified method for single & abtest
func (e *Executor) Delete(stmt string) error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestDelete(stmt)
	case "single":
		return e.SingleTestDelete(stmt)
	}
	panic("unhandled select switch")
}

// ExecDDL offer unified method for single & abtest
func (e *Executor) ExecDDL(stmt string) error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestExecDDL(stmt)
	case "single":
		return e.SingleTestExecDDL(stmt)
	}
	panic("unhandled select switch")
}

// TxnBegin offer unified method for single & abtest
func (e *Executor) TxnBegin() error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestTxnBegin()
	case "single":
		return e.SingleTestTxnBegin()
	}
	panic("unhandled txn begin switch")
}

// TxnCommit offer unified method for single & abtest
func (e *Executor) TxnCommit() error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestTxnCommit()
	case "single":
		return e.SingleTestTxnCommit()
	}
	panic("unhandled txn commit switch")
}

// TxnRollback offer unified method for single & abtest
func (e *Executor) TxnRollback() error {
	e.Lock()
	defer e.Unlock()
	switch e.mode {
	case "abtest":
		return e.ABTestTxnRollback()
	case "single":
		return e.SingleTestTxnRollback()
	}
	panic("unhandled txn rollback switch")
}

