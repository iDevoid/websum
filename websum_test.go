package websum

import (
	"reflect"
	"testing"
)

func TestSummarizeWeb(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantSum Summary
		wantErr bool
	}{
		{
			name: "github login HTML 5",
			args: args{
				url: "https://github.com/login",
			},
			wantSum: Summary{
				HTMLVersion: "HTML 5",
				Title:       "Sign in to GitHub · GitHub",
				HeadingsCount: map[string]int{
					"h1": 1,
					"h2": 0,
					"h3": 0,
					"h4": 0,
					"h5": 0,
					"h6": 0,
				},
				LinksCount: Links{
					External:     1,
					Internal:     10,
					Inaccessable: 1,
				},
				ContainLogin: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, err := SummarizeWeb(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("SummarizeWeb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("SummarizeWeb() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
