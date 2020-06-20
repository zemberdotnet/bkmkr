package main

type Forms []FormResp

type FormResp struct {
	ID    string
	Born  int64
	First string
	Last  string
}
