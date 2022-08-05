package helper

import "testing"
func TestHelloWorld(t *testing.T){
	result := HelloWorld("res")
	if result!= "Hello dino"{
		// unit test failed
		// panic("result is not 'Hello dino'")
		t.FailNow()
	}
}
func TestHelloWorldJC(t *testing.T){
	result := HelloWorld("jc")
	if result!= "Hello jc"{
		// unit test failed
		
		t.Fail()
	}
}