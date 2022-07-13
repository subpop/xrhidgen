package xrhidgen

func ptrbool(b bool) *bool                { return &b }
func ptrfloat64(f float64) *float64       { return &f }
func ptrstring(s string) *string          { return &s }
func ptrslicestring(s []string) *[]string { return &s }
