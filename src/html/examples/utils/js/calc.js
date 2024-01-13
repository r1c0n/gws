// variable for resultField
const resultField = document.getElementById('result')

// event listener for keyboard input
document.addEventListener('keydown', handleKeyPress)

// variable to store memory value
let memoryValue = 0

// function to handle keyboard input
function handleKeyPress(event) {
  // get the key pressed
  const keyPressed = event.key

  // check if it's a number
  if (/[0-9]/.test(keyPressed)) {
    appendNumber(keyPressed)
  }
  // check if it's an operator
  else if (['+', '-', '*', '/'].includes(keyPressed)) {
    appendOperator(keyPressed)
  }
  // check for decimal point
  else if (keyPressed === '.' || keyPressed === ',') {
    appendDecimalPoint()
  }
  // check for enter or equal sign
  else if (keyPressed === 'Enter' || keyPressed === '=') {
    calculateResult()
  }
  // check for escape key
  else if (keyPressed === 'Escape') {
    clearResult()
  }
  // check for 'm' key for memory
  else if (keyPressed === 'm') {
    addToMemory()
  }
  // check for 'r' key for recalling from memory
  else if (keyPressed === 'r') {
    recallMemory()
  }
  // check for backspace key
  else if (keyPressed === 'Backspace') {
    backspace()
  }
}

// function to append a number to the result
function appendNumber(num) {
  resultField.value += num
}

// function to append an operator to the result
function appendOperator(operator) {
  const currentValue = resultField.value
  const lastChar = currentValue.charAt(currentValue.length - 1)

  // check if last character is an operator, replace it
  if (['+', '-', '*', '/'].includes(lastChar)) {
    resultField.value = currentValue.slice(0, -1) + operator
  }
  // otherwise, append the operator
  else {
    resultField.value += operator
  }
}

// function to append a decimal point to the result
function appendDecimalPoint() {
  const currentValue = resultField.value
  if (!currentValue.includes('.')) {
    resultField.value += '.'
  }
}

// function to calculate the result
function calculateResult() {
  const currentValue = resultField.value
  try {
    const result = eval(currentValue)
    if (result === undefined || isNaN(result)) {
      throw new Error('Invalid calculation')
    }
    resultField.value = result
  } catch (error) {
    resultField.value = 'Invalid Input'
  }
}

// function to clear the result
function clearResult() {
  resultField.value = ''
}

// function to add the current value to memory
function addToMemory() {
  const currentValue = parseFloat(resultField.value)
  if (!isNaN(currentValue)) {
    memoryValue += currentValue
  }
}

// function to recall the value from memory
function recallMemory() {
  resultField.value = memoryValue
}

// function to handle backspace
function backspace() {
  const currentValue = resultField.value
  resultField.value = currentValue.slice(0, -1)
}
