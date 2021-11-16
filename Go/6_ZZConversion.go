func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	row, step := 0, 1
	for i := range s {
		res[row] = append(res[row], s[i])
		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}
		row += step
	}
	var buf bytes.Buffer
	for _, v := range res {
		buf.Write(v)
	}
	return buf.String()
}