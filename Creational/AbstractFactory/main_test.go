package main

import (
	"testing"
)

func GetShoeDetails(s Shoe)(string,int){
	return s.GetLogo(),s.GetSize()
}

func GetShortDetails(s Short)(string,int){
	return s.GetLogo(),s.GetSize()
}


func TestSportFactory(t *testing.T) {
	t.Run("Test GetsportFactory",func(t *testing.T){
		adidasFactory,_ := GetSportFactory("adidas")
		var wont = &Adidas{}
		got := adidasFactory
		assertSportFactory(t,wont,got)
	})

	t.Run("Make shoe of adidas",func(t *testing.T){
		adidasFactory,_ := GetSportFactory("adidas")
		adidasShoe := adidasFactory.MakeShoe()
		wontLogo := "adidas"
		wontSize := 14
		gotShoeLogo,goShoeSize := GetShoeDetails(adidasShoe)

		assertShoeLogo(t,wontLogo,gotShoeLogo)
		assertShoeSize(t,wontSize,goShoeSize)
	})
	t.Run("Make short of adidas",func(t *testing.T){
		adidasFactory,_ := GetSportFactory("adidas")
		adidasShoe := adidasFactory.MakeShort()
		wontLogo := "adidas"
		wontSize := 14
		gotShortLogo,goShortSize := GetShortDetails(adidasShoe)

		assertShortLogo(t,wontLogo,gotShortLogo)
		assertShortSize(t,wontSize,goShortSize)
	})
}

func assertSportFactory(t *testing.T,wont SportFactory,got SportFactory){
	if wont!=got{
		t.Error("expected",wont,"got",got)
	}
}

func assertShoeLogo(t *testing.T,wont string,got string){
	t.Helper()
	if wont!=got{
		t.Errorf("expected %v got %v",wont,got)
	}
}

func assertShoeSize(t *testing.T,wont int,got int) {
	t.Helper()
	if wont != got {
		t.Errorf("expected %v got %v", wont, got)
	}
}

func assertShortLogo(t *testing.T,wont string,got string){
	t.Helper()
	if wont!=got{
		t.Errorf("expected %v got %v",wont,got)
	}
}

func assertShortSize(t *testing.T,wont int,got int){
	t.Helper()
	if wont!=got{
		t.Errorf("expected %v got %v",wont,got)
	}
}