package main

import (
	"fmt"
	"math"
	"testing"
)

// Tolerance for checking floating point balances precisely
const balanceDelta = 1e-9

func TestBankAccount_Deposit(t *testing.T) {
	t.Run("Valid Deposits Increase Balance", func(t *testing.T) {
		account := &BankAccount{Owner: "Alice", Balance: 100.0}

		account.Deposit(50.50)
		if math.Abs(account.Balance-150.50) > balanceDelta {
			t.Errorf("After depositing 50.50, Balance = %f; want 150.50", account.Balance)
		}

		account.Deposit(25.0)
		if math.Abs(account.Balance-175.50) > balanceDelta {
			t.Errorf("After depositing another 25.00, Balance = %f; want 175.50", account.Balance)
		}
	})
}

func TestBankAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name        string
		initBalance float64
		amount      float64
		wantBalance float64
		expectErr   bool
	}{
		{
			name:        "Successful exact withdrawal",
			initBalance: 100.0,
			amount:      100.0,
			wantBalance: 0.0,
			expectErr:   false,
		},
		{
			name:        "Successful partial withdrawal",
			initBalance: 250.25,
			amount:      50.25,
			wantBalance: 200.0,
			expectErr:   false,
		},
		{
			name:        "Fails due to insufficient funds",
			initBalance: 50.0,
			amount:      100.0,
			wantBalance: 50.0, // Balance must remain untouched on failure
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &BankAccount{Owner: "Bob", Balance: tt.initBalance}
			err := account.Withdraw(tt.amount)

			// Verify error behavior
			if tt.expectErr && err == nil {
				t.Errorf("%s: expected an insufficient funds error, but got nil", tt.name)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("%s: unexpected error returned: %v", tt.name, err)
			}

			// Verify final state of the account balance
			if math.Abs(account.Balance-tt.wantBalance) > balanceDelta {
				t.Errorf("%s: Balance = %f; want %f", tt.name, account.Balance, tt.wantBalance)
			}
		})
	}
}

func TestBankAccount_String(t *testing.T) {
	// Verifies compliance with the fmt.Stringer interface using pointer instances
	tests := []struct {
		name       string
		account    *BankAccount
		wantOutput string
	}{
		{
			name:       "Standard formatting with decimal place padding",
			account:    &BankAccount{Owner: "Alice", Balance: 250.0},
			wantOutput: "Alice: $250.00",
		},
		{
			name:       "Formatting with fractional change values",
			account:    &BankAccount{Owner: "Charlie", Balance: 123.456},
			wantOutput: "Charlie: $123.46", // Fmt.Sprintf rounding verification
		},
		{
			name:       "Zero value representation",
			account:    &BankAccount{Owner: "Dave", Balance: 0.0},
			wantOutput: "Dave: $0.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := tt.account.String()
			if gotOutput != tt.wantOutput {
				t.Errorf("%s: String() = %q; want %q", tt.name, gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestStringerInterfaceAssignment(t *testing.T) {
	t.Run("Pointer strictly satisfies fmt.Stringer", func(t *testing.T) {
		// Validates that *BankAccount satisfies Go's standard core formatting contract
		var _ fmt.Stringer = &BankAccount{}
	})
}
