package go_mock

func GetS3File(ctx context.Context, s3 string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s3, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}