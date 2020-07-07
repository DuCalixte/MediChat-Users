package utilSocket

import (
  "math/rand"
)

var phrases = []string{
    "An investment in education pays the best interest.",
    "Anyone who stops learning is old, whether at twenty or eighty.",
    "Education costs money, but then so does ignorance.",
    "Education is not preparation for life. Education is life itself.",
    "Get over the idea that only children should spend their time studying. Be a student as long as you still have something to learn, and this will mean all your life.",
    "I am always doing that which I cannot do in order that I may learn how to do it.",
"I am always ready to learn although I do not always like being taught.",
 "I hope that in this year to come, you make mistakes. Because if you are making mistakes, then you are making new things, trying new things, learning, living, pushing yourself, changing yourself, changing your world.",
    "I learned the value of hard work by working hard.",
"If the goal you've set for yourself has a 100 percent chance of success, then frankly you aren't aiming high enough.",
    "If you hold a cat by the tail you learn things you cannot learn any other way.",
    "If you're determined to learn, no one can stop you.",
"It is good to have an end to journey toward, but it is the journey that matters in the end.",
"Knowledge is of no value unless you put it into practice.",
"Live as if you were to die tomorrow. Learn as if you were to live forever.",
    "Never discourage anyone who continually makes progress, no matter how slow.",
    "Self education is the only kind of education there is.",
    "The noblest pleasure is the joy of understanding.",
    "Theory is knowledge that doesn't work. Practice is when everything works and you don't know why.",
"There are no foreign lands. It is the traveller only who is foreign.",
"There are no secrets to success. It is the result of preparation, hard work, and learning from failure.",
    "You can be discouraged by failure, or you can learn from it. So go ahead and make mistakes, make all you can. Because, remember that's where you'll find success - on the far side of failure.",
    "You can get help from teachers, but you're going to have to learn a lot by yourself, sitting alone in a room.",
    "You don't learn to walk by following rules. You learn by doing, and by falling over.",
    "You'll never know everything about anything.",
  }

func RandomPhrases() string {
  index := rand.Int() % len(phrases)
  return phrases[index]
}
