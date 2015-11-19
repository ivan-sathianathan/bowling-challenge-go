package main

import ("fmt"
        "math/rand"
        "time"
      )

  var scorecard = []int{}
  // var frames = [][]int{}
  var new = [][]int{}
  var framesTotals = []int{}
  var frame = []int{}
  var bowl1 int
  var bowl2 int

func bowl() {
  for i :=1; i <=10; i++ {
    rand.Seed(time.Now().Unix())
    bowl1 = rand.Intn(11)
  if bowl1 == 10 {
    bowl2 = 0
  } else {
    rand.Seed(time.Now().Unix())
    bowl2 = rand.Intn(11-bowl1)
  }
  pushScores(bowl1, bowl2)
  }
}

func pushScores(bowl1, bowl2 int) {
  scorecard = append(scorecard, bowl1, bowl2)
}

func addScores(scorecard []int) int {
  sum := 0
  for _, value := range scorecard {
    sum += value
  }
  return sum
}

func createFrames(scorecard []int) [][]int {
  frames := make([][]int, 10)
  for i, idx := 0, 0; i < 10; i, idx = i + 1, idx + 2 {
    frames[i] = scorecard[idx:idx+2]
  }
  // frames = append(frames)
  return frames
}

func calcBonusScore(frames [][]int) ([]int) {
  for i := 0; i < len(frames); i++ {
    frame := frames[i]
    sum := 0
      for _, value := range frame {
        sum += value
      }
      framesTotals = append(framesTotals, sum)
  }
  return framesTotals

}


// func appendFunc() {
//
// }


func main () {
  // bowl()
  pushScores(2, 5)
  pushScores(3, 4)
  pushScores(5, 5)
  pushScores(10, 0)
  pushScores(2, 4)
  pushScores(5, 4)
  pushScores(5, 4)
  pushScores(5, 4)
  pushScores(5, 4)
  pushScores(5, 4)

  fmt.Println(scorecard)
  fmt.Println(addScores(scorecard))
  fmt.Println(createFrames(scorecard))
  fmt.Println(calcBonusScore(createFrames(scorecard)))
  // fmt.Println(frames)
}
