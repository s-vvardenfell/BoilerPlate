package utility

// func DoRequest(client *http.Client, method, url string, body io.Reader) *http.Response {
// 	req, err := http.NewRequest(method, url, body)
// 	if err != nil {
// 		logrus.Fatalf("cannot create NewRequest to %s, %v", url, err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		logrus.Fatalf("cannot get %s, %v", url, err)
// 	}
// 	return resp
// }

// func YamlToStruct(byteValue []byte, strct interface{}) error {
// 	if err := yaml.Unmarshal(byteValue, &strct); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func JsonToStruct(byteValue []byte, strct interface{}) error {
// 	if err := json.Unmarshal(byteValue, &strct); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func JsonToMap(byteValue []byte) (map[string]interface{}, error) {
// 	var result map[string]interface{}
// 	if err := json.Unmarshal([]byte(byteValue), &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func BytesFromReader(r io.Reader) []byte {
// 	byteValue, err := io.ReadAll(r)
// 	if err != nil {
// 		logrus.Fatalf("error reading in BytesFromReader(), %v", err)
// 	}
// 	return byteValue
// }

// func ElementByLocator(byteValue []byte, val, attr string) (string, error) {
// 	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(byteValue))
// 	if err != nil {
// 		return "", err
// 	}

// 	sel := doc.Find(val)
// 	res, ok := sel.Eq(0).Attr(attr)
// 	if ok {
// 		return res, nil
// 	}
// 	return "", fmt.Errorf("element '%s' with attr '%s' not found", val, attr)
// }

// func ElementsByLocators(byteValue []byte, m map[string]string) error {
// 	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(byteValue))
// 	if err != nil {
// 		return err
// 	}

// 	for lctr, attr := range m {
// 		sel := doc.Find(lctr)
// 		val, ok := sel.Eq(0).Attr(attr)
// 		if !ok {
// 			return fmt.Errorf("cannot parse element with locator %s and attribute %s", lctr, attr)
// 		} else {
// 			m[lctr] = val
// 		}
// 	}
// 	return nil
// }

// func createTemplate(name, t string) *template.Template {
// 	return template.Must(template.New(name).Parse(t))
// }

// func PrepareRequestBody(keyword, query string) *bytes.Buffer {
// 	var r bytes.Buffer
// 	t := createTemplate("template", query)
// 	err := t.Execute(&r, url.QueryEscape(keyword))
// 	if err != nil {
// 		logrus.Fatalf("cannot make template query with word %s, %w", keyword, err)
// 	}
// 	return &r
// }
