package grater

import (
	"log"
	"errors"
	)

type outFunc func(*Grater, string) error

func logOut(g *Grater, s string) error{
	if g == nil{
		err := errors.New("Grater is nil")
		return err
	}

	log.Println("Grater:", s)
	return nil
}

func writerOut(g *Grater, s string) error{
	if g == nil{
		err := errors.New("Grater is nil")
		return err
	}
	if g.writer == nil{
		err := errors.New("Writer is nil")
		return err
	}
	
	_, err := g.writer.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}
