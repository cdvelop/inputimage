package inputimage

import "github.com/cdvelop/model"

func Image() *model.Input {
	in := input{}

	return &model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

func (input) Name() string {
	return "Image"
}

func (input) HtmlName() string {
	return "file"
}

func (input) ValidateField(data_in string, skip_validation bool, options ...string) error {
	return nil
}

func (input) GoodTestData() (out []string) {

	return
}

func (input) WrongTestData() (out []string) {
	return
}
