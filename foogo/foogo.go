package foogo

type Foogo struct {
	Bar string
}

func (f *Foogo) SayYourThing() string {
	return f.Bar
}
