function addTodo() {
  var input = document.getElementById('newTodo')
  var todoList = document.getElementById('todoList')

  if (input.value.trim() !== '') {
    var li = document.createElement('li')
    li.className = 'todo-item'
    li.innerHTML = `
  <span>${input.value}</span>
  <button class="delete-btn" onclick="removeTodo(this)">X</button>
`
    todoList.appendChild(li)
    input.value = ''
    li.style.animation = 'fadeIn 0.5s'
  }
}

function removeTodo(btn) {
  var todoItem = btn.parentElement
  todoItem.style.animation = 'fadeOut 0.5s'
  setTimeout(function () {
    todoItem.remove()
  }, 500)
}
