package keyboards

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestCreateKb(t *testing.T) {
	m := ListCancel
	replyMk := m.CreateKB()
	rm := &replyMk
	fmt.Printf("len is %v", len(rm.InlineKeyboard))
	//fmt.Printf("%v", rm.InlineKeyboard[1][1])
	//fmt.Printf("%v", rm.ReplyKeyboard[1][1])
	//fmt.Printf("%v", rm.InlineKeyboard[2][2])
	//fmt.Printf("%v", rm.ReplyKeyboard[2][2])
	for _, item := range rm.InlineKeyboard {
		for j, k := range item {
			//fmt.Printf("\nindex = %v, item = %v", idx, item)
			fmt.Printf("\nindex = %v, text = %v, data = %v", j, k.Text, k.Data)
		}
	}

	activeCommand := "/search_subs_by_title"
	index := strings.Index(activeCommand, "/") // var mode yt.CompareMode
	if index != 0 {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( command is: '%v', index is: '%v', program counter is %v, file and line is %v, %v ) ",
			activeCommand, index, programCounter, file, line)
	}
}

func TestCreateBtnRow(t *testing.T) {
	replyMk := CreateBtnRow(SearchTargetOptions)
	rm := replyMk
	fmt.Printf("len is %v", len(rm))
	//fmt.Printf("%v", rm.InlineKeyboard[1][1])
	//fmt.Printf("%v", rm.ReplyKeyboard[1][1])
	//fmt.Printf("%v", rm.InlineKeyboard[2][2])
	//fmt.Printf("%v", rm.ReplyKeyboard[2][2])
	for _, item := range rm {
		fmt.Printf("text = %v, data = %v", item.Text, item.Data)
	}

	activeCommand := "/search_subs_by_title"
	index := strings.Index(activeCommand, "/") // var mode yt.CompareMode
	if index != 0 {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( command is: '%v', index is: '%v', program counter is %v, file and line is %v, %v ) ",
			activeCommand, index, programCounter, file, line)
	}
}
