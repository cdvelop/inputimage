package inputimage

func (i input) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	return `<input id="` + id + `" type="file" id="avatar" name="` + field_name + `" accept="image/,.png,.jpeg,.jpg" />`
}
