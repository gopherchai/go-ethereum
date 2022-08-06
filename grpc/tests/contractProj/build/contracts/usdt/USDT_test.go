// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdt 

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TestUSDTCaller_BalanceOf(t *testing.T) {
	type fields struct {
		contract *bind.BoundContract
	}
	type args struct {
		opts   *bind.CallOpts
		_owner common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "tbalanceOf",
			fields: fields{
				
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_USDT := &USDTCaller{
				contract: tt.fields.contract,
			}
			got, err := _USDT.BalanceOf(tt.args.opts, tt.args._owner)
			if (err != nil) != tt.wantErr {
				t.Errorf("USDTCaller.BalanceOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USDTCaller.BalanceOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
