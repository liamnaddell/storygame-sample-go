package parts

import s "github.com/skilstak/storyeng-go"

s.Parts["Start"] = s.Part{
	ENTER: "Welcome to this story game.",
	INPUT: "Name",
}

s.Parts["Name"] = s.Part{
	ENTER: "My name is Norman. What's yours?",
	oninput: func(e) {
		s.Data["name"] = e.Line
	},
	LEAVE: "Nice to meet you {{.name}}.",
}
