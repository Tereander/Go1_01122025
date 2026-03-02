package rectangle // проект main на весь проет один

type Rectangle struct {
	A, B int // эти поля экспортируемые
	color string // это поле - нет
}


func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}

