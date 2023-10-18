package main

import (
    "fmt"
    "math/big"
)

type LargeNumber struct {
    value *big.Int
}

func NewLargeNumber(hexValue string) *LargeNumber {
    value, success := new(big.Int).SetString(hexValue, 16)
    if !success {
        panic("Invalid hex value")
    }
    return &LargeNumber{value}
}

func (ln *LargeNumber) SetHex(hexValue string) {
    value, success := new(big.Int).SetString(hexValue, 16)
    if !success {
        panic("Invalid hex value")
    }
    ln.value.Set(value)
}

func (ln *LargeNumber) GetHex() string {
    return fmt.Sprintf("%x", ln.value)
}

func (ln *LargeNumber) Invert() {
    ln.value.Not(ln.value)
}

func (ln *LargeNumber) XOR(other *LargeNumber) {
    ln.value.Xor(ln.value, other.value)
}

func (ln *LargeNumber) OR(other *LargeNumber) {
    ln.value.Or(ln.value, other.value)
}

func (ln *LargeNumber) AND(other *LargeNumber) {
    ln.value.And(ln.value, other.value)
}

func (ln *LargeNumber) ShiftRight(bits int) {
    ln.value.Rsh(ln.value, uint(bits))
}

func (ln *LargeNumber) ShiftLeft(bits int) {
    ln.value.Lsh(ln.value, uint(bits))
}

func (ln *LargeNumber) Add(other *LargeNumber) {
    ln.value.Add(ln.value, other.value)
}

func (ln *LargeNumber) Subtract(other *LargeNumber) {
    ln.value.Sub(ln.value, other.value)
}

func (ln *LargeNumber) Mod(modulus *LargeNumber) {
    ln.value.Mod(ln.value, modulus.value)
}

func (ln *LargeNumber) Multiply(other *LargeNumber) {
    ln.value.Mul(ln.value, other.value)
}

func main() {
    numberA := NewLargeNumber("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
    numberB := NewLargeNumber("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
    numberC := NewLargeNumber("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
    numberD := NewLargeNumber("70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb")
    numberE := NewLargeNumber("33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc")
    numberF := NewLargeNumber("22e962951cb6cd2ce279ab0e2095825c141d48ef3ca9dabf253e38760b57fe03")
    numberK := NewLargeNumber("7d7deab2affa38154326e96d350deee1")
    numberN := NewLargeNumber("97f92a75b3faf8939e8e98b96476fd22")

    numberA.XOR(numberB)
    fmt.Println("XOR:", numberA.GetHex())

    numberC.Add(numberD)
    fmt.Println("Add:", numberC.GetHex())

    numberE.Subtract(numberF)
    fmt.Println("Sub:", numberE.GetHex())

    numberK.Multiply(numberN)
    fmt.Println("Mul:", numberK.GetHex())

}