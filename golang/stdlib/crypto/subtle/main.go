package main

import (
	"crypto/subtle"
	"log"
)

func main() {
	for _, tc := range []struct {
		A byte
		B byte
	}{
		{'a', 'a'}, // 1
		{'a', 'b'}, // 0
		{'b', 'a'}, // 0
	} {
		r := subtle.ConstantTimeByteEq(tc.A, tc.B)
		log.Println(r)
	}

	for _, tc := range []struct {
		A []byte
		B []byte
	}{
		{[]byte("A"), []byte("A")}, // 1
		{[]byte("A"), []byte("B")}, // 0
		{[]byte("B"), []byte("A")}, // 0
	} {
		r := subtle.ConstantTimeCompare(tc.A, tc.B)
		log.Println(r)
	}

	for _, tc := range []struct {
		A []byte
		B []byte
	}{
		{[]byte("A"), []byte("A")}, // {[65] [65]}
		{[]byte("A"), []byte("B")}, // {[66] [66]}
		{[]byte("B"), []byte("A")}, // {[65] [65]}
	} {
		subtle.ConstantTimeCopy(1, tc.A, tc.B)
		log.Println(tc)
	}

	for _, tc := range []struct {
		A []byte
		B []byte
	}{
		{[]byte("A"), []byte("A")}, // {[65] [65]}
		{[]byte("A"), []byte("B")}, // {[65] [66]}
		{[]byte("B"), []byte("A")}, // {[66] [65]}
	} {
		subtle.ConstantTimeCopy(0, tc.A, tc.B)
		log.Println(tc)
	}

	for _, tc := range []struct {
		A int32
		B int32
	}{
		{1, 1}, // 1
		{1, 2}, // 0
		{2, 1}, // 0
	} {
		r := subtle.ConstantTimeEq(tc.A, tc.B)
		log.Println(r)
	}

	for _, tc := range []struct {
		A int
		B int
	}{
		{1, 1}, // 1
		{1, 2}, // 1
		{2, 1}, // 0
	} {
		r := subtle.ConstantTimeLessOrEq(tc.A, tc.B)
		log.Println(r)
	}

	for _, tc := range []struct {
		A int
		B int
	}{
		{1, 1}, // 1
		{1, 2}, // 1
		{2, 1}, // 2
	} {
		r := subtle.ConstantTimeSelect(1, tc.A, tc.B)
		log.Println(r)
	}

	for _, tc := range []struct {
		A int
		B int
	}{
		{1, 1}, // 1
		{1, 2}, // 2
		{2, 1}, // 1
	} {
		r := subtle.ConstantTimeSelect(0, tc.A, tc.B)
		log.Println(r)
	}
}
