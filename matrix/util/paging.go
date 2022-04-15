package util

func Paginate(data []interface{}, page, size int) []interface{} {
	if data == nil || len(data) == 0 {
		return data
	}

	count := len(data)
	var pageCount int
	if count%size == 0 {
		pageCount = count / size
	} else {
		pageCount = count/size + 1
	}

	var fromIndex, toIndex int
	if page > pageCount {
		page = pageCount
	}
	if page != pageCount {
		fromIndex = (page - 1) * size
		toIndex = fromIndex + size
	} else {
		fromIndex = (page - 1) * size
		toIndex = count
	}

	return data[fromIndex:toIndex]
}
