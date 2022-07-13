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

type Float64Flag struct {
	Value *float64
}

func (f Float64Flag) String() string {
	if f.Value != nil {
		return strconv.FormatFloat(*f.Value, 'G', -1, 64)
	}
	return ""
}

func (f *Float64Flag) Set(v string) error {
	if f.Value == nil {
		f.Value = new(float64)
	}
	var err error
	*f.Value, err = strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	return nil
}
