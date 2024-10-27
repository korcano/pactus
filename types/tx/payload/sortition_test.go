package payload

import (
	"io"
	"testing"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/types/amount"
	"github.com/pactus-project/pactus/util"
	"github.com/stretchr/testify/assert"
)

func TestSortitionType(t *testing.T) {
	pld := SortitionPayload{}
	assert.Equal(t, TypeSortition, pld.Type())
}

func TestSortitionDecoding(t *testing.T) {
	tests := []struct {
		raw      []byte
		value    amount.Amount
		readErr  error
		basicErr error
	}{
		{
			raw:      []byte{},
			value:    0,
			readErr:  io.EOF,
			basicErr: nil,
		},
		{
			raw: []byte{
				0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
				0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				0x11, 0x12, 0x13, 0x14, // address
			},
			value:    0,
			readErr:  io.ErrUnexpectedEOF,
			basicErr: nil,
		},
		{
			raw: []byte{
				0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
				0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				0x11, 0x12, 0x13, 0x14, 0x15, // address
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
				0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
				0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, // proof
			},
			value:    0,
			readErr:  io.ErrUnexpectedEOF,
			basicErr: nil,
		},
		{
			raw: []byte{
				0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
				0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				0x11, 0x12, 0x13, 0x14, 0x15, // address
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
				0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
				0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F, // proof
			},
			value:    0,
			readErr:  nil,
			basicErr: nil,
		},
		{
			raw: []byte{
				0x02, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
				0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
				0x11, 0x12, 0x13, 0x14, 0x15, // address
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
				0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
				0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F, // proof
			},
			value:   0,
			readErr: nil,
			basicErr: BasicCheckError{
				Reason: "address is not a validator address: pc1zqgpsgpgxquyqjzstpsxsurcszyfpx9q4350xtk",
			},
		},
	}

	for no, tt := range tests {
		pld := SortitionPayload{}
		r := util.NewFixedReader(len(tt.raw), tt.raw)
		err := pld.Decode(r)
		if tt.readErr != nil {
			assert.ErrorIs(t, err, tt.readErr)
		} else {
			assert.NoError(t, err)

			for i := 0; i < pld.SerializeSize(); i++ {
				w := util.NewFixedWriter(i)
				assert.Error(t, pld.Encode(w), "encode test %v failed", no)
			}
			w := util.NewFixedWriter(pld.SerializeSize())
			assert.NoError(t, pld.Encode(w))
			assert.Equal(t, pld.SerializeSize(), len(w.Bytes()))
			assert.Equal(t, tt.raw, w.Bytes())

			// Basic check
			if tt.basicErr != nil {
				assert.ErrorIs(t, pld.BasicCheck(), tt.basicErr)
			} else {
				assert.NoError(t, pld.BasicCheck())

				// Check signer
				if tt.raw[0] != 0 {
					assert.Equal(t, crypto.Address(tt.raw[:21]), pld.Signer())
				} else {
					assert.Equal(t, crypto.TreasuryAddress, pld.Signer())
				}

				assert.Equal(t, tt.value, pld.Value())
				assert.Nil(t, pld.Receiver())
			}
		}
	}
}
