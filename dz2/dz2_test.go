package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_topWords(t *testing.T) {
	type args struct {
		text     string
		countTop int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Т1", args{text: "тест тест", countTop: 10}, []string{"тест"}},
		{"Т2", args{text: "тест тест ттест", countTop: 10}, []string{"тест", "ттест"}},
		{"Т3", args{text: "тест тест ттест ттест тт аа бб", countTop: 2}, []string{"тест", "ттест"}},
		{"Т4", args{text: "Принимая во внимание показатели успешности, перспективное планирование, а также свежий взгляд на привычные вещи - безусловно открывает новые горизонты для приоритизации разума над эмоциями. А также явные признаки победы институционализации призывают нас к новым свершениям, которые, в свою очередь, должны быть превращены в посмешище, хотя само их существование приносит несомненную пользу обществу. Высокий уровень вовлечения представителей целевой аудитории является четким доказательством простого факта: существующая теория, в своём классическом представлении, допускает внедрение системы массового участия. Равным образом, реализация намеченных плановых заданий создаёт необходимость включения в производственный план целого ряда внеочередных мероприятий с учётом комплекса существующих финансовых и административных условий. А ещё действия представителей оппозиции формируют глобальную экономическую сеть и при этом - подвергнуты целой серии независимых исследований. Активно развивающиеся страны третьего мира являются только методом политического участия и подвергнуты целой серии независимых исследований. Каждый из нас понимает очевидную вещь: глубокий уровень погружения обеспечивает актуальность системы обучения кадров, соответствующей насущным потребностям. В частности, выбранный нами инновационный путь требует от нас анализа форм воздействия. Выбранный нами инновационный путь предполагает независимые способы реализации своевременного выполнения сверхзадачи.", countTop: 5}, []string{"нас", "также", "инновационный", "нами", "выбранный"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topWords(tt.args.text, tt.args.countTop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topWords() = %v, want %v", got, tt.want)
				fmt.Printf("%s: Fail\n", tt.name)
			} else {
				fmt.Printf("%s: Ok\n", tt.name)
			}
		})
	}
}
