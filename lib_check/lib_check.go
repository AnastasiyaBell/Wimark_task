package lib_check

//import "fmt"
/*
lookup_count is a counter for every word in a newspaper
if some word is not in the newspaper (or maybe the word was used completely if it has duplicates) we have to receive impossibility (false)
*/

func Check (paper[] string, piece[] string ) bool{
	var flag bool = true
	lookup_count := make(map[string]int)

	for i := 0; i < len(paper); i++ {
		lookup_count[paper[i]] = lookup_count[paper[i]] + 1
		//fmt.Println(lookup_count[paper[i]])
	}

	for i := 0; i < len(piece); i++ {
		if lookup_count[piece[i]]-1 < 0 {
                        //fmt.Println(lookup_count[piece[i]])
			//fmt.Println(piece[i])
			flag = false
			break
		}
		lookup_count[piece[i]] = lookup_count[piece[i]] - 1
	}
	
	return flag
}
