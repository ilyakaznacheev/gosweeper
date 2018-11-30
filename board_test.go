package main

import (
	"testing"
)

func TestGenerateFields(t *testing.T) {
	type args struct {
		width      int
		height     int
		x          int
		y          int
		mineNumber int
	}
	tests := []struct {
		name       string
		args       args
		mineNumber int
	}{
		{
			name: "1",
			args: args{
				width:      3,
				height:     3,
				x:          2,
				y:          2,
				mineNumber: 2,
			},
			mineNumber: 2,
		},
		{
			name: "2",
			args: args{
				width:      3,
				height:     3,
				x:          2,
				y:          2,
				mineNumber: 8,
			},
			mineNumber: 8,
		},
		{
			name: "3",
			args: args{
				width:      100,
				height:     100,
				x:          11,
				y:          22,
				mineNumber: 30,
			},
			mineNumber: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := generateFields(tt.args.width, tt.args.height, tt.args.x, tt.args.y, tt.args.mineNumber)
			mineCount := 0
			for xIdx := range fields {
				for yIdx := range fields[xIdx] {
					if fields[xIdx][yIdx].mine {
						mineCount++
					}
				}
			}
			if mineCount != tt.mineNumber {
				t.Errorf("wrong number of mines generated %d, but expected %d", mineCount, tt.mineNumber)
			}
		})
	}
}

func TestNewBoard(t *testing.T) {
	type args struct {
		width      uint
		height     uint
		x          uint
		y          uint
		mineNumber uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// normal board creation
		{
			name: "1",
			args: args{
				width:      10,
				height:     10,
				x:          5,
				y:          5,
				mineNumber: 5,
			},
			wantErr: false,
		},
		// point is outside board
		{
			name: "2",
			args: args{
				width:      10,
				height:     10,
				x:          10,
				y:          5,
				mineNumber: 5,
			},
			wantErr: true,
		},
		// point is outside board
		{
			name: "3",
			args: args{
				width:      10,
				height:     10,
				x:          5,
				y:          10,
				mineNumber: 5,
			},
			wantErr: true,
		},
		// too much mines
		{
			name: "4",
			args: args{
				width:      3,
				height:     3,
				x:          2,
				y:          2,
				mineNumber: 30,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, err := NewBoard(tt.args.width, tt.args.height, tt.args.x, tt.args.y, tt.args.mineNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("error is %s, but error expectation is: %v", err, tt.wantErr)
				return
			}
			if board == nil && !tt.wantErr {
				t.Error("board object was not created")
			}
		})
	}
}

func TestBoardGetStatus(t *testing.T) {
	type fields struct {
		fields [][]*Field
		width  int
		height int
	}
	type args struct {
		x uint
		y uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// a mine
		{
			name: "1",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{false},
					},
				},
				width:  1,
				height: 2,
			},
			args: args{
				x: 0,
				y: 0,
			},
			want:    StatusMine,
			wantErr: false,
		},
		// not a mine
		{
			name: "2",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{false},
					},
				},
				width:  1,
				height: 2,
			},
			args: args{
				x: 0,
				y: 1,
			},
			want:    1,
			wantErr: false,
		},
		// all neighbours are mines
		{
			name: "3",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{true},
						&Field{true},
					},
					{
						&Field{true},
						&Field{false},
						&Field{true},
					},
					{
						&Field{true},
						&Field{true},
						&Field{true},
					},
				},
				width:  3,
				height: 3,
			},
			args: args{
				x: 1,
				y: 1,
			},
			want:    8,
			wantErr: false,
		},
		// some neighbours are mines
		{
			name: "4",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{false},
						&Field{true},
						&Field{false},
					},
					{
						&Field{true},
						&Field{false},
						&Field{true},
					},
					{
						&Field{false},
						&Field{true},
						&Field{false},
					},
				},
				width:  3,
				height: 3,
			},
			args: args{
				x: 1,
				y: 1,
			},
			want:    4,
			wantErr: false,
		},
		// corner case 1
		{
			name: "5",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{false},
						&Field{true},
					},
					{
						&Field{true},
						&Field{true},
					},
				},
				width:  2,
				height: 2,
			},
			args: args{
				x: 0,
				y: 0,
			},
			want:    3,
			wantErr: false,
		},
		// corner case 2
		{
			name: "6",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{false},
					},
					{
						&Field{true},
						&Field{true},
					},
				},
				width:  2,
				height: 2,
			},
			args: args{
				x: 0,
				y: 1,
			},
			want:    3,
			wantErr: false,
		},
		// corner case 3
		{
			name: "7",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{true},
					},
					{
						&Field{false},
						&Field{true},
					},
				},
				width:  2,
				height: 2,
			},
			args: args{
				x: 1,
				y: 0,
			},
			want:    3,
			wantErr: false,
		},
		// corner case 4
		{
			name: "8",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
						&Field{true},
					},
					{
						&Field{true},
						&Field{false},
					},
				},
				width:  2,
				height: 2,
			},
			args: args{
				x: 1,
				y: 1,
			},
			want:    3,
			wantErr: false,
		},
		// error point outside board
		{
			name: "9",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
					},
				},
				width:  1,
				height: 1,
			},
			args: args{
				x: 1,
				y: 0,
			},
			want:    0,
			wantErr: true,
		},
		// error point outside board
		{
			name: "10",
			fields: fields{
				fields: [][]*Field{
					{
						&Field{true},
					},
				},
				width:  1,
				height: 1,
			},
			args: args{
				x: 0,
				y: 1,
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				fields: tt.fields.fields,
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			got, err := b.GetStatus(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("error is %s, but error expectation is: %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mine number is %v, but expected %v", got, tt.want)
			}
		})
	}
}
