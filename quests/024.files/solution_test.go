package files
import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the files Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestProcessFile(t *testing.T){
	tests:=[]struct{
		name string
		input string
		expectedError bool
		expectedOut string
	}{
		{
			name:          "less_than_10_words",
			input:         "only four words here",
			expectedError: true,
			expectedOut:   "",
		},
		{
			name:          "exactly_10_words",
			input:         "the quick brown fox jumps over the lazy sleeping dog",
			expectedError: false,
			expectedOut:   "THE QUICK BROWN FOX JUMPS over the lazy sleeping dog",
		},
		{
			name:          "more_than_10_words_middle_unchanged",
			input:         "the quick brown fox jumps right over the lazy sleeping dog",
			expectedError: false,
			expectedOut:   "THE QUICK BROWN FOX JUMPS right over the lazy sleeping dog",
		},
	}
	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T){
			os.Remove("task.txt");
			err:=os.WriteFile("task.txt",[]byte(tt.input),0644)
			if err!=nil{
				t.Fatal(err)
			}
			defer os.Remove("task.txt")
			err=ProcessFile()
			if tt.expectedError{
				if err==nil{
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err!=nil{
				t.Errorf("expected nil but got %v",err)
				return
			}
			raw,err:=os.ReadFile("task.txt")
			if err!=nil{
				t.Errof("could not read task.txt after calling ProcessFile()")
				return
			}
			got:=strings.TrimSpace(string(raw))
			want:=tt.expectedOut
			if got!=want{
				t.Errorf("expected %v but got %v",want,got)
			}
		})
	}

}