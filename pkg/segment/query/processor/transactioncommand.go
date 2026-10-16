// Copyright (c) 2026 The SigScalr Authors.
// SPDX-License-Identifier: Apache-2.0

package processor

import (
	"errors"

	"github.com/siglens/siglens/pkg/segment/query/iqr"
	"github.com/siglens/siglens/pkg/segment/structs"
)

type transactionProcessor struct {
	options *structs.TransactionArguments
}

// Paged Out fork: the command is not implemented in this pipeline; report it instead of crashing the server.
func (p *transactionProcessor) Process(iqr *iqr.IQR) (*iqr.IQR, error) {
	return nil, errors.New("transaction command is not supported")
}

func (p *transactionProcessor) Rewind() {
}

func (p *transactionProcessor) Cleanup() {
}

func (p *transactionProcessor) GetFinalResultIfExists() (*iqr.IQR, bool) {
	return nil, false
}
