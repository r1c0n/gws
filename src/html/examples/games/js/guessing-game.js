// generate a random number between 1 and 100
let randomNumber = Math.floor(Math.random() * 100) + 1

// select elements with class names 'guesses', 'lastResult', and 'lowOrHi'
const guesses = document.querySelector('.guesses')
const lastResult = document.querySelector('.lastResult')
const lowOrHi = document.querySelector('.lowOrHi')

// select elements with class names 'guessSubmit' and 'guessField'
const guessSubmit = document.querySelector('.guessSubmit')
const guessField = document.querySelector('.guessField')

// initialize guess count and resetButton variable
let guessCount = 1
let resetButton

// function to check the user's guess
function checkGuess() {
  // get the user's guess from the input field
  let userGuess = Number(guessField.value)

  // if it's the first guess, display 'Previous guesses: '
  if (guessCount === 1) {
    guesses.textContent = 'Previous guesses: '
  }

  // append the user's guess to the list of previous guesses
  guesses.textContent += userGuess + ' '

  // if the user's guess is correct
  if (userGuess === randomNumber) {
    lastResult.textContent = 'Congratulations! You got it right!'
    lastResult.style.backgroundColor = 'green'
    lowOrHi.textContent = ''

    // call the setGameOver function
    setGameOver()
  }
  // if the user has used all 10 guesses
  else if (guessCount === 10) {
    lastResult.textContent = '!!!GAME OVER!!!'

    // call the setGameOver function
    setGameOver()
  }
  // if the guess was incorrect
  else {
    lastResult.textContent = 'Wrong!'
    lastResult.style.backgroundColor = 'red'

    // provide a hint based on whether the guess was too high or too low
    if (userGuess < randomNumber) {
      lowOrHi.textContent = 'Last guess was too low!'
    } else if (userGuess > randomNumber) {
      lowOrHi.textContent = 'Last guess was too high!'
    }
  }

  // increment the guess count, clear the input field, and focus on it
  guessCount++
  guessField.value = ''
  guessField.focus()
}

// add an event listener to the 'Submit guess' button
guessSubmit.addEventListener('click', checkGuess)

// function to set the game over state
function setGameOver() {
  // disable the input field and the 'Submit guess' button
  guessField.disabled = true
  guessSubmit.disabled = true

  // create a 'Start new game' button
  resetButton = document.createElement('button')
  resetButton.textContent = 'Start new game'
  document.body.appendChild(resetButton)

  // add an event listener to the 'Start new game' button
  resetButton.addEventListener('click', resetGame)
}

// function to reset the game
function resetGame() {
  // reset guess count and clear result paragraphs
  guessCount = 1
  const resetParas = document.querySelectorAll('.resultParas p')
  for (let i = 0; i < resetParas.length; i++) {
    resetParas[i].textContent = ''
  }

  // remove the 'Start new game' button
  resetButton.parentNode.removeChild(resetButton)

  // enable the input field and 'Submit guess' button, clear the input field, and focus on it
  guessField.disabled = false
  guessSubmit.disabled = false
  guessField.value = ''
  guessField.focus()

  // reset the background color of the last result paragraph
  lastResult.style.backgroundColor = 'white'

  // generate a new random number
  randomNumber = Math.floor(Math.random() * 100) + 1
}
