package requests

import (
	"net/http"
	"testing"
)

func TestDefaultHttpGetter_Get(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		d        DefaultHttpGetter
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
		{
			name:     "Positive case",
			d:        DefaultHttpGetter{},
			args:     args{url: URLAllUsers},
			wantResp: &http.Response{StatusCode: 200},
			wantErr:  false,
		},
		{
			name:     "404 error test case",
			d:        DefaultHttpGetter{},
			args:     args{url: "https://fakestoreapi.com/users2"},
			wantResp: &http.Response{StatusCode: 404},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.d.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultHttpGetter.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.StatusCode != tt.wantResp.StatusCode {
				t.Errorf("DefaultHttpGetter.Get() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
