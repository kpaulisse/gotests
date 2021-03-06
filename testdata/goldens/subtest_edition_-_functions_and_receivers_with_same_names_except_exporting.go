package testdata

import "testing"

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeNameSameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SameTypeName{}
			got, err := tr.SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeNameSameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SameTypeName{}
			got, err := tr.sameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeNameSameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &sameTypeName{}
			got, err := tr.SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("sameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeNameSameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &sameTypeName{}
			got, err := tr.sameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("sameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}
