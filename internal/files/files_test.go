package files

import (
	"os"
	"testing"
)

func Test_safeWriter_writeln(t *testing.T) {
	f, err := os.Create("fileForTest")

	type args struct {
		s string
	}
	tests := []struct {
		name string
		sw   *safeWriter
		args args
	}{
		{
			name: "Test for safeWriter_writeln",
			sw: &safeWriter{
				w:   f,
				err: err,
			},
			args: args{
				s: "",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sw.writeln(tt.args.s)
		})
	}
}

func TestSaveFile(t *testing.T) {
	type args struct {
		name string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test for SaveFile1",
			args: args{
				name: "test2",
				data: "test data",
			},
			wantErr: false,
		},
		{
			name: "Test for SaveFile2",
			args: args{
				name: ":",
				data: "test data",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFile(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantText string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := GetFileFromURL(tt.args.url); gotText != tt.wantText {
				t.Errorf("GetFileFromURL() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func TestSaveTxt(t *testing.T) {
	type args struct {
		name string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test for SaveTxt1",
			args: args{
				name: "test3",
				data: "test data",
			},
			wantErr: false,
		},
		{
			name: "Test for SaveTxt2",
			args: args{
				name: ":",
				data: "test data",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveTxt(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveTxt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test for Exists1",
			args: args{
				name: "test3",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test for Exists2",
			args: args{
				name: "777",
			},
			want:    false,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exists(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
