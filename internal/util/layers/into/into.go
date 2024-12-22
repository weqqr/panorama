package into

func SliceOf[C any, R any, RAddr interface {
	*R
	Into() C
}](r []R, err error) ([]C, error) {
	if err != nil {
		return nil, err
	}

	var domainResults []C

	for _, result := range r {
		domainResults = append(domainResults, RAddr(&result).Into())
	}

	return domainResults, nil
}
