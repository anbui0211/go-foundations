package unittest

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	type args struct {
		address  string
		ward     string
		district string
		province string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Case 1 origin",
			args: args{
				address:  "191 Lê Lợi",
				ward:     "Phường Hải Châu 1",
				district: "Quận Hải Châu",
				province: "Thành phố Đà Nẵng",
			},
			want: "191 Lê Lợi, P.Hải Châu 1, Q.Hải Châu, Đà Nẵng",
		},
		{
			name: "Case 2 trim",
			args: args{
				address:  "191      Lê Lợi",
				ward:     "Phường Hải Châu      1",
				district: "    Quận Hải Châu",
				province: "Thành       phố Đà Nẵng",
			},
			want: "191 Lê Lợi, P.Hải Châu 1, Q.Hải Châu, Đà Nẵng",
		},
		{
			name: "Case 2 duplicate address",
			args: args{
				address:  "191     Lê Lợi, Quận Hải Châu, Thành phố Đà Nẵng",
				ward:     "Phường Hải Châu      1",
				district: "    Quận Hải Châu",
				province: "Thành phố Đà Nẵng",
			},
			want: "191 Lê Lợi Quận Hải Châu Thành phố Đà Nẵng, P.Hải Châu 1, Q.Hải Châu, Đà Nẵng",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocation(tt.args.address, tt.args.ward, tt.args.district, tt.args.province); got != tt.want {
				t.Errorf("GetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
