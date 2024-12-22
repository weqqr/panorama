package from

func SliceOf[R interface{ From(C) R }, C any](c []C, err error) ([]R, error) {
	if err != nil {
		return nil, err
	}

	var domainResults []R

	for _, result := range c {
		domainResults = append(domainResults, (*new(R)).From(result))
	}

	return domainResults, nil
}
