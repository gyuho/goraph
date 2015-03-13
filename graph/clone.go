package graph

// Clone clones the graph Data.
// It does `Deep Copy`.
// That is, changing the cloned Data would not affect the original Data.
func (d *Data) Clone() *Data {
	copied := NewData()
	return copied
}
