package pkg

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCalculateShardingKey(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				id: 87040,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logrus.Infof("%+v", CalculateShardingKey(tt.args.id))
		})
	}
}
