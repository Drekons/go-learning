package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := topWords("Принимая во внимание показатели успешности, перспективное планирование, а также свежий взгляд на привычные вещи - безусловно открывает новые горизонты для приоритизации разума над эмоциями. А также явные признаки победы институционализации призывают нас к новым свершениям, которые, в свою очередь, должны быть превращены в посмешище, хотя само их существование приносит несомненную пользу обществу. Высокий уровень вовлечения представителей целевой аудитории является четким доказательством простого факта: существующая теория, в своём классическом представлении, допускает внедрение системы массового участия. Равным образом, реализация намеченных плановых заданий создаёт необходимость включения в производственный план целого ряда внеочередных мероприятий с учётом комплекса существующих финансовых и административных условий. А ещё действия представителей оппозиции формируют глобальную экономическую сеть и при этом - подвергнуты целой серии независимых исследований. Активно развивающиеся страны третьего мира являются только методом политического участия и подвергнуты целой серии независимых исследований. Каждый из нас понимает очевидную вещь: глубокий уровень погружения обеспечивает актуальность системы обучения кадров, соответствующей насущным потребностям. В частности, выбранный нами инновационный путь требует от нас анализа форм воздействия. Выбранный нами инновационный путь предполагает независимые способы реализации своевременного выполнения сверхзадачи.", 10)

	fmt.Println(words)
}

type WordTop struct {
	Word  string
	Count int
}

type ByCount []WordTop

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func topWords(text string, countTop int) []string {
	var textArr []string = strings.Split(text, " ")
	var words []WordTop
	var returnArr []string

	for _, word := range textArr {
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ToLower(word)

		if len(word) <= 2 {
			continue
		}

		var newFlag bool = true
		for i, top := range words {
			if top.Word == word {
				words[i].Count++
				newFlag = false
				break
			}
		}

		if newFlag {
			words = append(words, WordTop{Word: word, Count: 1})
		}
	}

	sort.Sort(ByCount(words))

	if len(words) > countTop {
		words = words[:countTop]
	}

	for _, word := range words {
		returnArr = append(returnArr, word.Word)
	}

	return returnArr
}
