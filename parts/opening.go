package parts

import s "github.com/skilstak/storyeng-go"

func init() {

	s.Parts["Start"] = s.Part{
		"onenter": "Welcome to this story game.",
		"oninput": "Name",
	}

	s.Parts["Name"] = s.Part{
		"onenter": "My name is Norman. What's yours?",
		"oninput": func (e s.Event) {
			s.Data["name"] = e.Line
		},
		"onleave": "Nice to meet you {{.name}}.",
	}

}
