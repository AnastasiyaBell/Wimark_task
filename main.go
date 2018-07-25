package main

import (
    "fmt"
    "os"
    "io"
    "strings"
    "./lib_check"
)

func main(){
    var paper_in []string
    var piece_in []string

    /* Take newspaper from "paper.txt" */   

    file1, err := os.Open("paper.txt")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file1.Close() 
    
    /* At first take strings from file as byte slice and 
    immediately cast to string. Then split on words. And delete empty strings*/
     
    data1 := make([]byte, 512)
    var paper_in1 []string
    for{
        n1, err := file1.Read(data1)
        if err == io.EOF{
            break
        }
        paper_in1 = append(paper_in1, strings.Split(string(data1[:n1]), "\n")...)
    }
    for j := 0; j < len(paper_in1); j++{
        paper_in = append(paper_in, strings.Split(paper_in1[j], " ")...)
    }
    for k := 0; ; k++{
        if k >= len(paper_in) {
            break
        }
        if len(paper_in[k]) == 0 {
            if k < len(paper_in) - 1{
                paper_in = append(paper_in[:k], paper_in[k+1:]...)
                k--
            } else{
                paper_in = paper_in[:k]
                k--
            }
        }
    }

    //for i := 0; i < len(paper_in); i++{
    //    fmt.Println("Paper word:", paper_in[i])
    //}
    
    /* Take note from "piece.txt" */
    
    file2, err := os.Open("piece.txt")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file2.Close()
    
    /* At first take strings from file as byte slice and
    immediately cast to string. Then split on words. And delete empty strings*/
    
    data2 := make([]byte, 512)
    var piece_in1 []string
    for{
        n2, err := file2.Read(data2)
        if err == io.EOF{
            break
        }
        piece_in1 = append(piece_in1, strings.Split(string(data2[:n2]), "\n")...)
        //fmt.Println(piece_in1)
        //fmt.Println(len(piece_in1))
    }
    for j := 0; j < len(piece_in1); j++{
        piece_in = append(piece_in, strings.Split(piece_in1[j], " ")...)
    }

    //fmt.Println("before deleting empty string", len(piece_in))
    
    for k := 0; ; k++{
        if k >= len(piece_in) {
            break
        }
        if len(piece_in[k]) == 0 {
            if k < len(piece_in) - 1{
                piece_in = append(piece_in[:k], piece_in[k+1:]...)
                k--
            } else{
                piece_in = piece_in[:k]
                k--
            }
        }
    }
    
    //fmt.Println("after deleting empty string", len(piece_in))

    //for i := 0; i < len(piece_in); i++{
    //    fmt.Println("Piece word:", piece_in[i])
    //}
    
    /* Find the answer if we can compose the note from the newspaper */
    
    answer := lib_check.Check(paper_in, piece_in)
    fmt.Println(answer)

}
