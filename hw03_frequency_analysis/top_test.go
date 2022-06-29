package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var textEngRus = `The voice spoke again. 
Голос заговорил снова: 
“I can’t hardly move with all these creeper things.” 
– Двинуться не дают, ух и цопкие они! 
The owner of the voice came backing out of the undergrowth so that twigs scratched on a greasy wind-breaker.
Тот, кому принадлежал голос, задом выбирался из кустов, с трудом выдирая у них свою грязную куртку.
The naked crooks of his knees were plump, caught and scratched by thorns.
Пухлые голые ноги коленками застряли в шипах и были все расцарапаны.
He bent down, removed the thorns carefully, and turned round.
Он наклонился, осторожно отцепил шипы и повернулся.
He was shorter than the fair boy and very fat.
Он был ниже светлого и очень толстый.
He came forward, searching out safe lodgments for his feet, and then looked up through thick spectacles. 
Сделал шаг, нащупав безопасную позицию, и глянул сквозь толстые очки. 
“Where’s the man with the megaphone?” 
– А где же дядька, который с мегафоном? 
The fair boy shook his head. 
Светлый покачал головой: 
“This is an island.
– Это остров.
At least I think it’s an island.
Так мне по крайней мере кажется.
That’s a reef out in the sea.
А там риф.
Perhaps there aren’t any grownups anywhere.” 
Может, даже тут вообще взрослых нет. 
The fat boy looked startled. `

var textSymbols = ",,,~~~ ~!@#$%^ &*()_ _ _ +?><} {.[]'\\"

var textDiff = "Привет пРивет привет? Привет, привет, ПРИВЕТ привет: приветы привета"

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10("", taskWithAsteriskIsCompleted), 0)
	})
	t.Run("symbols", func(t *testing.T) {
		require.Len(t, Top10(textSymbols, taskWithAsteriskIsCompleted), 0)
	})
	t.Run("Different words - positive", func(t *testing.T) {
		require.Len(t, Top10(textDiff, false), 9)
	})

	t.Run("Russian English text", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"the",
				"и",
				"and",
				"boy",
				"he",
				"his",
				"of",
				"out",
				"s",
				"a",
			}
			require.Equal(t, expected, Top10(textEngRus, taskWithAsteriskIsCompleted))
		} else {
			expected := []string{
				"the",
				"The",
				"и",
				"and",
				"He",
				"boy",
				"his",
				"of",
				"out",
				"–",
			}
			require.Equal(t, expected, Top10(textEngRus, taskWithAsteriskIsCompleted))
		}
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text, taskWithAsteriskIsCompleted))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text, taskWithAsteriskIsCompleted))
		}
	})
}
