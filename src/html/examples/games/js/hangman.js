// list of words for the game
const words = ['javascript', 'html', 'css', 'python', 'ruby', 'java']

// choose a random word from the list
let randomWord = words[Math.floor(Math.random() * words.length)]

// initialize variables
let guessedWord = '_'.repeat(randomWord.length)
let attempts = 6

// display the initial state of the game
$('#word').text(guessedWord)

// function to check the user's guess
$('#check').click(function () {
  let guess = $('#guess').val().toLowerCase()
  if (guess.length === 1) {
    if (randomWord.includes(guess)) {
      for (let i = 0; i < randomWord.length; i++) {
        if (randomWord[i] === guess) {
          guessedWord =
            guessedWord.substr(0, i) + guess + guessedWord.substr(i + 1)
        }
      }
      $('#word').text(guessedWord)
      if (guessedWord === randomWord) {
        $('#message').text('Congratulations! You guessed the word!')
        $('#check').prop('disabled', true)
      }
    } else {
      attempts--
      $('#attempts').text(`Attempts left: ${attempts}`)
      if (attempts === 0) {
        $('#message').text(
          `You ran out of attempts. The word was ${randomWord}.`
        )
        $('#check').prop('disabled', true)
      }
    }
    $('#guess').val('')
  }
})
