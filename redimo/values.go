package redimo

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"math/big"
)

type Value interface {
	AsBytes() ([]byte, bool)
	AsString() (string, bool)
	AsNumeric() (*big.Float, bool)
	toAV() dynamodb.AttributeValue
}

type BytesValue struct {
	bytes []byte
}

func (bv BytesValue) AsBytes() (out []byte, ok bool)       { return bv.bytes, true }
func (bv BytesValue) AsString() (out string, ok bool)      { return }
func (bv BytesValue) AsNumeric() (out *big.Float, ok bool) { return }
func (bv BytesValue) toAV() dynamodb.AttributeValue {
	return dynamodb.AttributeValue{
		B: bv.bytes,
	}
}

type StringValue struct {
	str string
}

func (sv StringValue) AsBytes() (out []byte, ok bool)       { return }
func (sv StringValue) AsString() (out string, ok bool)      { return sv.str, true }
func (sv StringValue) AsNumeric() (out *big.Float, ok bool) { return }
func (sv StringValue) toAV() dynamodb.AttributeValue {
	return dynamodb.AttributeValue{
		S: aws.String(sv.str),
	}
}

type NumericValue struct {
	bf *big.Float
}

func (nv NumericValue) AsBytes() (out []byte, ok bool)       { return }
func (nv NumericValue) AsString() (out string, ok bool)      { return }
func (nv NumericValue) AsNumeric() (out *big.Float, ok bool) { return nv.bf, true }
func (nv NumericValue) toAV() dynamodb.AttributeValue {
	return dynamodb.AttributeValue{
		N: aws.String(nv.bf.String()),
	}
}