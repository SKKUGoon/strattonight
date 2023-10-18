package ws

import "sync"

type RaceSafeStore struct {
	sync.Mutex
	Integer   int
	Number    float64
	NumberArr []float64
}

func (s *RaceSafeStore) IntegerAdd(v int) {
	s.Lock()
	defer s.Unlock()
	s.Integer += v
}

func (s *RaceSafeStore) IntegerMul(v int) {
	s.Lock()
	defer s.Unlock()
	s.Integer *= v
}

func (s *RaceSafeStore) IntegerSet(v int) {
	s.Lock()
	defer s.Unlock()
	s.Integer = v
}

func (s *RaceSafeStore) NumberAdd(v float64) {
	s.Lock()
	defer s.Unlock()
	s.Number += v
}

func (s *RaceSafeStore) NumberMul(v float64) {
	s.Lock()
	defer s.Unlock()
	s.Number *= v
}

func (s *RaceSafeStore) NumberSet(v float64) {
	s.Lock()
	defer s.Unlock()
	s.Number = v
}

func (s *RaceSafeStore) Append(v float64) {
	s.Lock()
	defer s.Unlock()
	s.NumberArr = append(s.NumberArr, v)
}

func (s *RaceSafeStore) Pop() {
	s.Lock()
	defer s.Unlock()
	s.NumberArr = s.NumberArr[:len(s.NumberArr)-1]
}

func (s *RaceSafeStore) Replace(idx int, value float64) {
	s.Lock()
	defer s.Unlock()
	s.NumberArr[idx] = value
}
