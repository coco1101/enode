// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"bytes"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/enode/common"
)

func TestPack(t *testing.T) {
	for i, test := range []struct {
		typ        string
		components []ArgumentMarshaling
		input      interface{}
		output     []byte
	}{
		{
			"uint8",
			nil,
			uint8(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint8[]",
			nil,
			[]uint8{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint16",
			nil,
			uint16(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint16[]",
			nil,
			[]uint16{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint32",
			nil,
			uint32(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint32[]",
			nil,
			[]uint32{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint64",
			nil,
			uint64(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint64[]",
			nil,
			[]uint64{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint256",
			nil,
			big.NewInt(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"uint256[]",
			nil,
			[]*big.Int{big.NewInt(1), big.NewInt(2)},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int8",
			nil,
			int8(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int8[]",
			nil,
			[]int8{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int16",
			nil,
			int16(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int16[]",
			nil,
			[]int16{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int32",
			nil,
			int32(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int32[]",
			nil,
			[]int32{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int64",
			nil,
			int64(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int64[]",
			nil,
			[]int64{1, 2},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int256",
			nil,
			big.NewInt(2),
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"int256[]",
			nil,
			[]*big.Int{big.NewInt(1), big.NewInt(2)},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
		},
		{
			"bytes1",
			nil,
			[1]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes2",
			nil,
			[2]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes3",
			nil,
			[3]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes4",
			nil,
			[4]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes5",
			nil,
			[5]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes6",
			nil,
			[6]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes7",
			nil,
			[7]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes8",
			nil,
			[8]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes9",
			nil,
			[9]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes10",
			nil,
			[10]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes11",
			nil,
			[11]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes12",
			nil,
			[12]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes13",
			nil,
			[13]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes14",
			nil,
			[14]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes15",
			nil,
			[15]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes16",
			nil,
			[16]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes17",
			nil,
			[17]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes18",
			nil,
			[18]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes19",
			nil,
			[19]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes20",
			nil,
			[20]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes21",
			nil,
			[21]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes22",
			nil,
			[22]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes23",
			nil,
			[23]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes24",
			nil,
			[24]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes25",
			nil,
			[25]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes26",
			nil,
			[26]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes27",
			nil,
			[27]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes28",
			nil,
			[28]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes29",
			nil,
			[29]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes30",
			nil,
			[30]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes31",
			nil,
			[31]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"bytes32",
			nil,
			[32]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"uint32[2][3][4]",
			nil,
			[4][3][2]uint32{{{1, 2}, {3, 4}, {5, 6}}, {{7, 8}, {9, 10}, {11, 12}}, {{13, 14}, {15, 16}, {17, 18}}, {{19, 20}, {21, 22}, {23, 24}}},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000b000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000f000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000110000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001300000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000015000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000170000000000000000000000000000000000000000000000000000000000000018"),
		},
		{
			"address[]",
			nil,
			[]common.Address{{1}, {2}},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000"),
		},
		{
			"bytes32[]",
			nil,
			[]common.Hash{{1}, {2}},
			common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000201000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"function",
			nil,
			[24]byte{1},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			"string",
			nil,
			"foobar",
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000006666f6f6261720000000000000000000000000000000000000000000000000000"),
		},
		{
			"string[]",
			nil,
			[]string{"hello", "foobar"},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002" + // len(array) = 2
				"0000000000000000000000000000000000000000000000000000000000000040" + // offset 64 to i = 0
				"0000000000000000000000000000000000000000000000000000000000000080" + // offset 128 to i = 1
				"0000000000000000000000000000000000000000000000000000000000000005" + // len(str[0]) = 5
				"68656c6c6f000000000000000000000000000000000000000000000000000000" + // str[0]
				"0000000000000000000000000000000000000000000000000000000000000006" + // len(str[1]) = 6
				"666f6f6261720000000000000000000000000000000000000000000000000000"), // str[1]
		},
		{
			"string[2]",
			nil,
			[]string{"hello", "foobar"},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000040" + // offset to i = 0
				"0000000000000000000000000000000000000000000000000000000000000080" + // offset to i = 1
				"0000000000000000000000000000000000000000000000000000000000000005" + // len(str[0]) = 5
				"68656c6c6f000000000000000000000000000000000000000000000000000000" + // str[0]
				"0000000000000000000000000000000000000000000000000000000000000006" + // len(str[1]) = 6
				"666f6f6261720000000000000000000000000000000000000000000000000000"), // str[1]
		},
		{
			"bytes32[][]",
			nil,
			[][]common.Hash{{{1}, {2}}, {{3}, {4}, {5}}},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002" + // len(array) = 2
				"0000000000000000000000000000000000000000000000000000000000000040" + // offset 64 to i = 0
				"00000000000000000000000000000000000000000000000000000000000000a0" + // offset 160 to i = 1
				"0000000000000000000000000000000000000000000000000000000000000002" + // len(array[0]) = 2
				"0100000000000000000000000000000000000000000000000000000000000000" + // array[0][0]
				"0200000000000000000000000000000000000000000000000000000000000000" + // array[0][1]
				"0000000000000000000000000000000000000000000000000000000000000003" + // len(array[1]) = 3
				"0300000000000000000000000000000000000000000000000000000000000000" + // array[1][0]
				"0400000000000000000000000000000000000000000000000000000000000000" + // array[1][1]
				"0500000000000000000000000000000000000000000000000000000000000000"), // array[1][2]
		},

		{
			"bytes32[][2]",
			nil,
			[][]common.Hash{{{1}, {2}}, {{3}, {4}, {5}}},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000040" + // offset 64 to i = 0
				"00000000000000000000000000000000000000000000000000000000000000a0" + // offset 160 to i = 1
				"0000000000000000000000000000000000000000000000000000000000000002" + // len(array[0]) = 2
				"0100000000000000000000000000000000000000000000000000000000000000" + // array[0][0]
				"0200000000000000000000000000000000000000000000000000000000000000" + // array[0][1]
				"0000000000000000000000000000000000000000000000000000000000000003" + // len(array[1]) = 3
				"0300000000000000000000000000000000000000000000000000000000000000" + // array[1][0]
				"0400000000000000000000000000000000000000000000000000000000000000" + // array[1][1]
				"0500000000000000000000000000000000000000000000000000000000000000"), // array[1][2]
		},

		{
			"bytes32[3][2]",
			nil,
			[][]common.Hash{{{1}, {2}, {3}}, {{3}, {4}, {5}}},
			common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000" + // array[0][0]
				"0200000000000000000000000000000000000000000000000000000000000000" + // array[0][1]
				"0300000000000000000000000000000000000000000000000000000000000000" + // array[0][2]
				"0300000000000000000000000000000000000000000000000000000000000000" + // array[1][0]
				"0400000000000000000000000000000000000000000000000000000000000000" + // array[1][1]
				"0500000000000000000000000000000000000000000000000000000000000000"), // array[1][2]
		},
		{
			// static tuple
			"tuple",
			[]ArgumentMarshaling{
				{Name: "a", Type: "int64"},
				{Name: "b", Type: "int256"},
				{Name: "c", Type: "int256"},
				{Name: "d", Type: "bool"},
				{Name: "e", Type: "bytes32[3][2]"},
			},
			struct {
				A int64
				B *big.Int
				C *big.Int
				D bool
				E [][]common.Hash
			}{1, big.NewInt(1), big.NewInt(-1), true, [][]common.Hash{{{1}, {2}, {3}}, {{3}, {4}, {5}}}},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001" + // struct[a]
				"0000000000000000000000000000000000000000000000000000000000000001" + // struct[b]
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" + // struct[c]
				"0000000000000000000000000000000000000000000000000000000000000001" + // struct[d]
				"0100000000000000000000000000000000000000000000000000000000000000" + // struct[e] array[0][0]
				"0200000000000000000000000000000000000000000000000000000000000000" + // struct[e] array[0][1]
				"0300000000000000000000000000000000000000000000000000000000000000" + // struct[e] array[0][2]
				"0300000000000000000000000000000000000000000000000000000000000000" + // struct[e] array[1][0]
				"0400000000000000000000000000000000000000000000000000000000000000" + // struct[e] array[1][1]
				"0500000000000000000000000000000000000000000000000000000000000000"), // struct[e] array[1][2]
		},
		{
			// dynamic tuple
			"tuple",
			[]ArgumentMarshaling{
				{Name: "a", Type: "string"},
				{Name: "b", Type: "int64"},
				{Name: "c", Type: "bytes"},
				{Name: "d", Type: "string[]"},
				{Name: "e", Type: "int256[]"},
				{Name: "f", Type: "address[]"},
			},
			struct {
				FieldA string `abi:"a"` // Test whether abi tag works
				FieldB int64  `abi:"b"`
				C      []byte
				D      []string
				E      []*big.Int
				F      []common.Address
			}{"foobar", 1, []byte{1}, []string{"foo", "bar"}, []*big.Int{big.NewInt(1), big.NewInt(-1)}, []common.Address{{1}, {2}}},
			common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000000000c0" + // struct[a] offset
				"0000000000000000000000000000000000000000000000000000000000000001" + // struct[b]
				"0000000000000000000000000000000000000000000000000000000000000100" + // struct[c] offset
				"0000000000000000000000000000000000000000000000000000000000000140" + // struct[d] offset
				"0000000000000000000000000000000000000000000000000000000000000220" + // struct[e] offset
				"0000000000000000000000000000000000000000000000000000000000000280" + // struct[f] offset
				"0000000000000000000000000000000000000000000000000000000000000006" + // struct[a] length
				"666f6f6261720000000000000000000000000000000000000000000000000000" + // struct[a] "foobar"
				"0000000000000000000000000000000000000000000000000000000000000001" + // struct[c] length
				"0100000000000000000000000000000000000000000000000000000000000000" + // []byte{1}
				"0000000000000000000000000000000000000000000000000000000000000002" + // struct[d] length
				"0000000000000000000000000000000000000000000000000000000000000040" + // foo offset
				"0000000000000000000000000000000000000000000000000000000000000080" + // bar offset
				"0000000000000000000000000000000000000000000000000000000000000003" + // foo length
				"666f6f0000000000000000000000000000000000000000000000000000000000" + // foo
				"0000000000000000000000000000000000000000000000000000000000000003" + // bar offset
				"6261720000000000000000000000000000000000000000000000000000000000" + // bar
				"0000000000000000000000000000000000000000000000000000000000000002" + // struct[e] length
				"0000000000000000000000000000000000000000000000000000000000000001" + // 1
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" + // -1
				"0000000000000000000000000000000000000000000000000000000000000002" + // struct[f] length
				"0000000000000000000000000100000000000000000000000000000000000000" + // common.Address{1}
				"0000000000000000000000000200000000000000000000000000000000000000"), // common.Address{2}
		},
		{
			// nested tuple
			"tuple",
			[]ArgumentMarshaling{
				{Name: "a", Type: "tuple", Components: []ArgumentMarshaling{{Name: "a", Type: "uint256"}, {Name: "b", Type: "uint256[]"}}},
				{Name: "b", Type: "int256[]"},
			},
			struct {
				A struct {
					FieldA *big.Int `abi:"a"`
					B      []*big.Int
				}
				B []*big.Int
			}{
				A: struct {
					FieldA *big.Int `abi:"a"` // Test whether abi tag works for nested tuple
					B      []*big.Int
				}{big.NewInt(1), []*big.Int{big.NewInt(1), big.NewInt(0)}},
				B: []*big.Int{big.NewInt(1), big.NewInt(0)}},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000040" + // a offset
				"00000000000000000000000000000000000000000000000000000000000000e0" + // b offset
				"0000000000000000000000000000000000000000000000000000000000000001" + // a.a value
				"0000000000000000000000000000000000000000000000000000000000000040" + // a.b offset
				"0000000000000000000000000000000000000000000000000000000000000002" + // a.b length
				"0000000000000000000000000000000000000000000000000000000000000001" + // a.b[0] value
				"0000000000000000000000000000000000000000000000000000000000000000" + // a.b[1] value
				"0000000000000000000000000000000000000000000000000000000000000002" + // b length
				"0000000000000000000000000000000000000000000000000000000000000001" + // b[0] value
				"0000000000000000000000000000000000000000000000000000000000000000"), // b[1] value
		},
		{
			// tuple slice
			"tuple[]",
			[]ArgumentMarshaling{
				{Name: "a", Type: "int256"},
				{Name: "b", Type: "int256[]"},
			},
			[]struct {
				A *big.Int
				B []*big.Int
			}{
				{big.NewInt(-1), []*big.Int{big.NewInt(1), big.NewInt(0)}},
				{big.NewInt(1), []*big.Int{big.NewInt(2), big.NewInt(-1)}},
			},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002" + // tuple length
				"0000000000000000000000000000000000000000000000000000000000000040" + // tuple[0] offset
				"00000000000000000000000000000000000000000000000000000000000000e0" + // tuple[1] offset
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" + // tuple[0].A
				"0000000000000000000000000000000000000000000000000000000000000040" + // tuple[0].B offset
				"0000000000000000000000000000000000000000000000000000000000000002" + // tuple[0].B length
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[0].B[0] value
				"0000000000000000000000000000000000000000000000000000000000000000" + // tuple[0].B[1] value
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[1].A
				"0000000000000000000000000000000000000000000000000000000000000040" + // tuple[1].B offset
				"0000000000000000000000000000000000000000000000000000000000000002" + // tuple[1].B length
				"0000000000000000000000000000000000000000000000000000000000000002" + // tuple[1].B[0] value
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), // tuple[1].B[1] value
		},
		{
			// static tuple array
			"tuple[2]",
			[]ArgumentMarshaling{
				{Name: "a", Type: "int256"},
				{Name: "b", Type: "int256"},
			},
			[2]struct {
				A *big.Int
				B *big.Int
			}{
				{big.NewInt(-1), big.NewInt(1)},
				{big.NewInt(1), big.NewInt(-1)},
			},
			common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" + // tuple[0].a
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[0].b
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[1].a
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), // tuple[1].b
		},
		{
			// dynamic tuple array
			"tuple[2]",
			[]ArgumentMarshaling{
				{Name: "a", Type: "int256[]"},
			},
			[2]struct {
				A []*big.Int
			}{
				{[]*big.Int{big.NewInt(-1), big.NewInt(1)}},
				{[]*big.Int{big.NewInt(1), big.NewInt(-1)}},
			},
			common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000040" + // tuple[0] offset
				"00000000000000000000000000000000000000000000000000000000000000c0" + // tuple[1] offset
				"0000000000000000000000000000000000000000000000000000000000000020" + // tuple[0].A offset
				"0000000000000000000000000000000000000000000000000000000000000002" + // tuple[0].A length
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" + // tuple[0].A[0]
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[0].A[1]
				"0000000000000000000000000000000000000000000000000000000000000020" + // tuple[1].A offset
				"0000000000000000000000000000000000000000000000000000000000000002" + // tuple[1].A length
				"0000000000000000000000000000000000000000000000000000000000000001" + // tuple[1].A[0]
				"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), // tuple[1].A[1]
		},
	} {
		typ, err := NewType(test.typ, test.components)
		if err != nil {
			t.Fatalf("%v failed. Unexpected parse error: %v", i, err)
		}
		output, err := typ.pack(reflect.ValueOf(test.input))
		if err != nil {
			t.Fatalf("%v failed. Unexpected pack error: %v", i, err)
		}

		if !bytes.Equal(output, test.output) {
			t.Errorf("input %d for typ: %v failed. Expected bytes: '%x' Got: '%x'", i, typ.String(), test.output, output)
		}
	}
}

func TestMethodPack(t *testing.T) {
	abi, err := JSON(strings.NewReader(jsondata2))
	if err != nil {
		t.Fatal(err)
	}

	sig := abi.Methods["slice"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)

	packed, err := abi.Pack("slice", []uint32{1, 2})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	var addrA, addrB = common.Address{1}, common.Address{2}
	sig = abi.Methods["sliceAddress"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{32}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes(addrA[:], 32)...)
	sig = append(sig, common.LeftPadBytes(addrB[:], 32)...)

	packed, err = abi.Pack("sliceAddress", []common.Address{addrA, addrB})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	var addrC, addrD = common.Address{3}, common.Address{4}
	sig = abi.Methods["sliceMultiAddress"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{64}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{160}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes(addrA[:], 32)...)
	sig = append(sig, common.LeftPadBytes(addrB[:], 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes(addrC[:], 32)...)
	sig = append(sig, common.LeftPadBytes(addrD[:], 32)...)

	packed, err = abi.Pack("sliceMultiAddress", []common.Address{addrA, addrB}, []common.Address{addrC, addrD})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	sig = abi.Methods["slice256"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)

	packed, err = abi.Pack("slice256", []*big.Int{big.NewInt(1), big.NewInt(2)})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	a := [2][2]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(2), big.NewInt(0)}}
	sig = abi.Methods["nestedArray"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0xa0}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes(addrC[:], 32)...)
	sig = append(sig, common.LeftPadBytes(addrD[:], 32)...)
	packed, err = abi.Pack("nestedArray", a, []common.Address{addrC, addrD})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	sig = abi.Methods["nestedArray2"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{0x20}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0x40}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0x80}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	packed, err = abi.Pack("nestedArray2", [2][]uint8{{1}, {1}})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}

	sig = abi.Methods["nestedSlice"].Id()
	sig = append(sig, common.LeftPadBytes([]byte{0x20}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0x02}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0x40}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{0xa0}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{1}, 32)...)
	sig = append(sig, common.LeftPadBytes([]byte{2}, 32)...)
	packed, err = abi.Pack("nestedSlice", [][]uint8{{1, 2}, {1, 2}})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(packed, sig) {
		t.Errorf("expected %x got %x", sig, packed)
	}
}

func TestPackNumber(t *testing.T) {
	tests := []struct {
		value  reflect.Value
		packed []byte
	}{
		// Protocol limits
		{reflect.ValueOf(0), common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000")},
		{reflect.ValueOf(1), common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")},
		{reflect.ValueOf(-1), common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")},

		// Type corner cases
		{reflect.ValueOf(uint8(math.MaxUint8)), common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000000000ff")},
		{reflect.ValueOf(uint16(math.MaxUint16)), common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000ffff")},
		{reflect.ValueOf(uint32(math.MaxUint32)), common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000ffffffff")},
		{reflect.ValueOf(uint64(math.MaxUint64)), common.Hex2Bytes("000000000000000000000000000000000000000000000000ffffffffffffffff")},

		{reflect.ValueOf(int8(math.MaxInt8)), common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000007f")},
		{reflect.ValueOf(int16(math.MaxInt16)), common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000007fff")},
		{reflect.ValueOf(int32(math.MaxInt32)), common.Hex2Bytes("000000000000000000000000000000000000000000000000000000007fffffff")},
		{reflect.ValueOf(int64(math.MaxInt64)), common.Hex2Bytes("0000000000000000000000000000000000000000000000007fffffffffffffff")},

		{reflect.ValueOf(int8(math.MinInt8)), common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80")},
		{reflect.ValueOf(int16(math.MinInt16)), common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8000")},
		{reflect.ValueOf(int32(math.MinInt32)), common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffff80000000")},
		{reflect.ValueOf(int64(math.MinInt64)), common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffff8000000000000000")},
	}
	for i, tt := range tests {
		packed := packNum(tt.value)
		if !bytes.Equal(packed, tt.packed) {
			t.Errorf("test %d: pack mismatch: have %x, want %x", i, packed, tt.packed)
		}
	}
}
