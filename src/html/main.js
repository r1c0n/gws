document.addEventListener('DOMContentLoaded', function () {
  var accordionItems = document.querySelectorAll('.accordion-item')

  accordionItems.forEach(function (item) {
    item.addEventListener('click', function () {
      var content = this.querySelector('.accordion-content')
      content.style.display =
        content.style.display === 'block' ? 'none' : 'block'
    })
  })
})
