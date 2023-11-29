package main

import "strconv"

type StringFlag struct {
	Value *string
}

func (s StringFlag) String() string {
	if s.Value != nil {
		return *s.Value
	}
	return ""
}

func (f *StringFlag) Set(v string) error {
	if f.Value == nil {
		f.Value = new(string)
	}
	*f.Value = v
	return nil
}

type BoolFlag struct {
	Value *bool
}

func (f BoolFlag) String() string {
	if f.Value != nil {
		return strconv.FormatBool(*f.Value)
	}
	return ""
}

func (f *BoolFlag) Set(v string) error {
	if f.Value == nil {
		f.Value = new(bool)
	}
	var err error
	*f.Value, err = strconv.ParseBool(v)
	if err != nil {
		return err
	}
	return nil
}

func (f BoolFlag) IsBoolFlag() bool { return true }

type Float32Flag struct {
	Value *float32
}

func (f Float32Flag) String() string {
	if f.Value != nil {
		return strconv.FormatFloat(float64(*f.Value), 'G', -1, 32)
	}
	return ""
}

func (f *Float32Flag) Set(v string) error {
	if f.Value == nil {
		f.Value = new(float32)
	}
	var err error
	var value float64
	value, err = strconv.ParseFloat(v, 32)
	if err != nil {
		return err
	}
	*f.Value = float32(value)
	return nil
}
