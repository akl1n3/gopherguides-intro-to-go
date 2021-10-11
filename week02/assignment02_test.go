package main

func main(){
        exp := []string{"John", "Paul", "George", "Ringo"}
        var act []string

        // Iterate through the exp variable and copy its values into act var
        for i := 0; i < len(exp); i++ {
                copy(act, exp)
        }

        // Iterate through the act var and assert that its contents match that
        // of the exp var
	func testEq(act, exp []string) bool {
    		if len(act) != len(exp) {
        		return false
    		}
    		for i := range a {
        		if act[i] != exp[i] {
            			return false
        		}	
    		}
    		return true
	}

}
