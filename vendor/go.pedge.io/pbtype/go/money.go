package pbtype

import (
	"fmt"

	"go.pedge.io/googleapis/google/type"
)

// NewMoney returns a new Money for the given CurrencyCode and valueMicros.
func NewMoney(currencyCode CurrencyCode, valueMicros int64) *Money {
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  valueMicros,
	}
}

// NewMoneyUSD returns a new Money for USD for the given valueMicros.
func NewMoneyUSD(valueMicros int64) *Money {
	return NewMoney(CurrencyCodeUSD, valueMicros)
}

// NewMoneyEUR returns a new Money for EUR for the given valueMicros.
func NewMoneyEUR(valueMicros int64) *Money {
	return NewMoney(CurrencyCodeEUR, valueMicros)
}

// NewMoneySimple returns a new Money for the given CurrencyCode and value.
//
// ValueMinorUnits will use the units of the CurrencyCode.
// If valueUnits is negative, valueMinorUnits will be converted to negative.
func NewMoneySimple(currencyCode CurrencyCode, valueUnits int64, valueMinorUnits int64) *Money {
	if valueUnits < 0 {
		valueMinorUnits = -valueMinorUnits
	}
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  unitsAndMicroPartToMicros(valueUnits, valueMinorUnits*currencyCode.minorMultiplier()),
	}
}

// NewMoneySimpleUSD returns a new Money for USD for the given and value.
//
// If valueDollars is negative, valueCents will be converted to negative.
func NewMoneySimpleUSD(valueDollars int64, valueCents int64) *Money {
	return NewMoneySimple(CurrencyCodeUSD, valueDollars, valueCents)
}

// NewMoneySimpleEUR returns a new Money for EUR for the given and value.
//
// If valueEuros is negative, valueCents will be converted to negative.
func NewMoneySimpleEUR(valueEuros int64, valueCents int64) *Money {
	return NewMoneySimple(CurrencyCodeEUR, valueEuros, valueCents)
}

// NewMoneyFloat returns a new Money for the given CurrencyCode and float value.
func NewMoneyFloat(currencyCode CurrencyCode, value float64) *Money {
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  floatUnitsToMicros(value),
	}
}

// NewMoneyFloatUSD returns a new Money for USD for the given and value.
func NewMoneyFloatUSD(valueDollars float64) *Money {
	return NewMoneyFloat(CurrencyCodeUSD, valueDollars)
}

// NewMoneyFloatEUR returns a new Money for EUR for the given and value.
func NewMoneyFloatEUR(valueEuros float64) *Money {
	return NewMoneyFloat(CurrencyCodeEUR, valueEuros)
}

// GoogleMoneyToMoney converts a google.type.Money to Money.
func GoogleMoneyToMoney(googleMoney *google_type.Money) (*Money, error) {
	currencyCode, err := CurrencyCodeSimpleValueOf(googleMoney.CurrencyCode)
	if err != nil {
		return nil, err
	}
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  unitsAndNanoPartToMicros(googleMoney.Units, googleMoney.Nanos),
	}, nil
}

// ToGoogleMoney converts the given Money into a google.type.Money
func (m *Money) ToGoogleMoney() *google_type.Money {
	units, nanoPart := microsToUnitsAndNanoPart(m.ValueMicros)
	return &google_type.Money{
		CurrencyCode: m.CurrencyCode.SimpleString(),
		Units:        units,
		Nanos:        nanoPart,
	}
}

// IsZero returns true if ValueMicros == 0.
func (m *Money) IsZero() bool {
	return m.ValueMicros == 0
}

// Units returns the units-only part of the value.
func (m *Money) Units() int64 {
	units, _ := microsToUnitsAndMicroPart(m.ValueMicros)
	return units
}

// SimpleString returns the simple string for the Money.
func (m *Money) SimpleString() string {
	units, microPart := microsToUnitsAndMicroPart(m.ValueMicros)
	if microPart < 0 {
		microPart = -microPart
	}
	return fmt.Sprintf("%d.%06d", units, microPart)
}

// MoneyMathable is the interface needed to perform money math operations.
type MoneyMathable interface {
	GetCurrencyCode() CurrencyCode
	GetValueMicros() int64
	Errors() []error
}

// MoneyMather performs math operations on Money.
type MoneyMather interface {
	MoneyMathable
	Plus(MoneyMathable) MoneyMather
	Minus(MoneyMathable) MoneyMather
	Times(MoneyMathable) MoneyMather
	Div(MoneyMathable) MoneyMather
	Min(MoneyMathable) MoneyMather
	Abs() MoneyMather
	PlusInt(int64) MoneyMather
	MinusInt(int64) MoneyMather
	TimesInt(int64) MoneyMather
	DivInt(int64) MoneyMather
	PlusFloat(float64) MoneyMather
	MinusFloat(float64) MoneyMather
	TimesFloat(float64) MoneyMather
	DivFloat(float64) MoneyMather
	Result() (*Money, error)
}

// Plus does the MoneyMather Plus operation.
func (m *Money) Plus(moneyMathable MoneyMathable) MoneyMather {
	return newMoneyMather(m).Plus(moneyMathable)
}

// Minus does the MoneyMather Minus operation.
func (m *Money) Minus(moneyMathable MoneyMathable) MoneyMather {
	return newMoneyMather(m).Minus(moneyMathable)
}

// Times does the MoneyMather Times operation.
func (m *Money) Times(moneyMathable MoneyMathable) MoneyMather {
	return newMoneyMather(m).Times(moneyMathable)
}

// Div does the MoneyMather Div operation.
func (m *Money) Div(moneyMathable MoneyMathable) MoneyMather {
	return newMoneyMather(m).Div(moneyMathable)
}

// Min does the MoneyMather Min operation.
func (m *Money) Min(moneyMathable MoneyMathable) MoneyMather {
	return newMoneyMather(m).Min(moneyMathable)
}

// Abs does the MoneyMather Abs operation.
func (m *Money) Abs() MoneyMather {
	return newMoneyMather(m).Abs()
}

// PlusInt does the MoneyMather Plus operation.
func (m *Money) PlusInt(value int64) MoneyMather {
	return newMoneyMather(m).PlusInt(value)
}

// MinusInt does the MoneyMather Minus operation.
func (m *Money) MinusInt(value int64) MoneyMather {
	return newMoneyMather(m).MinusInt(value)
}

// TimesInt does the MoneyMather Times operation.
func (m *Money) TimesInt(value int64) MoneyMather {
	return newMoneyMather(m).TimesInt(value)
}

// DivInt does the MoneyMather Div operation.
func (m *Money) DivInt(value int64) MoneyMather {
	return newMoneyMather(m).DivInt(value)
}

// PlusFloat does the MoneyMather Plus operation.
func (m *Money) PlusFloat(value float64) MoneyMather {
	return newMoneyMather(m).PlusFloat(value)
}

// MinusFloat does the MoneyMather Minus operation.
func (m *Money) MinusFloat(value float64) MoneyMather {
	return newMoneyMather(m).MinusFloat(value)
}

// TimesFloat does the MoneyMather Times operation.
func (m *Money) TimesFloat(value float64) MoneyMather {
	return newMoneyMather(m).TimesFloat(value)
}

// DivFloat does the MoneyMather Div operation.
func (m *Money) DivFloat(value float64) MoneyMather {
	return newMoneyMather(m).DivFloat(value)
}

// GetCurrencyCode returns the CurrencyCode.
func (m *Money) GetCurrencyCode() CurrencyCode {
	return m.CurrencyCode
}

// GetValueMicros returns the ValueMicros.
func (m *Money) GetValueMicros() int64 {
	return m.ValueMicros
}

// Errors returns nil.
func (m *Money) Errors() []error {
	return nil
}

type moneyMather struct {
	cc   CurrencyCode
	vm   int64
	errs []error
}

func newMoneyMather(money *Money) *moneyMather {
	var errs []error
	if money.CurrencyCode == CurrencyCode_CURRENCY_CODE_NONE {
		errs = append(errs, fmt.Errorf("pbtype: cannot use Money with CurrencyCode_CURRENCY_CODE_NONE"))
	}
	return &moneyMather{
		cc:   money.CurrencyCode,
		vm:   money.ValueMicros,
		errs: errs,
	}
}

func (m *moneyMather) Plus(moneyMathable MoneyMathable) MoneyMather {
	if !m.ok(moneyMathable) {
		return m
	}
	m.vm += moneyMathable.GetValueMicros()
	return m
}

func (m *moneyMather) Minus(moneyMathable MoneyMathable) MoneyMather {
	if !m.ok(moneyMathable) {
		return m
	}
	m.vm -= moneyMathable.GetValueMicros()
	return m
}

func (m *moneyMather) Times(moneyMathable MoneyMathable) MoneyMather {
	if !m.ok(moneyMathable) {
		return m
	}
	m.vm *= moneyMathable.GetValueMicros()
	return m
}

func (m *moneyMather) Div(moneyMathable MoneyMathable) MoneyMather {
	if !m.ok(moneyMathable) {
		return m
	}
	if moneyMathable.GetValueMicros() == 0 {
		m.errs = append(m.errs, fmt.Errorf("pbtype: cannot divide by 0"))
		return m
	}
	m.vm /= moneyMathable.GetValueMicros()
	return m
}

func (m *moneyMather) Min(moneyMathable MoneyMathable) MoneyMather {
	if !m.ok(moneyMathable) {
		return m
	}
	if m.vm > moneyMathable.GetValueMicros() {
		m.vm = moneyMathable.GetValueMicros()
	}
	return m
}

func (m *moneyMather) Abs() MoneyMather {
	if m.vm < 0 {
		m.vm = -m.vm
	}
	return m
}

func (m *moneyMather) PlusInt(value int64) MoneyMather {
	m.vm += value
	return m
}

func (m *moneyMather) MinusInt(value int64) MoneyMather {
	m.vm -= value
	return m
}

func (m *moneyMather) TimesInt(value int64) MoneyMather {
	m.vm *= value
	return m
}

func (m *moneyMather) DivInt(value int64) MoneyMather {
	if value == 0 {
		m.errs = append(m.errs, fmt.Errorf("pbtype: cannot divide by 0"))
		return m
	}
	m.vm /= value
	return m
}

func (m *moneyMather) PlusFloat(value float64) MoneyMather {
	m.vm += int64(value)
	return m
}

func (m *moneyMather) MinusFloat(value float64) MoneyMather {
	m.vm -= int64(value)
	return m
}

func (m *moneyMather) TimesFloat(value float64) MoneyMather {
	m.vm = int64(float64(m.vm) * value)
	return m
}

func (m *moneyMather) DivFloat(value float64) MoneyMather {
	if value == 0.0 {
		m.errs = append(m.errs, fmt.Errorf("pbtype: cannot divide by 0"))
		return m
	}
	m.vm = int64(float64(m.vm) / value)
	return m
}

func (m *moneyMather) Result() (*Money, error) {
	if len(m.errs) > 0 {
		return nil, fmt.Errorf("%v", m.errs)
	}
	return &Money{
		CurrencyCode: m.cc,
		ValueMicros:  m.vm,
	}, nil
}

func (m *moneyMather) GetCurrencyCode() CurrencyCode {
	return m.cc
}

func (m *moneyMather) GetValueMicros() int64 {
	return m.vm
}

func (m *moneyMather) Errors() []error {
	return m.errs
}

func (m *moneyMather) ok(moneyMathable MoneyMathable) bool {
	errs := moneyMathable.Errors()
	if len(errs) > 0 {
		m.errs = append(m.errs, errs...)
		return false
	}
	if m.cc != moneyMathable.GetCurrencyCode() {
		m.errs = append(m.errs, fmt.Errorf("pbtype: mismatched CurrencyCodes: %s %s", m.cc, moneyMathable.GetCurrencyCode()))
		return false
	}
	return true
}

func unitsAndMicroPartToMicros(units int64, micros int64) int64 {
	return unitsToMicros(units) + micros
}

func unitsAndNanoPartToMicros(units int64, nanos int32) int64 {
	return unitsToMicros(units) + int64(nanos/1000)
}

func microsToUnitsAndMicroPart(micros int64) (int64, int64) {
	return micros / 1000000, micros % 1000000
}

func microsToUnitsAndNanoPart(micros int64) (int64, int32) {
	return micros / 1000000, int32(micros%1000000) * 1000
}

func unitsToMicros(units int64) int64 {
	return units * 1000000
}

func floatUnitsToMicros(floatUnits float64) int64 {
	return int64(floatUnits * 1000000.0)
}