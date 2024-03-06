package db

import "context"

func (s *sDB) FetchJobStatus(ctx context.Context, userId string) (int, error) {
	return 0, nil
}
func (s *sDB) UpJobStatus(ctx context.Context, userId string, seqId string) error {
	return nil
}
