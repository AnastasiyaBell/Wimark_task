package lib_check

//import "fmt"
/*
lookup_count is a counter for every word in a newspaper
if some word is not in the newspaper (or maybe the word was used completely if it has duplicates) we have to receive impossibility (false)
*/

func Check (paper[] string, piece[] string ) bool{
	var flag bool = true
	lookup_count := make(map[string]int)
        
        paper_ch := make(chan int)
        piece_ch := make(chan int)
        
        
       	go func () {
                for i := 0; i < len(paper); i++ {
		        lookup_count[paper[i]] = lookup_count[paper[i]] + 1
		        //fmt.Println("Paper", lookup_count[paper[i]], paper[i])
	        }
                paper_ch <- 1
        }()
        
	go func () {
                for i := 0; i < len(piece); i++ {
		        /*if lookup_count[piece[i]]-1 < 0 {
                                //fmt.Println(lookup_count[piece[i]])
			        //fmt.Println(piece[i])
			        flag = false
			        break
		        }*/
		        lookup_count[piece[i]] = lookup_count[piece[i]] - 1
                        //fmt.Println("Piece", lookup_count[piece[i]], piece[i])        
                }
                piece_ch <- 1
        }()
        
	<-paper_ch
        <-piece_ch

        for i := 0; i < len(piece); i++ {
                if lookup_count[piece[i]] < 0 {
                        flag = false
                        //fmt.Println(piece[i], lookup_count[piece[i]])
                        break
                }
        }
	return flag
}
