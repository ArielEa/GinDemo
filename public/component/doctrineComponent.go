package component

var ARRAY_MODE string = "array"

var SINGLE_MODE string = "oneornull"

var searchSql string

var returnMode string = ARRAY_MODE

func DonctrineInit(sqlMessgae string) {

	searchSql = "abc"

	GetArrayResult()
}

func GetArrayResult() ([][]interface{}, error) {

	var records [][]interface{}

	return records, nil
}

func GetOneOrNullResult() ([]string, error) {
	var record []string

	return record, nil
}

func GetHandleTime() (string, error) {
	return "時間は234s", nil
}

func GetExplainMessage() (string, error) {
	return "", nil
}

func GetIntId() (int, error) {
	return 2, nil
}
