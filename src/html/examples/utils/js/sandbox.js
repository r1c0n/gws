function runJQuery() {
  $(document).ready(function () {
    $('button').click(function () {
      $(this).text('Clicked!')
    })
  })
}

function runLodash() {
  var array = [1, 2, 3, 4, 5]
  var squared = _.map(array, function (num) {
    return num * num
  })
  alert(squared)
}

function runSweetAlert() {
  swal('Good job!', 'You clicked the button!', 'success')
}

function runMoment() {
  var now = moment()
  alert('Current date and time: ' + now.format('YYYY-MM-DD HH:mm:ss'))
}
