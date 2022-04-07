package main

// Интерфейсы — это инструменты для определения наборов действий и поведения. Они позволяют объектам опираться на абстракции, а не фактические реализации других объектов

type Writer interface {
	Write([]byte) (int, error)
}

func main()  {

}
