package main

import (
	"github.com/supme/directEmail"
	"time"
	"fmt"
)

func main() {
	email := directEmail.New()

	email.FromName = "Отправитель"

	email.FromEmail = "sender@example.com"
	email.FromName = "Отправитель"

	email.ToEmail = "user@example.com"
	email.ToName = "Получатель"

	email.Header(fmt.Sprintf("Message-ID: <test_message_%d>", time.Now().Unix()))
	email.Header("Content-Language: ru")
	email.Subject = "Тест отправки email"

	email.Part(directEmail.TypeTextHTML, `
<h2>Тема: «Первоначальный общекультурный цикл: гипотеза и теории»</h2>
<p>Беллетристика, в первом приближении, возможна. Искусство имитирует "кодекс деяний". Одиночество иллюстрирует непосредственный статус художника.</p>
<p>Художественная элита, в первом приближении, выстраивает психологический параллелизм, таким образом, все перечисленные признаки архетипа и мифа подтверждают, что действие механизмов мифотворчества сродни механизмам художественно-продуктивного мышления. Прекрасное многопланово начинает резкий стиль. Либидо начинает мимезис. Бессознательное, в том числе, аккумулирует комплекс агрессивности, именно об этом комплексе движущих сил писал З.Фрейд в теории сублимации. Биографический метод продолжает архетип. Ритм, на первый взгляд, готично трансформирует непосредственный структурализм, именно об этом комплексе движущих сил писал З.Фрейд в теории сублимации.</p>
<p>Типическое, так или иначе, выстраивает модернизм. Модернизм, как бы это ни казалось парадоксальным, монотонно начинает модернизм, так Г.Корф формулирует собственную антитезу. Ф.Шилер, Г.Гете, Ф.Шлегели и А.Шлегели выразили типологическую антитезу классицизма и романтизма через противопоставление искусства "наивного" и "сентиментального", поэтому мимезис использует самодостаточный синтаксис искусства. Иррациональное в творчестве изящно представляет собой неизменный фарс.</p>
	`)

	email.Part(directEmail.TypeTextPlain, `
Тема: «Первоначальный общекультурный цикл: гипотеза и теории»
Беллетристика, в первом приближении, возможна. Искусство имитирует "кодекс деяний". Одиночество иллюстрирует непосредственный статус художника.
Художественная элита, в первом приближении, выстраивает психологический параллелизм, таким образом, все перечисленные признаки архетипа и мифа подтверждают, что действие механизмов мифотворчества сродни механизмам художественно-продуктивного мышления. Прекрасное многопланово начинает резкий стиль. Либидо начинает мимезис. Бессознательное, в том числе, аккумулирует комплекс агрессивности, именно об этом комплексе движущих сил писал З.Фрейд в теории сублимации. Биографический метод продолжает архетип. Ритм, на первый взгляд, готично трансформирует непосредственный структурализм, именно об этом комплексе движущих сил писал З.Фрейд в теории сублимации.
Типическое, так или иначе, выстраивает модернизм. Модернизм, как бы это ни казалось парадоксальным, монотонно начинает модернизм, так Г.Корф формулирует собственную антитезу. Ф.Шилер, Г.Гете, Ф.Шлегели и А.Шлегели выразили типологическую антитезу классицизма и романтизма через противопоставление искусства "наивного" и "сентиментального", поэтому мимезис использует самодостаточный синтаксис искусства. Иррациональное в творчестве изящно представляет собой неизменный фарс.
`)

//	email.Attachment("/home/aagafonov/Golang/src/github.com/supme/directEmail/example/StickAndCarrot.jpg")

	email.Render()
	print("\n", string(email.GetRawMessage()), "\n")

	err := email.Send()
	if err != nil {
		print("Send email with error:", err.Error())
	}

}

