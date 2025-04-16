package storage_monitoring

import (
	"container/heap"
	"reflect"
	"sync"
	"testing"
)

func Test_GetBestStorageServerAddress(t *testing.T) {
	t.Parallel()

	type args struct {
		heap StorageHeap
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "best storage server with max disk space",
			args: args{
				heap: StorageHeap{
					{
						Address:   "server1",
						FreeSpace: 1000,
					},
					{
						Address:   "server2",
						FreeSpace: 5000,
					},
					{
						Address:   "server3",
						FreeSpace: 0,
					},
				},
			},
			want: "server2",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Service{storageHeap: tt.args.heap, mu: sync.RWMutex{}}
			heap.Init(&s.storageHeap)

			got, err := s.GetBestStorageServerAddress()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBestStorageServerAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBestStorageServerAddress() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
