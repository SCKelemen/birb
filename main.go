package main 



/*
🐣	:Hatching Chick
🐥	:Front-Facing Baby Chick:			
🪶	:feather
🍗	:Poultry Leg:	
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

	
	BIRD // 🐦
	DODO // 🦤
	PARROT // 🦜
	CHICKEN // 🐔
	EAGLE // 🦅
	CHICK // 🐤
	TURKEY // 🦃
	PENGUIN // 🐧
	OWL // 🦉
	PEACOCK // 🦚
	DOVE // 🕊️
	DUCK // 🦆
	ROOSTER // 🐓
	SWAN // 🦢
	FLAMINGO // 🦩

)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	BIRD:	"🐦", 
	DODO:	"🦤", 
	PARROT:	"🦜", 
	CHICKEN:	"🐔", 
	EAGLE:	"🦅", 
	CHICK:	"🐤", 
	TURKEY:	"🦃", 
	PENGUIN:	"🐧", 
	OWL:	"🦉", 
	PEACOCK:	"🦚", 
	DOVE:	"🕊️", 
	DUCK:	"🦆", 
	ROOSTER:	"🐓", 
	SWAN:	"🦢", 
	FLAMINGO:	"🦩", 
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
