/*
Some error handling routines
*/

package err_chk

//Panic on error
func Check(e error) {
	//Generic error test
	if e != nil {
		panic(e)
	}
}
