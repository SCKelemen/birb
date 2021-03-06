package main 



/*
đŖ	:Hatching Chick
đĨ	:Front-Facing Baby Chick:			
đĒļ	:feather
đ	:Poultry Leg:	
*/

import "strconv"

type TokenKind int

type Token struct {
	TokenKind TokenKind
	Literal   string
}
const (
	ILLEGAL TokenKind = iota
	EOF 

	
	BIRD // đĻ
	DODO // đĻ¤
	PARROT // đĻ
	CHICKEN // đ
	EAGLE // đĻ
	CHICK // đ¤
	TURKEY // đĻ
	PENGUIN // đ§
	OWL // đĻ
	PEACOCK // đĻ
	DOVE // đī¸
	DUCK // đĻ
	ROOSTER // đ
	SWAN // đĻĸ
	FLAMINGO // đĻŠ

)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	BIRD:	"đĻ", 
	DODO:	"đĻ¤", 
	PARROT:	"đĻ", 
	CHICKEN:	"đ", 
	EAGLE:	"đĻ", 
	CHICK:	"đ¤", 
	TURKEY:	"đĻ", 
	PENGUIN:	"đ§", 
	OWL:	"đĻ", 
	PEACOCK:	"đĻ", 
	DOVE:	"đī¸", 
	DUCK:	"đĻ", 
	ROOSTER:	"đ", 
	SWAN:	"đĻĸ", 
	FLAMINGO:	"đĻŠ", 
}

func (token TokenKind) String() string {
	s := ""
	if 0 <= token && token < TokenKind(len(tokens)) {
		s = tokens[token]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(token)) + ")"
	}
	return s
}
