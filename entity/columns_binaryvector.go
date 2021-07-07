// Code generated by go generate; DO NOT EDIT
// This file is generated by go genrated at 2021-07-05 22:52:01.455014639 &#43;0800 CST m=&#43;0.001674288

//package entity defines entities used in sdk
package entity 

import "github.com/milvus-io/milvus-sdk-go/internal/proto/schema"


// columnBinaryVector generated columns type for BinaryVector
type columnBinaryVector struct {
	name   string
	dim    int
	values [][]byte
}

func (c *columnBinaryVector) Name() string {
	return c.name
}

func (c *columnBinaryVector) Type() FieldType {
	return FieldTypeBinaryVector
}

func (c *columnBinaryVector) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		FieldName: c.name,
	}

	data := make([]byte, 0, len(c.values)* c.dim)

	for _, vector := range c.values {
		data = append(data, vector...)
	}

	fd.Field = &schema.FieldData_Vectors{
		Vectors: &schema.VectorField{
			Dim: int64(c.dim),
			
			Data: &schema.VectorField_BinaryVector{
				BinaryVector: data,
			},
			
		},
	}
	return fd
}

func NewColumnBinaryVector(name string,dim int, values [][]byte) Column {
	return &columnBinaryVector {
		name:   name,
		dim:    dim,
		values: values,
	}
}